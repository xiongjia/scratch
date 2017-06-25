'use strict';

import { createStore } from 'redux';

import reducer from './reducers';

export default createStore(reducer);

