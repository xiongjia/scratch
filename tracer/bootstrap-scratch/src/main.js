'use strict';

import * as misc from './misc.js';
import each from 'lodash/each';

const dbg = misc.mkDbgLog('main');

const appConf = {
  debug: process.env.ENV_DEBUG
};

misc.initDbgLog(appConf);
dbg('app conf: %j', appConf);

each([ 'test1', 'test2' ], (item) => {
  dbg('item: %s', item);
});
