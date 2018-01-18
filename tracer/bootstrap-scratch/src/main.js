'use strict';

import * as misc from './misc.js';
import each from 'lodash/each';

/* app conf */
const appConf = {
  debug: process.env.ENV_DEBUG
};

const dbg = misc.mkDbgLog('main');
misc.initDbgLog(appConf);
dbg('app conf: %j', appConf);

each([ 'test1', 'test2' ], (item) => {
  dbg('item: %s', item);
});
