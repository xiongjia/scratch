'use strict';

import * as misc from './misc.js';

/* app conf */
const appConf = {
  debug: process.env.ENV_DEBUG
};

const dbg = misc.mkDbgLog('main');
misc.initDbgLog(appConf);
dbg('app conf: %j', appConf);

$('#btnDefault').on('click', () => {
  dbg('btn clicked (default)');
  alert('test1');
});
