'use strict';

import React from 'react';
import { render } from 'react-dom';
import debug from 'debug';
import axios from 'axios';

const dbg = (() => {
  const prefix = 'localtest';
  debug.enable(`${prefix}:*`);
  return debug(`${prefix}:index`);
})();

(() => {
  dbg('start tests');
  axios.get('/api/testdata').then((rep) => {
    dbg('api test data: %j', rep.data);
    dbg('api test status: %d', rep.status);
  }).catch(err => {
    dbg('api test data err: %s', err.toString());
  });
})();

const App = () => {
  return (
    <div>
      <p>test here</p>
    </div>
  );
};

render(<App />, document.getElementById('root'));

