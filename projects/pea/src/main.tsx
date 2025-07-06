import '@ant-design/v5-patch-for-react-19'
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { HashRouter } from 'react-router-dom'
import './index.css'
import App from './App.tsx'

const root = document.getElementById('root')
if (root) {
  createRoot(root).render(
    <StrictMode>
      <HashRouter basename={'/'}>
        <App />
      </HashRouter>
    </StrictMode>,
  )
}
