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

exports.allPassed = (callback) => {
  let funcA, funcB, funcC, promise;

  logger('--------------------------------------------------------');
  logger('Task A [Pass] -> Task B [Pass] -> Task C [Pass] -> Final');
  logger('--------------------------------------------------------');
  /* Output of below code snippet:
   *
   *   20:11:14 Begin A: data = start
   *   20:11:15 End A (1000ms)
   *   20:11:15 Begin B: data = A
   *   20:11:16 End B (1000ms)
   *   20:11:16 Begin C: data = B
   *   20:11:17 End C (1000ms)
   *   20:11:17 Final: C
   */
  callback = callback || function () {};
  funcA = misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000 });
  funcB = misc.mkTestFunc({ name: 'B', ret: 'B', delay: 1000 });
  funcC = misc.mkTestFunc({ name: 'C', ret: 'C', delay: 1000 });
  promise = Promise.resolve('start');
  promise = promise.then((data) => misc.mkPromise(funcA, data))
    .then((data) => misc.mkPromise(funcB, data))
    .then((data) => misc.mkPromise(funcC, data));
  promise.then((data) => funcFinal(undefined, data), (err) => funcFinal(err))
    .then(() => callback());
};

exports.waterfall = (callback) => {
  let funcA, funcB, funcC, promise;

  logger('------------------------------------------------');
  logger('Task A [Pass] -> Task B [Err] -X-> Task C [Skip]');
  logger('                    |                           ');
  logger('                    +--> Final [Error]          ');
  logger('------------------------------------------------');
  /* Output of below code snippet:
   *
   *   00:10:05 Begin A: data = start
   *   00:10:06 End A (1000ms)
   *   00:10:06 Begin B: data = A
   *   00:10:07 End B (1001ms)
   *   00:10:07 Final ERR: Error: B Err
   */
  callback = callback || function () {};
  funcA = misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000 });
  funcB = misc.mkTestFunc({ name: 'B', ret: 'B', delay: 1000, err: 'B Err' });
  funcC = misc.mkTestFunc({ name: 'C', ret: 'C', delay: 1000 });
  promise = Promise.resolve('start');
  promise = promise.then((data) => misc.mkPromise(funcA, data))
    .then((data) => misc.mkPromise(funcB, data))
    .then((data) => misc.mkPromise(funcC, data));
  promise.then((data) => funcFinal(undefined, data), (err) => funcFinal(err))
    .then(() => callback());
};

exports.ignoreErr = (callback) => {
  let funcA, funcB, funcC, promise;

  logger('----------------------------------------------------------------');
  logger('Task A [Ignore] -> Task B [Ignore] -X-> Task C [Ignore] -> Final');
  logger('----------------------------------------------------------------');
  /* Output of below code snippet:
   *
   *   21:18:53 Begin A: data = start
   *   21:18:54 End A (1001ms)
   *   21:18:54 Begin B: data = Error: A Err
   *   21:18:55 End B (1002ms)
   *   21:18:55 Begin C: data = Error: B Err
   *   21:18:56 End C (1000ms)
   *   21:18:56 Final: Error: C Err
   */
  callback = callback || function () {};
  funcA = misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000, err: 'A Err' });
  funcB = misc.mkTestFunc({ name: 'B', ret: 'B', delay: 1000, err: 'B Err' });
  funcC = misc.mkTestFunc({ name: 'C', ret: 'C', delay: 1000, err: 'C Err' });
  promise = Promise.resolve('start');
  promise = promise.then((data) => misc.mkPromise(funcA, data))
    .catch((err) => err.toString())
    .then((data) => misc.mkPromise(funcB, data))
    .catch((err) => err.toString())
    .then((data) => misc.mkPromise(funcC, data))
    .catch((err) => err.toString());
  promise.then((data) => funcFinal(undefined, data), (err) => funcFinal(err))
    .then(() => callback());
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

