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
<img src='${imgGnu}' class='test'>
`;

const elem = document.querySelector('.test');
const draggable = new Draggabilly( elem);
draggable.on('dragEnd', (evt) => {
  logger.debug('dragend', evt);
});
