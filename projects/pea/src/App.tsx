import { ConfigProvider } from 'antd'
import { HashRouter, Route, Routes, Navigate } from 'react-router-dom'
import { Home, Playground } from '&/view'
import './App.css'

const App = () => {
  return (
    <ConfigProvider>
      <HashRouter>
        <Routes>
          <Route path="home" element={<Home />} />
          <Route path="playground" element={<Playground />} />
          <Route path="/" element={<Navigate to="home" />} />
        </Routes>
      </HashRouter>
    </ConfigProvider>
  )
}

export default App
