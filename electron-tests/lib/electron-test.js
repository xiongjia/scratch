'use strict';

const misc = require('./misc.js');

exports.run = (ctx) => {
  let logger, wndMain;

  logger = misc.mkLogger({ handler: ctx.logger });
  const cfg = ctx.cfg;
  const electron = require('electron');
  const app = electron.app;
  const mainWindow = require('./wnd-main.js');

  app.on('ready', () => {
    logger.debug('electron app ready');
    wndMain = mainWindow.create({
      logger: logger,
      electron: electron,
      cfg: cfg
    });

    if (cfg.enableDevTools) {
      logger.debug('open dev tools');
      wndMain.openDevTools();
    }
  }).on('window-all-closed', () => {
    logger.debug('all windows are closed');
    app.quit();
  });
};
