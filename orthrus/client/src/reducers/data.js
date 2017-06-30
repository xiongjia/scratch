'use strict';

import * as misc from '../misc.js';
const logger = misc.getLogger('data');

const defaultState = { data: 'testdata', num: 100};

export default function reducer(state = defaultState, action) {
  logger.debug('action: %s', action.type);

  if (action.type === 'GET_TEST_DATA') {
    logger.debug('get test data: ', action.testData);
    const { data, num } = action.testData;
    return {...state, desc: `get test data: ${data}, ${num}`};
  } else if (action.type === 'GET_TEST_DATA_ERR') {
    logger.debug('get test err: ', action.err);
    return {...state, desc: `get test err, ${action.err}`};
  }
  return state;
}
