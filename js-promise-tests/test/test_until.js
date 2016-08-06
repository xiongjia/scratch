'use strict';

const misc = require('../lib/misc.js');
const logger = misc.logger;

function funcFinal(err, data) {
  if (err) {
    logger(`Final ERR: ${err.toString()}`);
    return;
  }
  logger(`Final: ${data}`);
}

function until(task, data, maxTimes, delay, callback) {
  let retryCnt;

  retryCnt = 0;
  callback = callback || function () {};
  function untilPromise(parentPromise) {
    return parentPromise.catch(() => {
      let retryPromise = new Promise((resolve, reject) => {
        task(data, (err, ret) => {
          if (!err) {
            resolve();
            callback(undefined, ret);
            callback = function () {};
            return;
          }
          retryCnt++;
          logger(`retry -> ${retryCnt}`);
          if (retryCnt >= maxTimes) {
            resolve();
            callback(new Error(`retry times > max times (${maxTimes})`));
            callback = function () {};
            return;
          }

          if (delay > 0) {
            setTimeout(() => { reject(untilPromise(retryPromise)); }, delay);
          }
          else {
            setImmediate(() => { reject(untilPromise(retryPromise)); });
          }
        });
      });
    });
  }
  untilPromise(Promise.reject());
}

exports.untilPassed = (callback) => {
  let retryTask;

  /* - Set the max retry times to 1000, retry delay 10ms.
   * - After 100 calls, the retryTask will be switched to success. 
   */
  callback = callback || function () {};
  const maxTimes = 1000;
  const retryDelay = 10;
  const retryCnt = 100;
  retryTask = misc.mkRetryFunc({
    name: 'A',
    err: 'retryTask Error',
    ret: 'A',
    onRetry: (opts) => opts.retryCnt >= retryCnt
  });
  until(retryTask, 'start', maxTimes,
    retryDelay, (err, data) => { funcFinal(err, data); callback(); });
};

exports.untilErr = (callback) => {
  let retryTask;

  /* - Set the max retry times to 1000, retry delay 10ms.
   * - The retryTask always retrun error. 
   * - This task will be terminated with error after 1000 times retry.
   */
  callback = callback || function () {};
  const maxTimes = 1000;
  const retryDelay = 10;
  retryTask = misc.mkRetryFunc({
    name: 'A',
    err: 'retryTask Error',
    ret: 'A',
    onRetry: (opts) => opts.retryCnt >= Infinity 
  });
  until(retryTask, 'start', maxTimes,
    retryDelay, (err, data) => { funcFinal(err, data); callback(); });
};

(() => {
  if (require.main !== module) {
    return;
  }
  Promise.resolve().then(() => {
    return new Promise((resolve, reject) => exports.untilPassed(resolve));
  }).then(() => {
    return new Promise((resolve, reject) => exports.untilErr(resolve));
  });
})();

