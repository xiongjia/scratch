'use strict';

import * as misc from './misc.js';
import * as axios from 'axios';

const logger = misc.getLogger('index');

const root = document.getElementById('root');
root.innerHTML = 'hello';

logger.debug('orthrus client dbg log');

/* axios tests */
axios.get('/api/test')
  .then(res => logger.debug('res: ', res))
  .catch(err => logger.debug('err: ', err));
