'use strict';

import * as log from 'loglevel';
log.setLevel(APP_LOGLEVEL);

export function getLogger(tag) {
  return log.getLogger(tag || 'orthrus');
}
