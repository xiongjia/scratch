'use strict';

exports.create = (ctx) => {
  let mainWnd;

  ctx = ctx || {};
  const { logger, cfg, electron } = ctx;
  const { BrowserWindow } = electron;
  const wndPage =`file://${cfg.assets}/index.html`;
  const wndSz = { width: 800, height: 600 };

  logger.debug('Creating main windows: %s (%j)', wndPage, wndSz, {});
  mainWnd = new BrowserWindow(wndSz);
  mainWnd.loadURL(wndPage);
  mainWnd.setMenu(null);
  return {
    openDevTools: () => { mainWnd.webContents.openDevTools(); }
  };
};
