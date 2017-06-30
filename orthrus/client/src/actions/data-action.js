'use strict';

export function getTestData(testData) {
  return { type: 'GET_TEST_DATA', testData: testData };
}

export function getTestDataErr(err) {
  return { type: 'GET_TEST_DATA_ERR', err: err };
}
