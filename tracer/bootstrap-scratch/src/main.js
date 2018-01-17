'use strict';

import each from 'lodash/each';

const appConf = {
  debug: process.env.ENV_DEBUG
};

console.log('app conf:', appConf);

each([ 'test1', 'test2' ], (item) => {
  console.log('item: %s', item);
});
