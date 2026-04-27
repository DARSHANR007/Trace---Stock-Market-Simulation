# Stage 1: Build your Go binary
FROM golang:1.25-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o server ./cmd/server/

# Stage 2: Run
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    ca-certificates \
    sqlite3 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the binary
COPY --from=builder /app/server .

# Copy the DB (already has all instrument data)
COPY instruments.db .

# Copy the UI folder (your app serves this as static files)
COPY ui/ ./ui/

EXPOSE 8000

ENTRYPOINT ["/app/server"]