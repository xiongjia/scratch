import React from "react"
import ReactDOM from "react-dom"
import { Provider } from "react-redux"

const app = document.getElementById('app')

ReactDOM.render(<Provider>
  <div><h1>test2</h1></div>
</Provider>, app);
