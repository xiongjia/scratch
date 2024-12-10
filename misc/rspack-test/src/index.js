import { entry } from './entry';

function render() {
  document.getElementById('root').innerHTML = `the answer to the universe is ${entry}`;
}
render();
