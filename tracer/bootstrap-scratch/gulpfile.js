'use strict';

const argv = require('yargs').argv;
const gulp = require('gulp');

const conf = {
  DEBUG: argv.debug
};

const dirs = {
  DEST: 'dist',
  DEST_CSS_MAP: '.'
};

gulp.task('sass', () => {
  const sourcemaps = require('gulp-sourcemaps');
  const sass = require('gulp-sass');
  const gulpif = require('gulp-if');
  const minifyCss = require('gulp-minify-css');

  const sassOpt = {
    outputStyle: 'nested',
    precison: 3,
    errLogToConsole: true,
    includePaths: ['node_modules/bootstrap-sass/assets/stylesheets']
  };

  return gulp.src(['src/main.scss'])
    .pipe(gulpif(conf.DEBUG, sourcemaps.init()))
    .pipe(sass(sassOpt))
    .pipe(minifyCss())
    .pipe(gulpif(conf.DEBUG, sourcemaps.write(dirs.DEST_CSS_MAP)))
    .pipe(gulp.dest(dirs.DEST));
});

gulp.task('fonts', () => {
  const src = ['node_modules/bootstrap-sass/assets/fonts/**/*'];
  return gulp.src(src)
    .pipe(gulp.dest('dist'));
});
