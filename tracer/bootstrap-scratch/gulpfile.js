'use strict';

const gulp = require('gulp');

const config = {
  debug: true
};


gulp.task('sass', () => {
  const sass = require('gulp-sass');
  const sassOpt = {
    outputStyle: 'nested',
    precison: 3,
    errLogToConsole: true,
    includePaths: ['node_modules/bootstrap-sass/assets/stylesheets']
  };
  return gulp.src(['src/main.scss'])
    .pipe(sass(sassOpt))
    .pipe(gulp.dest('dist'));
});

gulp.task('fonts', () => {
  const src = ['node_modules/bootstrap-sass/assets/fonts/**/*'];
  return gulp.src(src)
    .pipe(gulp.dest('dist'));
});
