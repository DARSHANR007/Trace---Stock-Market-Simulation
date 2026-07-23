package market

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const websocketGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

const (
	wsOpcodeContinuation = 0x0
	wsOpcodeText         = 0x1
	wsOpcodeBinary       = 0x2
	wsOpcodeClose        = 0x8
	wsOpcodePing         = 0x9
	wsOpcodePong         = 0xA
)

// LiveTick is a raw websocket message captured from a market feed.
type LiveTick struct {
	ReceivedAt time.Time
	Text       string
	Raw        []byte
}

// WebSocketClient is a small standalone websocket client for live market data.
type WebSocketClient struct {
	endpoint    string
	headers     http.Header
	dialTimeout time.Duration

	mu     sync.Mutex
	conn   net.Conn
	reader *bufio.Reader
}

// NewWebSocketClient creates a websocket client without wiring it into other services.
func NewWebSocketClient(endpoint string, headers http.Header) *WebSocketClient {
	clonedHeaders := make(http.Header)
	for key, values := range headers {
		clonedHeaders[key] = append([]string(nil), values...)
	}

	return &WebSocketClient{
		endpoint:    endpoint,
		headers:     clonedHeaders,
		dialTimeout: 10 * time.Second,
	}
}

// NewBearerWebSocketClient creates a websocket client with an Authorization bearer token.
func NewBearerWebSocketClient(endpoint, accessToken string) *WebSocketClient {
	headers := make(http.Header)
	if accessToken != "" {
		headers.Set("Authorization", "Bearer "+accessToken)
	}
	return NewWebSocketClient(endpoint, headers)
}

// Connect opens the websocket connection and completes the RFC6455 handshake.
func (c *WebSocketClient) Connect(ctx context.Context) error {
	if c == nil {
		return errors.New("websocket client is nil")
	}
	if c.endpoint == "" {
		return errors.New("websocket endpoint is required")
	}

	parsedURL, err := url.Parse(c.endpoint)
	if err != nil {
		return fmt.Errorf("parse websocket endpoint: %w", err)
	}
	if parsedURL.Scheme != "ws" && parsedURL.Scheme != "wss" {
		return fmt.Errorf("unsupported websocket scheme %q", parsedURL.Scheme)
	}

	address := parsedURL.Host
	if !strings.Contains(address, ":") {
		if parsedURL.Scheme == "wss" {
			address += ":443"
		} else {
			address += ":80"
		}
	}

	dialer := &net.Dialer{Timeout: c.dialTimeout}
	var rawConn net.Conn
	if parsedURL.Scheme == "wss" {
		tlsConn, err := tls.DialWithDialer(dialer, "tcp", address, &tls.Config{ServerName: parsedURL.Hostname()})
		if err != nil {
			return fmt.Errorf("dial websocket tls connection: %w", err)
		}
		rawConn = tlsConn
	} else {
		rawConn, err = dialer.DialContext(ctx, "tcp", address)
		if err != nil {
			return fmt.Errorf("dial websocket connection: %w", err)
		}
	}

	requestPath := parsedURL.RequestURI()
	if requestPath == "" {
		requestPath = "/"
	}

	keyBytes := make([]byte, 16)
	if _, err := rand.Read(keyBytes); err != nil {
		_ = rawConn.Close()
		return fmt.Errorf("generate websocket key: %w", err)
	}
	secKey := base64.StdEncoding.EncodeToString(keyBytes)

	requestHeaders := make(http.Header)
	for key, values := range c.headers {
		requestHeaders[key] = append([]string(nil), values...)
	}
	requestHeaders.Set("Host", parsedURL.Host)
	requestHeaders.Set("Upgrade", "websocket")
	requestHeaders.Set("Connection", "Upgrade")
	requestHeaders.Set("Sec-WebSocket-Version", "13")
	requestHeaders.Set("Sec-WebSocket-Key", secKey)

	if err := sendWebSocketHandshake(rawConn, parsedURL.Host, requestPath, requestHeaders); err != nil {
		_ = rawConn.Close()
		return err
	}

	response, err := http.ReadResponse(bufio.NewReader(rawConn), &http.Request{Method: http.MethodGet})
	if err != nil {
		_ = rawConn.Close()
		return fmt.Errorf("read websocket handshake response: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusSwitchingProtocols {
		_ = rawConn.Close()
		return fmt.Errorf("websocket handshake failed: %s", response.Status)
	}

	expectedAccept := websocketAcceptKey(secKey)
	if response.Header.Get("Sec-WebSocket-Accept") != expectedAccept {
		_ = rawConn.Close()
		return errors.New("websocket handshake validation failed")
	}

	reader := bufio.NewReader(rawConn)

	c.mu.Lock()
	c.conn = rawConn
	c.reader = reader
	c.mu.Unlock()

	return nil
}

// Close shuts down the websocket connection.
func (c *WebSocketClient) Close() error {
	if c == nil {
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return nil
	}

	_ = c.writeFrameLocked(wsOpcodeClose, nil)
	err := c.conn.Close()
	c.conn = nil
	c.reader = nil
	return err
}

// SendText sends a small text message to the websocket server.
func (c *WebSocketClient) SendText(payload []byte) error {
	if c == nil {
		return errors.New("websocket client is nil")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return errors.New("websocket is not connected")
	}

	return c.writeFrameLocked(wsOpcodeText, payload)
}

// Stream starts a reader goroutine and returns channels for live data and errors.
func (c *WebSocketClient) Stream(ctx context.Context) (<-chan LiveTick, <-chan error, error) {
	if err := c.Connect(ctx); err != nil {
		return nil, nil, err
	}

	dataCh := make(chan LiveTick)
	errorCh := make(chan error, 1)

	go func() {
		defer close(dataCh)
		defer close(errorCh)
		defer c.Close()

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			message, err := c.readMessage()
			if err != nil {
				if !errors.Is(err, io.EOF) && !errors.Is(err, context.Canceled) {
					select {
					case errorCh <- err:
					default:
					}
				}
				return
			}

			select {
			case dataCh <- LiveTick{ReceivedAt: time.Now().UTC(), Text: string(message), Raw: message}:
			case <-ctx.Done():
				return
			}
		}
	}()

	return dataCh, errorCh, nil
}

