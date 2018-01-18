'use strict';

import debug from 'debug';

export const initDbgLog = (opts) => {
  if (opts.debug) {
    debug.enable('scratch:*');
  } else {
    debug.disable();
  }
};

export const mkDbgLog = (prefix) => debug('scratch:' + prefix);
