'use strict';

let logger = exports.logger = (msg) => {
  console.log(`[${(new Date()).toLocaleTimeString()}] %s`, msg);
};

exports.mkTestFunc = (opts) => {
  return function (data, callback) {
    let begin = Date.now();

    opts = opts || {};
    callback = callback || function () {};

    logger(`Begin ${opts.name}: data = ${data}`);
    setTimeout(() => {
      logger(`End ${opts.name} (${ Date.now() - begin }ms)`);
      callback(opts.err ? new Error(opts.err) : undefined, opts.ret);
    }, opts.delay);
  };
};

exports.mkPromise = (func, data) => {
  return new Promise((resolve, reject) => {
    func(data, (err, ret) => {
      if (err) {
        reject(err);
      }
      else {
        resolve(ret);
      }
    });
  });
};
