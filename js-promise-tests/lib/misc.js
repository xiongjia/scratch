'use strict';

let logger = exports.logger = (msg) => {
  console.log(`${(new Date()).toLocaleTimeString()} %s`, msg);
};

exports.mkTestFunc = (opts) => {
  return function (callback) {
    let begin = Date.now();
    opts = opts || {};
    callback = callback || function () {};
    logger(`Begin ${opts.name}`);
    setTimeout(() => {
      logger(`End ${opts.name} (${ Date.now() - begin }ms)`);
      callback(opts.err ? new Error(opts.err) : undefined);
    }, opts.delay);
  };
};

exports.mkPromise = (func) => {
  return new Promise((resolve, reject) => {
    func((err) => {
      if (err) {
        reject(err);
      }
      else {
        resolve();
      }
    });
  });
};
