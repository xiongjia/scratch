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

exports.mkRetryFunc = (opts) => {
  let retryCnt, onRetry;
  retryCnt = 0;
  onRetry = opts.onRetry || function () {};
  return function (data, callback) {
    logger(`Begin ${opts.name}(${retryCnt}): data = ${data}`);
    setTimeout(() => {
      let retryCheck;
      retryCnt++;
      retryCheck = onRetry({ data: data, retryCnt: retryCnt });
      logger(`Retry ${opts.name}[${retryCnt}]: ${retryCheck}`);
      if (retryCheck) {
        logger(`End Retry ${opts.name}; no err`);
        callback(undefined, opts.ret);
      }
      else {
        logger(`End Retry ${opts.name}; err: ${opts.err}`);
        callback(new Error(`${opts.err} (${retryCnt})`));
      }
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

