'use strict';

const misc = require('./misc.js');

exports.create = (ctx) => {
  let mainWnd;

  ctx = ctx || {};
  const logger = ctx.logger;
  const electron = ctx.electron;
  const cfg = ctx.cfg;

  const BrowserWindow = electron.BrowserWindow;
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
