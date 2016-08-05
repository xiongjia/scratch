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

exports.race = () => {
  let funcA, funcB;

  logger('----------------------------------------');
  logger('Task A [Pass (3seconds)] -+             ');
  logger('                          |-> Final [B] ');
  logger('Task B [Pass (1seconds)] -+             ');
  logger('----------------------------------------');
  /* Output of below code snippet:
   *
   *   [23:10:50] Begin A: data = start
   *   [23:10:50] Begin B: data = start
   *   [23:10:51] End B (1001ms)
   *   [23:10:51] Final: B
   *   [23:10:53] End A (3001ms)
   */
  funcA = misc.mkTestFunc({ name: 'A', ret: 'A', delay: 3000 });
  funcB = misc.mkTestFunc({ name: 'B', ret: 'B', delay: 1000 });
  Promise.race([
    misc.mkPromise(funcA, 'start'),
    misc.mkPromise(funcB, 'start'),
  ]).then((data) => funcFinal(undefined, data), (err) => funcFinal(err));
};

(() => {
  if (require.main !== module) {
    return;
  }
  exports.race();
})();

