
import React from 'react';
import { MarketDepth } from '../types';

interface OrderBookProps {
  depth: MarketDepth;
}

const OrderBook: React.FC<OrderBookProps> = ({ depth }) => {
  const maxQty = Math.max(...[...depth.buy, ...depth.sell].map(o => o.quantity));
  const totalBuy = depth.buy.reduce((a, b) => a + b.quantity, 0);
  const totalSell = depth.sell.reduce((a, b) => a + b.quantity, 0);

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-16">
      {/* Bid Side */}
      <div className="space-y-6">
        <div className="flex justify-between items-end">
          <span className="text-[11px] font-black text-emerald-500 tracking-[0.2em] uppercase">Bids</span>
          <div className="text-right">
            <div className="text-[8px] opacity-40 uppercase tracking-widest mb-1">Total Demand</div>
            <div className="font-serif text-3xl font-bold leading-none">{totalBuy.toLocaleString()}</div>
          </div>
        </div>
        <div className="space-y-1.5">
          {depth.buy.map((order, i) => (
            <div key={`buy-${i}`} className="relative h-9 flex items-center px-1 group">
              <div 
                className="absolute right-0 top-0 bottom-0 bg-emerald-500/20 group-hover:bg-emerald-500/30 transition-all duration-300 rounded-sm"
                style={{ width: `${(order.quantity / maxQty) * 100}%` }}
              />
              <div className="relative w-full flex justify-between text-[13px] font-mono px-3">
                <span className="font-bold">{order.quantity}</span>
                <span className="text-emerald-500 font-bold">{order.price.toFixed(2)}</span>
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* Ask Side */}
      <div className="space-y-6">
        <div className="flex justify-between items-end">
          <span className="text-[11px] font-black text-rose-500 tracking-[0.2em] uppercase">Asks</span>
          <div className="text-right">
            <div className="text-[8px] opacity-40 uppercase tracking-widest mb-1">Total Supply</div>
            <div className="font-serif text-3xl font-bold leading-none">{totalSell.toLocaleString()}</div>
          </div>
        </div>
        <div className="space-y-1.5">
          {depth.sell.map((order, i) => (
            <div key={`sell-${i}`} className="relative h-9 flex items-center px-1 group">
              <div 
                className="absolute left-0 top-0 bottom-0 bg-rose-500/20 group-hover:bg-rose-500/30 transition-all duration-300 rounded-sm"
                style={{ width: `${(order.quantity / maxQty) * 100}%` }}
              />
              <div className="relative w-full flex justify-between text-[13px] font-mono px-3">
                <span className="text-rose-500 font-bold">{order.price.toFixed(2)}</span>
                <span className="font-bold">{order.quantity}</span>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default OrderBook;