// StreamUpstoxLiveData is a convenience wrapper for opening a live feed stream.
func StreamUpstoxLiveData(ctx context.Context, endpoint, accessToken string) (<-chan LiveTick, <-chan error, error) {
	client := NewBearerWebSocketClient(endpoint, accessToken)
	return client.Stream(ctx)
}

func sendWebSocketHandshake(conn net.Conn, host, requestPath string, headers http.Header) error {
	var builder strings.Builder
	builder.WriteString("GET ")
	builder.WriteString(requestPath)
	builder.WriteString(" HTTP/1.1\r\n")
	builder.WriteString("Host: ")
	builder.WriteString(host)
	builder.WriteString("\r\n")

	for key, values := range headers {
		for _, value := range values {
			builder.WriteString(key)
			builder.WriteString(": ")
			builder.WriteString(value)
			builder.WriteString("\r\n")
		}
	}

	builder.WriteString("\r\n")
	_, err := io.WriteString(conn, builder.String())
	return err
}

func websocketAcceptKey(key string) string {
	sum := sha1.Sum([]byte(key + websocketGUID))
	return base64.StdEncoding.EncodeToString(sum[:])
}

func (c *WebSocketClient) readMessage() ([]byte, error) {
	c.mu.Lock()
	reader := c.reader
	conn := c.conn
	c.mu.Unlock()

	if reader == nil || conn == nil {
		return nil, errors.New("websocket is not connected")
	}

	var fragments [][]byte
	var started bool

	for {
		frame, err := readFrame(reader)
		if err != nil {
			return nil, err
		}

		switch frame.opcode {
		case wsOpcodeText, wsOpcodeBinary:
			started = true
			fragments = append(fragments, frame.payload)
			if frame.fin {
				return joinFragments(fragments), nil
			}
		case wsOpcodeContinuation:
			if !started {
				return nil, errors.New("unexpected websocket continuation frame")
			}
			fragments = append(fragments, frame.payload)
			if frame.fin {
				return joinFragments(fragments), nil
			}
		case wsOpcodePing:
			c.mu.Lock()
			err = c.writeFrameLocked(wsOpcodePong, frame.payload)
			c.mu.Unlock()
			if err != nil {
				return nil, err
			}
		case wsOpcodePong:
			continue
		case wsOpcodeClose:
			return nil, io.EOF
		default:
			if frame.fin {
				return frame.payload, nil
			}
		}
	}
}

