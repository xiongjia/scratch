'use strict';

import * as misc from './misc.js';

/* app conf */
const appConf = {
  debug: process.env.ENV_DEBUG
};

const dbg = misc.mkDbgLog('main');
misc.initDbgLog(appConf);
dbg('app conf: %j', appConf);

/* buttons tests */
const btnDefault = $('#btnDefault');
misc.updateTooltip({ element: btnDefault, text: misc.genTooltipTxt() });
btnDefault.on('click', () => {
  misc.updateTooltip({ element: btnDefault, text: misc.genTooltipTxt() });
  dbg('btn clicked (default)');
  alert('btn clicked (default)');
});

/* dropdown menu in tabcontent-testtab.html */
const btnPick = $('#btnPick');
$('#dropdownMenuItems li a').on('click', function () {
  const item = $(this).text();
  dbg('menu item selected: %s', item);
  btnPick.text(item);
});
