'use strict';

export default function reducer(state = 100, action) {
  if (action.type === 'INC') {
    return state + 1;
  } else if (action.type === 'DEC') {
    return state - 1;
  }
  return state;
}