type websocketFrame struct {
	fin     bool
	opcode  byte
	payload []byte
}

func readFrame(reader *bufio.Reader) (websocketFrame, error) {
	header := make([]byte, 2)
	if _, err := io.ReadFull(reader, header); err != nil {
		return websocketFrame{}, err
	}

	frame := websocketFrame{
		fin:    header[0]&0x80 != 0,
		opcode: header[0] & 0x0f,
	}

	masked := header[1]&0x80 != 0
	length := uint64(header[1] & 0x7f)

	switch length {
	case 126:
		extra := make([]byte, 2)
		if _, err := io.ReadFull(reader, extra); err != nil {
			return websocketFrame{}, err
		}
		length = uint64(binary.BigEndian.Uint16(extra))
	case 127:
		extra := make([]byte, 8)
		if _, err := io.ReadFull(reader, extra); err != nil {
			return websocketFrame{}, err
		}
		length = binary.BigEndian.Uint64(extra)
	}

	if length > uint64(maxFrameSize) {
		return websocketFrame{}, errors.New("websocket frame too large")
	}

	var maskKey [4]byte
	if masked {
		if _, err := io.ReadFull(reader, maskKey[:]); err != nil {
			return websocketFrame{}, err
		}
	}

	payload := make([]byte, length)
	if _, err := io.ReadFull(reader, payload); err != nil {
		return websocketFrame{}, err
	}

	if masked {
		for i := range payload {
			payload[i] ^= maskKey[i%4]
		}
	}

	frame.payload = payload
	return frame, nil
}

const maxFrameSize = 1 << 20

func joinFragments(fragments [][]byte) []byte {
	if len(fragments) == 0 {
		return nil
	}

	if len(fragments) == 1 {
		return fragments[0]
	}

	total := 0
	for _, fragment := range fragments {
		total += len(fragment)
	}

	merged := make([]byte, 0, total)
	for _, fragment := range fragments {
		merged = append(merged, fragment...)
	}
	return merged
}

func (c *WebSocketClient) writeFrameLocked(opcode byte, payload []byte) error {
	if c.conn == nil {
		return errors.New("websocket is not connected")
	}

	var frame []byte
	frame = append(frame, 0x80|opcode)

	length := len(payload)
	switch {
	case length <= 125:
		frame = append(frame, 0x80|byte(length))
	case length <= 65535:
		frame = append(frame, 0x80|126)
		extra := make([]byte, 2)
		binary.BigEndian.PutUint16(extra, uint16(length))
		frame = append(frame, extra...)
	default:
		frame = append(frame, 0x80|127)
		extra := make([]byte, 8)
		binary.BigEndian.PutUint64(extra, uint64(length))
		frame = append(frame, extra...)
	}

	var maskKey [4]byte
	if _, err := rand.Read(maskKey[:]); err != nil {
		return fmt.Errorf("generate websocket mask: %w", err)
	}
	frame = append(frame, maskKey[:]...)

	maskedPayload := make([]byte, len(payload))
	for i := range payload {
		maskedPayload[i] = payload[i] ^ maskKey[i%4]
	}
	frame = append(frame, maskedPayload...)

	_, err := c.conn.Write(frame)
	return err
}
