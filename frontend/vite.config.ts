import path from 'path'
import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, '.', '')
  const isDevelopment = mode === 'development'

  return {
    base: isDevelopment ? '/' : '/seeprice/',

    plugins: [react()],

    server: {
      proxy: {
        '/getprice': 'http://localhost:8000',
        '/suggestions': 'http://localhost:8000',
        '/ohlc': 'http://localhost:8000',
        '/login': 'http://localhost:8000',
        '/callback': 'http://localhost:8000',
        '/health': 'http://localhost:8000',
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
