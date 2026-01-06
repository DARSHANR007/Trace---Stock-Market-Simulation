
import React from 'react';
import { StockData } from '../types';

interface MarketStatsProps {
  data: StockData;
}

const MarketStats: React.FC<MarketStatsProps> = ({ data }) => {
  return (
    <>
      <div className="space-y-10">
        <h4 className="text-[10px] uppercase font-bold tracking-[0.5em] secondary-text border-l-2 border-current pl-4">Volume Distribution</h4>
        <div className="space-y-8">
          <div className="flex flex-col">
            <span className="text-[9px] font-bold secondary-text uppercase tracking-widest mb-2 opacity-60">Traded Accumulation</span>
            <span className="text-3xl font-mono leading-none">{data.volume.toLocaleString()}</span>
          </div>
          <div className="flex flex-col">
            <span className="text-[9px] font-bold secondary-text uppercase tracking-widest mb-2 opacity-60">Mean Traded Price</span>
            <span className="text-3xl font-mono leading-none">₹{data.average_price.toFixed(2)}</span>
          </div>
        </div>
      </div>

      <div className="space-y-10">
        <h4 className="text-[10px] uppercase font-bold tracking-[0.5em] secondary-text border-l-2 border-current pl-4">System Bounds</h4>
        <div className="space-y-8">
          <div className="flex flex-col">
            <span className="text-[9px] font-bold text-emerald-600/80 uppercase tracking-widest mb-2">Cap (Upper)</span>
            <span className="text-3xl font-mono leading-none text-emerald-500 font-bold">₹{data.upper_circuit_limit.toFixed(2)}</span>
          </div>
          <div className="flex flex-col">
            <span className="text-[9px] font-bold text-rose-600/80 uppercase tracking-widest mb-2">Floor (Lower)</span>
            <span className="text-3xl font-mono leading-none text-rose-500 font-bold">₹{data.lower_circuit_limit.toFixed(2)}</span>
          </div>
        </div>
      </div>
    </>
  );
};

export default MarketStats;
