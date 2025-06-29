import { resolve } from 'node:path'
import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'
import biomePlugin from 'vite-plugin-biome'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), biomePlugin()],
  resolve: {
    alias: [{ find: '&', replacement: resolve(__dirname, './src') }],
  },
})
