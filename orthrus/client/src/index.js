'use strict';

import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';

import Layout from './components/layout.js';
import * as misc from './misc.js';
import store from './store.js';

const logger = misc.getLogger('index');

logger.debug('orthrus client');
const root = document.getElementById('root');
ReactDOM.render(
  <Provider store={store}>
    <Layout
      data={99}
      desc={'test desc'}
      dispatch={store.dispatch}
    />
  </Provider>, root);
