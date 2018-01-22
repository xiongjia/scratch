'use strict';

import debug from 'debug';
import moment from 'moment';

export const initDbgLog = (opts) => {
  if (opts.debug) {
    debug.enable('scratch:*');
  } else {
    debug.disable();
  }
};

export const mkDbgLog = (prefix) => debug('scratch:' + prefix);

export const updateTooltip = (opts) => {
  opts.element.tooltip('hide')
    .attr(opts.attr || 'title', opts.text || '')
    .tooltip('fixTitle');
};

export const genTooltipTxt = () => moment().format('LTS');
