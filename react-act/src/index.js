import React from "react"
import { render } from "react-dom"
import Button from '@mui/material/Button';


function Cat({ name }) {
    return (
        <div>
          <Button variant="contained">Hello {name}</Button>
        </div>
      );
}

render(<Cat name="test" />, document.getElementById("root"))
