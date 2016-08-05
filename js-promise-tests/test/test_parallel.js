'use strict';

const misc = require('../lib/misc.js');
const logger = misc.logger;

function funcFinal(data) {
  logger(`Final: ${data}`);
}

exports.allPassed = (callback) => {
  let funcA, funcB, funcC;

  logger('--------------------------------------');
  logger('Task A [Pass] -+                      ');
  logger('Task B [Pass] -|-> Final [A,B,C Pass] ');
  logger('Task C [Pass] -+                      ');
  logger('--------------------------------------');
  /* Output of below code snippet:
   *
   *   [23:23:40] Begin A: data = start
   *   [23:23:40] Begin B: data = start
   *   [23:23:40] Begin C: data = start
   *   [23:23:41] End A (1001ms)
   *   [23:23:42] End C (2001ms)
   *   [23:23:43] End B (3002ms)
   *   [23:23:43] Final: A,B,C
   */
  callback = callback || function () {};
  funcA = misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000 });
  funcB = misc.mkTestFunc({ name: 'B', ret: 'B', delay: 3000 });
  funcC = misc.mkTestFunc({ name: 'C', ret: 'C', delay: 2000 });
  Promise.all([
    misc.mkPromise(funcA, 'start')
      .then((val) => val, (err) => err.toString()),
    misc.mkPromise(funcB, 'start')
      .then((val) => val, (err) => err.toString()),
    misc.mkPromise(funcC, 'start')
      .then((val) => val, (err) => err.toString())
  ]).then((data) => funcFinal(data), (ignore) => {})
    .then(()=> callback());
};

exports.taskCErr = (callback) => {
  let funcA, funcB, funcC;

  logger('---------------------------------------------');
  logger('Task A [Pass] -+                             ');
  logger('Task B [Pass] -|-> Final [A,B Pass; C Error] ');
  logger('Task C [Err ] -+                             ');
  logger('---------------------------------------------');
  /* Output of below code snippet:
   *
   *   [23:23:37] Begin A: data = start
   *   [23:23:37] Begin B: data = start
   *   [23:23:37] Begin C: data = start
   *   [23:23:38] End A (1002ms)
   *   [23:23:39] End C (2001ms)
   *   [23:23:40] End B (3001ms)
   *   [23:23:40] Final: A,B,Error: C Err
   */
  callback = callback || function () {};
  funcA = misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000 });
  funcB = misc.mkTestFunc({ name: 'B', ret: 'B', delay: 3000 });
  funcC = misc.mkTestFunc({ name: 'C', ret: 'C', delay: 2000, err: 'C Err' });
  Promise.all([
    misc.mkPromise(funcA, 'start')
      .then((val) => val, (err) => err.toString()),
    misc.mkPromise(funcB, 'start')
      .then((val) => val, (err) => err.toString()),
    misc.mkPromise(funcC, 'start')
      .then((val) => val, (err) => err.toString())
  ]).then((data) => funcFinal(data), (ignore) => {})
    .then(()=> callback());
};

(() => {
  if (require.main !== module) {
    return;
  }
  Promise.resolve().then(() => {
    return new Promise((resolve, reject) => exports.taskCErr(resolve));
  }).then(() => {
    return new Promise((resolve, reject) => exports.allPassed(resolve));
  });
})();

