import React from 'react'
import { render } from 'react-dom'
import { Provider } from 'react-redux'

render(
  <Provider>
    <div><h1>test</h1></div>
  </Provider>,
  document.getElementById('root'))
