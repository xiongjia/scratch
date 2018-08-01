'use strict';

exports.logger = (msg) => {
  console.log(`[${(new Date()).toLocaleTimeString()}] %s`, msg);
};


