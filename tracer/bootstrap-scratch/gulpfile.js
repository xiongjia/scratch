'use strict';

const os = require('os');
const argv = require('yargs').argv;
const gulp = require('gulp');
const gutil = require('gulp-util');
const seq = require('gulp-sequence');
const buildTM = new Date();

const conf = {
  BROWSER: (() => {
    if (argv.browser) {
      return argv.browser;
    }
    if (os.platform() === 'win32') {
      return 'chrome.exe';
    } else {
      return 'google chrome';
    }
  })(),
  DEBUG: !!argv.debug,
  NAME: 'bootstrap-scratch',
  DESC: 'bootstrap scratch',
  VER: '0.1',
  DEV_NAME: 'xiong-jia.le',
  DEV_URL: 'https://github.com/xiongjia',
  BUILD_TM: buildTM.toISOString(),
  BUILD_TS: buildTM.valueOf()
};

const dirs = {
  SRC: 'src',
  SRC_BOOTSTRAP_SASS: 'node_modules/bootstrap-sass',
  SRC_JQUERY: 'node_modules/jquery',
  DEST: 'dist',
  DEST_ASSETS: 'dist/assets',
  DEST_CSS: 'dist/css',
  DEST_CSS_MAP: '.',
  DEST_FONTS: 'dist/fonts',
  DEST_JS: 'dist/js',
  DEST_JS_MAP: '.'
};

require('./gulp/task-fonts.js')(conf, dirs);
require('./gulp/task-clean.js')(conf, dirs);
require('./gulp/task-lint.js')(conf, dirs);
require('./gulp/task-assets.js')(conf, dirs);
require('./gulp/task-style.js')(conf, dirs);
require('./gulp/task-script.js')(conf, dirs);
require('./gulp/task-page.js')(conf, dirs);
require('./gulp/task-serv.js')(conf, dirs);

gutil.log('bootstrap scratch');
gutil.log('conf = %j', conf);
gutil.log('dirs = %j', dirs);

gulp.task('build', (cb) => {
  return seq('lint', [ 'fonts', 'index' ])(cb);
});

gulp.task('default', (cb) => seq('clean', 'build')(cb));
