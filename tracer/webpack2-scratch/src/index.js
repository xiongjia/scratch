'use strict';

import * as log from 'loglevel';
import Draggabilly from 'Draggabilly';

log.setLevel(APP_LOGLEVEL);
const logger = log.getLogger('index');

require('./assets/style/root.css');

const imgGnu = require('./assets/img/zoro.png');
const root = document.getElementById('root');

logger.debug('app started');

root.innerHTML = `
<strong>Webpack2</strong> tests <br>
Production: ${APP_PRODUCTION} <br>

begin<br>

<p id="t1"></p>

end<br>

<img src='${imgGnu}' class='test'>
`;

const elem = document.querySelector('.test');
const draggable = new Draggabilly( elem);
draggable.on('dragEnd', (evt) => {
  logger.debug('dragend', evt);
});


const testElem1 = document.querySelector('#t1');
testElem1.className = 'testClassName';
testElem1.appendChild(document.createTextNode('abc'));

import data from './assets/json/data.json';
logger.debug('data', data);

// import * as axios from 'axios';
// axios tests
// axios.get('http://api.fanfou.com/statuses/home_timeline.json')
//   .then(res => logger.debug(res))
//   .catch(err => logger.debug(err));
