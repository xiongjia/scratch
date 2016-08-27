'use strict';

const _ = require('lodash');

exports.mkLogger = (opts) => {
  let handler, log;

  const LEVELS = ['error', 'info', 'warn', 'debug'];
  opts = opts || {};
  handler = opts.handler || {};
  log = {
    disable: () => {
      _.each(LEVELS, (level) => {
        log[level] = function () {};
      });
      return log;
    },
    enable: () => {
      _.each(LEVELS, (level) => {
        log[level] = handler[level] || function () {};
      });
      return log;
    }
  };
  return log.enable();
};
