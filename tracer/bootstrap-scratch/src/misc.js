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

export const genMockChartData = () => {
  const data = {
    labels: ['M', 'T', 'W', 'T', 'F', 'S', 'S'],
    datasets: []
  };
  data.datasets.push({
    label: 'apples',
    data: [12, 19, 3, 17, 6, 3, 7],
    backgroundColor: 'rgba(153,255,51,0.4)'
  });
  data.datasets.push({
    label: 'oranges',
    data: [2, 29, 5, 5, 2, 3, 10],
    backgroundColor: 'rgba(255,153,0,0.4)'
  });
  return { type: 'line', data: data };
};
