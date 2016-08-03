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
   *   23:12:54 Begin A
   *   23:12:55 End A (1018ms)
   *   23:12:55 Begin B
   *   23:12:56 End B (1001ms)
   *   23:12:56 Begin C
   *   23:12:57 End C (1000ms)
   *   23:12:57 Final
   *
   */
  callback = callback || function () {};
  funcA = misc.mkTestFunc({ name: 'A', delay: 1000, err: null });
  funcB = misc.mkTestFunc({ name: 'B', delay: 1000, err: null });
  funcC = misc.mkTestFunc({ name: 'C', delay: 1000, err: null });
  promise = Promise.resolve();
  promise = promise.then(() => misc.mkPromise(funcA))
    .then(() => misc.mkPromise(funcB))
    .then(() => misc.mkPromise(funcC));
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
   *   23:35:18 Begin A
   *   23:35:19 End A (1001ms)
   *   23:35:19 Begin B
   *   23:35:20 End B (1000ms)
   *   23:35:20 Final ERR: Error: B Error
   *
   */
  callback = callback || function () {};
  funcA = misc.mkTestFunc({ name: 'A', delay: 1000, err: null });
  funcB = misc.mkTestFunc({ name: 'B', delay: 1000, err: 'B Error' });
  funcC = misc.mkTestFunc({ name: 'C', delay: 1000, err: null });
  promise = Promise.resolve();
  promise = promise.then(() => misc.mkPromise(funcA))
    .then(() => misc.mkPromise(funcB))
    .then(() => misc.mkPromise(funcC));
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

