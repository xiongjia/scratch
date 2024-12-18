import { EntryPage } from '&/view/index'
import { StyleProvider, px2remTransformer } from '@ant-design/cssinjs'
import { ConfigProvider } from 'antd'
import { Route, Routes } from 'react-router-dom'
import './App.css'

const px2rem = px2remTransformer({ rootValue: 16 })

const App = () => {
  return (
    <ConfigProvider
      form={{
        colon: false,
      }}
    >
      <StyleProvider hashPriority="high" transformers={[px2rem]}>
        <Routes>
          <Route path="/" element={<EntryPage />} />
          <Route path="/test" element={<EntryPage />} />
          <Route path="*" element={<EntryPage />} />
        </Routes>
      </StyleProvider>
    </ConfigProvider>
  )
}

export default App
