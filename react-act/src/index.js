import React from "react";
import { render } from "react-dom";



function Cat({ name }) {
 return <h1>The cat's name is {name}</h1>;
}

render(<Cat name="test" />, document.getElementById("root"));
