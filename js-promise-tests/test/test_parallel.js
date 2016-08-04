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
  let funcA, funcB, funcC;

  logger('-------------------------');
  logger('Task A [Pass] -+         ');
  logger('Task B [Pass] -|-> Final ');
  logger('Task C [Pass] -+         ');
  logger('-------------------------');
  /* Output of below code snippet:
   *
   * 21:46:35 Begin A: data = start
   * 21:46:35 Begin B: data = start
   * 21:46:35 Begin C: data = start
   * 21:46:36 End A (1001ms)
   * 21:46:37 End C (2001ms)
   * 21:46:38 End B (3000ms)
   * 21:46:38 Final: A,B,C
   *
   */
  callback = callback || function () {};
  funcA = misc.mkTestFunc({ name: 'A', ret: 'A', delay: 1000 });
  funcB = misc.mkTestFunc({ name: 'B', ret: 'B', delay: 3000 });
  funcC = misc.mkTestFunc({ name: 'C', ret: 'C', delay: 2000 });
  Promise.all([
    misc.mkPromise(funcA, 'start'),
    misc.mkPromise(funcB, 'start'),
    misc.mkPromise(funcC, 'start')
  ]).then((data) => funcFinal(undefined, data), (err) => funcFinal(err))
    .then(()=> callback());
};

(() => {
  if (require.main !== module) {
    return;
  }
  exports.allPassed();
})();

