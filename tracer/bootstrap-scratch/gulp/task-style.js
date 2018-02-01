'use strict';

const gulp = require('gulp');
const gulpif = require('gulp-if');
const rev = require('gulp-rev');

const uncssBootstrapIgore = [
  /\.affix/,
  /\.alert/,
  /\.close/,
  /\.collaps/,
  /\.fade/,
  /\.has/,
  /\.help/,
  /\.in/,
  /\.modal/,
  /\.open/,
  /\.popover/,
  /\.tooltip/
];

exports = module.exports = (conf, dirs) => {
  gulp.task('sass', [ 'clean:css' ], () => {
    const sourcemaps = require('gulp-sourcemaps');
    const sass = require('gulp-sass');
    const cleanCSS = require('gulp-clean-css');
    const postcss = require('gulp-postcss');
    const uncss = require('postcss-uncss');

    const sassOpt = {
      outputStyle: 'nested',
      precison: 6,
      errLogToConsole: true,
      includePaths: [ dirs.SRC_BOOTSTRAP_SASS + '/assets/stylesheets' ]
    };

    const postcssPlugin = [
      uncss({
        html: [ dirs.SRC + '/**/*.html' ],
        ignore: [ ...uncssBootstrapIgore ]
      }),
    ];

    return gulp.src([ dirs.SRC + '/**/*.scss' ])
      .pipe(gulpif(conf.DEBUG, sourcemaps.init()))
      .pipe(sass(sassOpt))
      .pipe(postcss(postcssPlugin))
      .pipe(gulpif(!conf.DEBUG, cleanCSS({
        level: { 1: {specialComments: 0} }
      })))
      .pipe(gulpif(conf.DEBUG, sourcemaps.write(dirs.DEST_CSS_MAP, {
        addComment: false
      })))
      .pipe(gulpif(!conf.DEBUG, rev()))
      .pipe(gulp.dest(dirs.DEST_CSS));
  });
};
