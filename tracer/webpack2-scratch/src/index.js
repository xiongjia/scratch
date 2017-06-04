'use strict';

require('./assets/style/root.css');

const imgGnu = require('./assets/img/zoro.png');
const root = document.getElementById('root');

root.innerHTML = `
<strong>Webpack2</strong> tests <br>
Production: ${PRODUCTION} <br>
<img src="${imgGnu}">
`;
