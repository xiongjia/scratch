'use strict';

((args) => {
  const path = require('path');
  const winston = require('winston');
  winston.cli();
  winston['default'].transports.console.level = 'silly';
  winston['default'].transports.console.colorize = false;
  winston.info('Creating electronTest, versions: %j', process.versions, {});
  const electronTest = require('../');
  electronTest.run({
    logger: winston,
    cfg: {
      assets: path.join(__dirname, '../assets'),
      enableDevTools: args.enableDevTools
    }
  });
})(require('yargs').argv);
