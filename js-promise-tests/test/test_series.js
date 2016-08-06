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

function series(tasks, ignoreErr, callback) {
  let promise;

  callback = callback || function () {};
  function mkSeriesPromise(task) {
    return (data) => misc.mkPromise(task, data);
  }
  promise = Promise.resolve('start');
  for (const task of tasks) {
    promise = promise.then(mkSeriesPromise(task));
    if (ignoreErr) {
      promise = promise.catch((err) => err.toString());
    }
  }
  promise.then((data) => funcFinal(undefined, data), (err) => funcFinal(err))
    .then(() => callback());
}

exports.allPassed = (callback) => {
  logger('--------------------------------------------------------');
  logger('Task A [Pass] -> Task B [Pass] -> Task C [Pass] -> Final');
  logger('--------------------------------------------------------');
  /* Output of below code snippet:
   *
   *   [20:11:14] Begin A: data = start
   *   [20:11:15] End A (1000ms)
   *   [20:11:15] Begin B: data = A
   *   [20:11:16] End B (1000ms)
   *   [20:11:16] Begin C: data = B
   *   [20:11:17] End C (1000ms)
   *   [20:11:17] Final: C
   */
  series([
    misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000 }),
    misc.mkTestFunc({ name: 'B', ret: 'B', delay: 1000 }),
    misc.mkTestFunc({ name: 'C', ret: 'C', delay: 1000 })
  ], false, callback);
};

exports.waterfall = (callback) => {
  logger('------------------------------------------------');
  logger('Task A [Pass] -> Task B [Err] -X-> Task C [Skip]');
  logger('                    |                           ');
  logger('                    +--> Final [Error]          ');
  logger('------------------------------------------------');
  /* Output of below code snippet:
   *
   *   [00:10:05] Begin A: data = start
   *   [00:10:06] End A (1000ms)
   *   [00:10:06] Begin B: data = A
   *   [00:10:07] End B (1001ms)
   *   [00:10:07] Final ERR: Error: B Err
   */
  series([
    misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000 }),
    misc.mkTestFunc({ name: 'B', ret: 'B', delay: 1000, err: 'B Err' }),
    misc.mkTestFunc({ name: 'C', ret: 'C', delay: 1000 })
  ], false, callback);
};

exports.ignoreErr = (callback) => {
  logger('----------------------------------------------------------------');
  logger('Task A [Ignore] -> Task B [Ignore] -X-> Task C [Ignore] -> Final');
  logger('----------------------------------------------------------------');
  /* Output of below code snippet:
   *
   *   [21:18:53] Begin A: data = start
   *   [21:18:54] End A (1001ms)
   *   [21:18:54] Begin B: data = Error: A Err
   *   [21:18:55] End B (1002ms)
   *   [21:18:55] Begin C: data = Error: B Err
   *   [21:18:56] End C (1000ms)
   *   [21:18:56] Final: Error: C Err
   */
  series([
    misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000, err: 'A Err' }),
    misc.mkTestFunc({ name: 'B', ret: 'B', delay: 1000, err: 'B Err' }),
    misc.mkTestFunc({ name: 'C', ret: 'C', delay: 1000, err: 'C Err' })
  ], true, callback);
};

(() => {
  if (require.main !== module) {
    return;
  }
  Promise.resolve().then(() => {
    return new Promise((resolve, reject) => exports.allPassed(resolve));
  }).then(() => {
    return new Promise((resolve, reject) => exports.waterfall(resolve));
  }).then(() => {
    return new Promise((resolve, reject) => exports.ignoreErr(resolve));
  });
})();

