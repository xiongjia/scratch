'use strict';

const misc = require('../lib/misc.js');
const logger = misc.logger;

function funcFinal(err) {
  if (err) {
    logger(`Final ERR: ${err.toString()}`);
    return;
  }
  logger('Final');
}

exports.simple = (callback) => {
  let funcA, funcB, funcC, promise;

  /* Task A [Pass] -> Task B [Pass] -> Task C [Pass] -> Final
   *
   * Output of below code snippet:
   *
   *   00:10:02 Begin A: data = start
   *   00:10:03 End A (1019ms)
   *   00:10:03 Begin B: data = A
   *   00:10:04 End B (1001ms)
   *   00:10:04 Begin C: data = B
   *   00:10:05 End C (1000ms)
   *   00:10:05 Final
   */
  callback = callback || function () {};
  funcA = misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000 });
  funcB = misc.mkTestFunc({ name: 'B', ret: 'B', delay: 1000 });
  funcC = misc.mkTestFunc({ name: 'C', ret: 'C', delay: 1000 });
  promise = Promise.resolve('start');
  promise = promise.then((data) => misc.mkPromise(funcA, data))
    .then((data) => misc.mkPromise(funcB, data))
    .then((data) => misc.mkPromise(funcC, data));
  promise.then(() => funcFinal(), (err) => funcFinal(err))
    .then(() => { callback(); });
};

exports.error = (callback) => {
  let funcA, funcB, funcC, promise;

  /* Task A [Pass] -> Task B [Err] -X-> Task C [Skip] 
   *                     |
   *                     +--> Final [Error]
   *
   * Output of below code snippet:
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
  promise.then(() => funcFinal(), (err) => funcFinal(err))
    .then(() => { callback(); });
};

(() => {
  Promise.resolve().then(() => {
    return new Promise((resolve, reject) => exports.simple(resolve));
  }).then(() => {
    return new Promise((resolve, reject) => exports.error(resolve));
  });
})();

