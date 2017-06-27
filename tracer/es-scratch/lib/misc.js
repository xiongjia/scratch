'use strict';

import util from 'util';

export function logger(tag = 'default') {
  let handler = console;

  const getLogger = (level) => {
    return function () {
      let msg = util.format.apply(null, arguments);
      (handler[level] || console[level] || console.info)(msg);
    };
  };
  return {
    error: getLogger('error'),
    warn: getLogger('warn'),
    info: getLogger('info'),
    debug: getLogger('debug')
  };
}
