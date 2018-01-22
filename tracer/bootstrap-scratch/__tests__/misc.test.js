'use strict';

import * as misc from '../src/misc.js';

test('tooltip text', () => {
  const tooltip = misc.genTooltipTxt();
  expect(tooltip).not.toBeUndefined();
});
