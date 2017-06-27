'use strict';

import * as misc from '../lib/misc.js';

const log = misc.logger();

function inc(num) {
  return function (target, key, descriptor) {
    const method = descriptor.value;
    descriptor.value = (...args) => {
      args[0] += num;
      return method.apply(target, args);
    };
  };
}

class Test {
  constructor(data = 100) {
    this.init(data);
  }

  @inc(10)
  init(data) {
    this.testData = data;
  }

  toString() {
    return `testData is ${this.testData}`;
  }
}

(() => {
  const test = new Test(9);
  log.debug('test : %s', test.toString());
})();
