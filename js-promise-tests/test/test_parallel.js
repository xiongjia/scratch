'use strict';

const misc = require('../lib/misc.js');
const logger = misc.logger;

function funcFinal(data) {
  logger(`Final: ${data}`);
}

function parallel(tasks, callback) {
  let promiseTasks;

  function mkParallelPromise(task) {
    return misc.mkPromise(task, 'start')
      .then((val) => val, (err) => err.toString());
  }
  promiseTasks = [];
  for (const task of tasks) {
    promiseTasks.push(mkParallelPromise(task));
  }
  Promise.all(promiseTasks)
    .then((data) => funcFinal(data), (ignore) => {})
    .then(()=> callback());
}

exports.allPassed = (callback) => {
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
  parallel([
    misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000 }),
    misc.mkTestFunc({ name: 'B', ret: 'B', delay: 3000 }),
    misc.mkTestFunc({ name: 'C', ret: 'C', delay: 2000 })
  ], callback);
};

exports.taskCErr = (callback) => {
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
  parallel([
    misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000 }),
    misc.mkTestFunc({ name: 'B', ret: 'B', delay: 3000 }),
    misc.mkTestFunc({ name: 'C', ret: 'C', delay: 2000, err: 'C Err' })
  ], callback);
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

