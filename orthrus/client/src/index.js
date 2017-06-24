'use strict';

import React from 'react';
import ReactDOM from 'react-dom';

import Layout from './components/layout.js';
import * as misc from './misc.js';

const logger = misc.getLogger('index');

logger.debug('orthrus client');
const root = document.getElementById('root');
ReactDOM.render(<Layout data={'test'} />, root);
