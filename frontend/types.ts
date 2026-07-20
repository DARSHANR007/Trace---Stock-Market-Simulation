
export interface OHLC {
  open: number;
  high: number;
  low: number;
  close: number;
}

export interface Order {
  quantity: number;
  price: number;
  orders: number;
}

export interface MarketDepth {
  buy: Order[];
  sell: Order[];
}

export interface StockData {
  ohlc: OHLC;
  depth: MarketDepth;
  timestamp: string;
  instrument_token: string;
  symbol: string;
  last_price: number;
  volume: number;
  average_price: number;
  oi: number;
  net_change: number;
  total_buy_quantity: number;
  total_sell_quantity: number;
  lower_circuit_limit: number;
  upper_circuit_limit: number;
  last_trade_time: string;
  oi_day_high: number;
  oi_day_low: number;
}

export interface APIResponse {
  status: string;
  data: {
    [key: string]: StockData;
  };
}

export interface SearchSuggestion {
  symbol: string;
  name: string;
  instrument_key: string;
  segment: string;
}
