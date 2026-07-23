
import React, { useState, useEffect, useCallback, useRef } from 'react';
import { 
  ArrowUp, 
  ArrowDown, 
  Search, 
  RefreshCw,
  Sun,
  Moon,
  ChevronRight
} from 'lucide-react';
import { APIResponse, SearchSuggestion, StockData } from './types';
import OrderBook from './components/OrderBook';
import MarketStats from './components/MarketStats';
import OHLCDisplay from './components/OHLCDisplay';

const App: React.FC = () => {
  const [symbol, setSymbol] = useState('HYUNDAI');
  const [data, setData] = useState<StockData | null>(null);
  const [suggestions, setSuggestions] = useState<SearchSuggestion[]>([]);
  const [activeSuggestionIndex, setActiveSuggestionIndex] = useState(-1);
  const [showSuggestions, setShowSuggestions] = useState(false);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isLight, setIsLight] = useState(false);

  const searchContainerRef = useRef<HTMLDivElement>(null);

  const fetchStockPrice = useCallback(async (searchSymbol: string) => {
    if (!searchSymbol) return;
    setLoading(true);
    setError(null);
    try {
      // Connect to the actual go backend endpoint
      const res = await fetch(`/getprice?symbol=${encodeURIComponent(searchSymbol.trim())}`);
      const json: APIResponse & { message?: string } = await res.json();
      
      if (res.ok && json.status === 'success' && json.data) {
        const key = Object.keys(json.data)[0];
        setData(json.data[key]);
      } else {
        throw new Error(json.message || 'Invalid symbol or no data received');
      }
    } catch (err: any) {
      console.error(err);
      setError(err?.message || 'Failed to connect to backend.');
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchStockPrice('HYUNDAI');
  }, [fetchStockPrice]);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (searchContainerRef.current && !searchContainerRef.current.contains(event.target as Node)) {
        setShowSuggestions(false);
      }
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  useEffect(() => {
    const query = symbol.trim();

    if (!query || !showSuggestions) {
      if (!query) {
        setSuggestions([]);
        setActiveSuggestionIndex(-1);
      }
      return;
    }

    const timer = window.setTimeout(async () => {
      try {
        const res = await fetch(`/suggestions?q=${encodeURIComponent(query)}`);
        if (!res.ok) return;

        const json = await res.json();
        if (json.status === 'success' && Array.isArray(json.data)) {
          setSuggestions(json.data as SearchSuggestion[]);
          setActiveSuggestionIndex(-1);
        } else {
          setSuggestions([]);
          setActiveSuggestionIndex(-1);
        }
      } catch (err) {
        console.error(err);
        setSuggestions([]);
        setActiveSuggestionIndex(-1);
      }
    }, 250);

    return () => window.clearTimeout(timer);
  }, [symbol, showSuggestions]);

  useEffect(() => {
    if (isLight) {
      document.body.classList.add('light');
    } else {
      document.body.classList.remove('light');
    }
  }, [isLight]);

  const handleSuggestionClick = (suggestion: SearchSuggestion) => {
    setSymbol(suggestion.symbol);
    setShowSuggestions(false);
    setSuggestions([]);
    setActiveSuggestionIndex(-1);
    fetchStockPrice(suggestion.instrument_key || suggestion.symbol);
  };

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    setShowSuggestions(false);

    if (showSuggestions && suggestions.length > 0 && activeSuggestionIndex >= 0) {
      handleSuggestionClick(suggestions[activeSuggestionIndex]);
      return;
    }

    if (symbol.trim()) {
      setSuggestions([]);
      setActiveSuggestionIndex(-1);
      fetchStockPrice(symbol.trim());
    }
  };

  const handleSearchKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (!showSuggestions || suggestions.length === 0) {
      if (event.key === 'ArrowDown' && suggestions.length > 0) {
        setShowSuggestions(true);
      }
      return;
    }

    if (event.key === 'ArrowDown') {
      event.preventDefault();
      setActiveSuggestionIndex((current) => (current + 1) % suggestions.length);
      return;
    }

    if (event.key === 'ArrowUp') {
      event.preventDefault();
      setActiveSuggestionIndex((current) => (current <= 0 ? suggestions.length - 1 : current - 1));
      return;
    }

    if (event.key === 'Enter' && activeSuggestionIndex >= 0) {
      event.preventDefault();
      handleSuggestionClick(suggestions[activeSuggestionIndex]);
      return;
    }

    if (event.key === 'Escape') {
      event.preventDefault();
      setShowSuggestions(false);
      setActiveSuggestionIndex(-1);
    }
  };

  const isPositive = data ? data.net_change >= 0 : false;
  const percentChange = data ? ((data.net_change / (data.last_price - data.net_change)) * 100).toFixed(2) : '0.00';

  return (
    <div className={`min-h-screen transition-colors duration-500`}>
      <div className="max-w-[1200px] mx-auto px-8 py-10">
        
        {/* Navigation / Header */}
        <nav className="flex items-center justify-between mb-8">
          <div className="flex items-center gap-12">
            {/* Decreased boldness from trace */}
            <h1 className="font-brand text-3xl font-medium tracking-tighter uppercase text-inherit">Trace</h1>
          </div>

          <div className="flex items-center gap-4">
            <button 
              onClick={() => setIsLight(!isLight)}
              className="p-2 rounded-full hover:bg-current/5 transition-colors"
              title="Toggle Theme"
            >
              {isLight ? <Moon size={20} /> : <Sun size={20} />}
            </button>
          </div>
        </nav>

        {/* Centered Search Area above stock price */}
        <div className="flex justify-center mb-16">
          <div ref={searchContainerRef} className="relative w-full max-w-sm">
            <form onSubmit={handleSearch} className="relative group w-full">
              <input
                type="text"
                value={symbol}
                onChange={(e) => {
                  setSymbol(e.target.value);
                  setShowSuggestions(true);
                }}
                onFocus={() => {
                  if (suggestions.length > 0) setShowSuggestions(true);
                }}
                onKeyDown={handleSearchKeyDown}
                className="w-full bg-transparent border-b border-current/10 focus:border-current py-3 text-xs font-mono uppercase tracking-widest focus:outline-none transition-all text-center placeholder:opacity-30"
                placeholder="ENTER SYMBOL OR NAME"
                autoComplete="off"
                spellCheck={false}
              />
              <div className="mt-2 text-[10px] font-mono uppercase tracking-[0.2em] opacity-30 text-center">
                Partial and case-insensitive search supported
              </div>
              <button type="submit" className="absolute right-0 top-1/2 -translate-y-1/2 opacity-40 group-focus-within:opacity-100 transition-opacity p-2">
                <Search className="w-4 h-4" />
              </button>
            </form>

            {showSuggestions && suggestions.length > 0 && (
              <div className="absolute left-0 right-0 top-full mt-4 z-20 overflow-hidden rounded-2xl border border-current/10 bg-[#0d0f14]/95 shadow-2xl backdrop-blur-md">
                {suggestions.map((suggestion, index) => {
                  const isActive = index === activeSuggestionIndex;

                  return (
                    <div
                      key={suggestion.instrument_key}
                      role="option"
                      aria-selected={isActive}
                      tabIndex={-1}
                      onMouseEnter={() => setActiveSuggestionIndex(index)}
                      onMouseDown={(event) => {
                        event.preventDefault();
                        handleSuggestionClick(suggestion);
                      }}
                      className={`flex w-full items-start justify-between gap-4 px-4 py-3 text-left transition-colors cursor-pointer ${isActive ? 'bg-current/10' : 'hover:bg-current/5'}`}
                    >
                      <div className="min-w-0">
                        <div className="text-sm font-mono tracking-widest uppercase">{suggestion.symbol}</div>
                        <div className="truncate text-[10px] opacity-50">{suggestion.name}</div>
                      </div>
                      <div className="flex shrink-0 items-center gap-2 text-[10px] font-mono uppercase tracking-[0.2em] opacity-40">
                        <span>{suggestion.segment}</span>
                        <ChevronRight size={14} />
                      </div>
                    </div>
                  );
                })}
              </div>
            )}
          </div>
        </div>

        {error && (
          <div className="mb-12 border border-rose-500/20 py-3 px-5 text-[10px] font-mono uppercase tracking-[0.2em] text-rose-500 text-center">
            {error}
          </div>
        )}

        {data ? (
          <div className="space-y-20">
            {/* Visual Hero */}
            <section className="flex flex-col items-center text-center gap-10">
              <div className="space-y-2">
                <span className="text-[10px] font-bold tracking-[0.4em] secondary-text uppercase">Equity Instrument</span>
                {/* Updated font to Quicksand and slightly decreased size */}
                <h2 className="font-quicksand text-4xl md:text-6xl lg:text-7xl font-semibold leading-none tracking-tight uppercase">
                  {data.symbol}
                </h2>
              </div>
              
              <div className="flex flex-col items-center gap-4">
                <div className="flex items-center gap-6">
                  {/* Refresh button near the stock price */}
                  <button 
                    onClick={() => fetchStockPrice(symbol)}
                    className={`p-3 rounded-full hover:bg-current/5 transition-colors flex items-center justify-center border border-current/5 ${loading ? 'animate-spin' : ''}`}
                    title="Refresh Data"
                  >
                    <RefreshCw size={20} />
                  </button>
                  <span className="font-serif text-6xl md:text-8xl lg:text-9xl tracking-tighter">
                    {data.last_price.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
                  </span>
                </div>
                <div className={`flex items-center gap-4 font-mono text-sm ${isPositive ? 'text-emerald-500' : 'text-rose-500'}`}>
                  <span className="flex items-center font-bold text-lg">
                    {isPositive ? <ArrowUp size={16} className="mr-1" /> : <ArrowDown size={16} className="mr-1" />}
                    {Math.abs(data.net_change).toFixed(2)}
                  </span>
                  <span className="text-xs opacity-70 px-2 py-0.5 rounded border border-current/20">
                    {isPositive ? '+' : '-'}{Math.abs(parseFloat(percentChange))}%
                  </span>
                </div>
                <div className="flex items-center gap-2 opacity-30 text-[10px] font-mono mt-2 uppercase tracking-widest">
                  <span>ID: {data.instrument_token}</span>
                  <span>•</span>
                  <span>SESSION LIVE</span>
                </div>
              </div>
            </section>

            {/* Content Core */}
            <div className="grid grid-cols-1 lg:grid-cols-12 gap-16 md:gap-24">
              <div className="lg:col-span-8 space-y-20">
                <div className="space-y-8">
                  <header className="flex justify-between items-baseline border-b border-current/10 pb-4">
                    <h3 className="text-[10px] font-bold tracking-[0.5em] uppercase secondary-text">Market Liquidity</h3>
                    <span className="text-[10px] font-mono opacity-30">Order depth analysis</span>
                  </header>
                  <OrderBook depth={data.depth} />
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-16">
                  <MarketStats data={data} />
                </div>
              </div>

              <div className="lg:col-span-4 space-y-16">
                <OHLCDisplay ohlc={data.ohlc} />
                
                <div className="pt-12 border-t border-current/10 space-y-6">
                  <div className="flex justify-between items-center text-[10px] font-bold tracking-[0.2em] secondary-text uppercase">
                    <span>Terminal Logic</span>
                  </div>
                  <div className="space-y-4 font-mono text-[11px] opacity-70">
                    <div className="flex justify-between items-center">
                      <span>Server Status</span>
                      <span className="text-right text-emerald-500 font-bold uppercase">Connected</span>
                    </div>
                    <div className="flex justify-between items-center">
                      <span>Last Fetch</span>
                      <span className="text-right font-bold">{new Date(data.timestamp).toLocaleTimeString()}</span>
                    </div>
                    <div className="flex justify-between items-center">
                      <span>Region</span>
                      <span className="text-right">NSE</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        ) : (
          !loading && (
            <div className="py-40 flex items-center justify-center border border-dashed border-current/10 rounded-lg">
              <p className="font-serif italic text-xl opacity-20">Standby... Awaiting signal trace.</p>
            </div>
          )
        )}

        <footer className="mt-40 pt-10 border-t border-current/5 flex justify-between items-center text-[9px] font-bold tracking-[0.3em] opacity-20 uppercase">
          <span>&copy; Trace {new Date().getFullYear()}</span>
          <span className="hidden md:inline">Universal Trading Interface</span>
          <span>v1.5.0-LIVE</span>
        </footer>
      </div>
    </div>
  );
};

export default App;
