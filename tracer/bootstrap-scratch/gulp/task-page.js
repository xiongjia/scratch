'use strict';

const gulp = require('gulp');
const path = require('path');
const gulpif = require('gulp-if');

exports = module.exports = (conf, dirs) => {
  gulp.task('index', [ 'sass', 'js', 'assets' ], () => {
    const inject = require('gulp-inject');
    const injectStr = require('gulp-inject-string');
    const htmlbeautify = require('gulp-html-beautify');
    const posthtml = require('gulp-posthtml');
    const htmlmin = require('gulp-htmlmin');
    const include = require('posthtml-include')({
      root: path.join(__dirname, '..', dirs.SRC)
    });
    const exp = require('posthtml-expressions')({
      locals: {
        title: conf.DESC,
        debug: conf.DEBUG,
        buildTM: conf.BUILD_TM,
        buildTS: conf.BUILD_TS,
        buildVer: conf.VER,
        home: { page: 'Home' }
      }
    });

    const items = gulp.src([
      dirs.DEST + '/**/*.css',
      dirs.DEST + '/**/jquery*.js',
      dirs.DEST + '/**/bootstrap*.js',
      dirs.DEST + '/**/bundle*.js'
    ], { read: false });

    return gulp.src([ dirs.SRC + '/index.html' ])
      .pipe(inject(items, { ignorePath: dirs.DEST + '/', relative: false }))
      .pipe(posthtml(() => ({
        plugins: [ include, exp ],
        options: {}
      })))
      .pipe(injectStr.after('</title>', '<!-- string inj1 -->'))
      .pipe(injectStr.after('</title>', '<!-- string inj2 -->'))
      .pipe(gulpif(conf.DEBUG, htmlbeautify({ indentSize: 2 })))
      .pipe(gulpif(!conf.DEBUG, htmlmin({ collapseWhitespace: true })))
      .pipe(gulp.dest(dirs.DEST));
  });
};
