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

function retryFixedTimes(maxTimes, retryTask, callback) {
  let promise;

  function mkRetryPromise(task, taskId) {
    return (err) => misc.mkPromise(task,
        `TaskId: ${taskId}; ${err.toString()}`);
  }

  callback = callback || function () {};
  promise = Promise.reject('start');
  for (let i = 0; i < maxTimes; ++i) {
    promise = promise.catch(mkRetryPromise(retryTask, i));
  }
  promise.then((data) => funcFinal(undefined, data), (err) => funcFinal(err))
    .then(() => callback());
}

exports.fixedTimesPass = (callback) => {
  let retryTask;

  logger('--------------------------------------------------------');
  logger(' A [Err; 1st] -> A [Err; 2nd] -> A [Pass; 3rd] -> Final ');
  logger('--------------------------------------------------------');
  /* Output of below code snippet:
   *
   *   [10:26:53] Begin A(0): data = TaskId: 0; start
   *   [10:26:54] Retry A[1]: false
   *   [10:26:54] End Retry A; err: retryTask Error
   *   [10:26:54] Begin A(1): data = TaskId: 1; Error: retryTask Error (1)
   *   [10:26:55] Retry A[2]: false
   *   [10:26:55] End Retry A; err: retryTask Error
   *   [10:26:55] Begin A(2): data = TaskId: 2; Error: retryTask Error (2)
   *   [10:26:56] Retry A[3]: true
   *   [10:26:56] End Retry A; no err
   *   [10:26:56] Final: A
   */
  retryTask = misc.mkRetryFunc({
    name: 'A',
    err: 'retryTask Error',
    ret: 'A',
    delay: 1000,
    onRetry: (opts) => opts.retryCnt >= 3
  });
  retryFixedTimes(5, retryTask, callback);
};

exports.fixedTimesErr = (callback) => {
  let retryTask;

  logger('-------------------------------------------------------------');
  logger(' A [Err; 1st] -> A [Err; 2nd] -> A [Err; 3rd] -> Final [Err] ');
  logger('-------------------------------------------------------------');
  /* Output of below code snippet:
   *
   *   [10:26:56] Begin A(0): data = TaskId: 0; start
   *   [10:26:57] Retry A[1]: false
   *   [10:26:57] End Retry A; err: retryTask Error
   *   [10:26:57] Begin A(1): data = TaskId: 1; Error: retryTask Error (1)
   *   [10:26:58] Retry A[2]: false
   *   [10:26:58] End Retry A; err: retryTask Error
   *   [10:26:58] Begin A(2): data = TaskId: 2; Error: retryTask Error (2)
   *   [10:26:59] Retry A[3]: false
   *   [10:26:59] End Retry A; err: retryTask Error
   *   [10:26:59] Final ERR: Error: retryTask Error (3)
   */
  callback = callback || function () {};
  retryTask = misc.mkRetryFunc({
    name: 'A',
    err: 'retryTask Error',
    ret: 'A',
    delay: 1000,
    onRetry: (opts) => opts.retryCnt >= Infinity
  });
  retryFixedTimes(3, retryTask, callback);
};

(() => {
  if (require.main !== module) {
    return;
  }
  Promise.resolve().then(() => {
    return new Promise((resolve, reject) => exports.fixedTimesPass(resolve));
  }).then(() => {
    return new Promise((resolve, reject) => exports.fixedTimesErr(resolve));
  });
})();

