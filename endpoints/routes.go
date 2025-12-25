package endpoints

// API Endpoints
const (
	URIhost                string = "https://api.upstox.com"
	URIAuthorise           string = "%s/index/dialog/authorize?apiKey=%s&redirect_uri=%s&response_type=code"
	URIAccessToken         string = "%s/index/oauth/token"
	URIProfile             string = "%s/index/profile"
	URIMasterContract      string = "%s/index/master-contract/%s"
	URILogout              string = "%s/index/logout"
	URIHoldings            string = "%s/live/profile/holdings"
	URIBalance             string = "%s/live/profile/balance"
	URIBalanceCommodity    string = "%s/live/profile/balance/commodity"
	URIPositions           string = "%s/live/profile/positions"
	URIPlaceOrder          string = "%s/live/orders"
	URIGetOrders           string = "%s/live/orders"
	URIGetOrdersInfo       string = "%s/live/orders/%s"
	URIModifyOrder         string = "%s/live/orders/%s"
	URICancelOrder         string = "%s/live/orders/%s"
	URICancelAllOrders     string = "%s/live/orders"
	URITradesInfo          string = "%s/live/orders/%s/trades"
	URITradeBook           string = "%s/live/trade-book"
	URILiveFeed            string = "%s/live/feed/now/%s/%s/%s"
	URILiveFeedSubscribe   string = "%s/live/feed/sub/%s/%s?symbol=%s"
	URILiveFeedUnSubscribe string = "%s/live/feed/unsub/%s/%s?symbol=%s"
	URIOHLC                string = "%s/historical/ohlc/%s/%s?%s"
	URISocketParams        string = "%s/live/socket-params"
	URISocketEndpoint      string = "wss://ws-api.upstox.com"
)
