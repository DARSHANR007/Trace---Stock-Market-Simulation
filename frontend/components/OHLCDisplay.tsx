
import React from 'react';
import { OHLC } from '../types';

interface OHLCDisplayProps {
  ohlc: OHLC;
}

const OHLCDisplay: React.FC<OHLCDisplayProps> = ({ ohlc }) => {
  const range = ohlc.high - ohlc.low;
  
  return (
    <div className="space-y-10">
      <h3 className="text-[10px] uppercase font-bold tracking-[0.5em] secondary-text border-l-2 border-current pl-4">Session Precision</h3>
      
      <div className="grid grid-cols-2 gap-x-12 gap-y-12">
        <div className="flex flex-col gap-1">
          <span className="text-[8px] font-bold secondary-text uppercase tracking-widest opacity-60">Open</span>
          <span className="text-xl font-mono">{ohlc.open.toFixed(2)}</span>
        </div>
        <div className="flex flex-col gap-1 text-right">
          <span className="text-[8px] font-bold secondary-text uppercase tracking-widest opacity-60">Close</span>
          <span className="text-xl font-mono">{ohlc.close.toFixed(2)}</span>
        </div>
        <div className="flex flex-col gap-1">
          <span className="text-[8px] font-bold text-emerald-600/80 uppercase tracking-widest">High</span>
          <span className="text-xl font-mono text-emerald-500">{ohlc.high.toFixed(2)}</span>
        </div>
        <div className="flex flex-col gap-1 text-right">
          <span className="text-[8px] font-bold text-rose-600/80 uppercase tracking-widest">Low</span>
          <span className="text-xl font-mono text-rose-500">{ohlc.low.toFixed(2)}</span>
        </div>
      </div>

      <div className="space-y-4 pt-6">
        <div className="flex justify-between items-baseline text-[10px] font-mono uppercase">
          <span className="secondary-text tracking-widest font-bold opacity-50">Spread Range</span>
          <span className="font-bold opacity-80">{range.toFixed(2)}</span>
        </div>
        <div className="relative h-4 w-full bg-current/5 rounded-full p-1 flex items-center overflow-hidden">
           <div className="absolute inset-x-0 h-[1px] bg-current opacity-10"></div>
           <div className="h-full bg-gradient-to-r from-rose-500/40 via-current/10 to-emerald-500/40 w-full rounded-full"></div>
           <div className="absolute left-1/2 -translate-x-1/2 h-full w-[2px] bg-current opacity-20"></div>
        </div>
      </div>
    </div>
  );
};

export default OHLCDisplay;
