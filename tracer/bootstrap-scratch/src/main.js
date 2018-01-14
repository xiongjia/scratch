'use strict';

import each from 'lodash/each';

each([ 'test1', 'test2' ], (item) => {
  console.log('item: %s', item);
});
