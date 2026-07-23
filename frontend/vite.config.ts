import path from 'path'
import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, '.', '')
  const isDevelopment = mode === 'development'
  const proxyTarget = env.VITE_API_PROXY_TARGET || 'http://localhost:8000'

  return {
    base: isDevelopment ? '/' : '/seeprice/',

    plugins: [react()],

    server: {
      proxy: {
        '/getprice': proxyTarget,
        '/suggestions': proxyTarget,
        '/ohlc': proxyTarget,
        '/login': proxyTarget,
        '/callback': proxyTarget,
        '/health': proxyTarget,
      },
    },

    define: {
      'process.env.GEMINI_API_KEY': JSON.stringify(env.GEMINI_API_KEY),
    },

    resolve: {
      alias: {
        '@': path.resolve(__dirname, '.'),
      },
    },
  }
})
