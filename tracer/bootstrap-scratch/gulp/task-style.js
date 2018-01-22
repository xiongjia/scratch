'use strict';

const gulp = require('gulp');
const gulpif = require('gulp-if');
const rev = require('gulp-rev');

exports = module.exports = (conf, dirs) => {
  gulp.task('sass', [ 'clean:css' ], () => {
    const sourcemaps = require('gulp-sourcemaps');
    const sass = require('gulp-sass');
    const cleanCSS = require('gulp-clean-css');

    const sassOpt = {
      outputStyle: 'nested',
      precison: 6,
      errLogToConsole: true,
      includePaths: [ dirs.SRC_BOOTSTRAP_SASS + '/assets/stylesheets' ]
    };

    return gulp.src([ dirs.SRC + '/**/*.scss' ])
      .pipe(gulpif(conf.DEBUG, sourcemaps.init()))
      .pipe(sass(sassOpt))
      .pipe(cleanCSS({compatibility: 'ie8'}))
      .pipe(gulpif(conf.DEBUG, sourcemaps.write(dirs.DEST_CSS_MAP, {
        addComment: false
      })))
      .pipe(gulpif(!conf.DEBUG, rev()))
      .pipe(gulp.dest(dirs.DEST_CSS));
  });
};
