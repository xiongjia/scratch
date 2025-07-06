import './App.css'
import { Playground } from '&/view'
import { Route, Routes } from 'react-router-dom'
import './styles/tailwind.css'

function App() {
  return (
    <Routes>
      <Route path="/" element={<Playground />} />
    </Routes>
  )
}

export default App
