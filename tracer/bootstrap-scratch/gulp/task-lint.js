'use strict';

const gulp = require('gulp');

exports = module.exports = (conf, dirs) => {
  gulp.task('lint', [ 'lint:js' ]);
  gulp.task('lint:js', () => {
    const eslint = require('gulp-eslint');
    return gulp.src([ '**/*.js', '!node_modules/**', '!dist/**' ])
      .pipe(eslint()).pipe(eslint.format())
      .pipe(eslint.failAfterError());
  });
};
