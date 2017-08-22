'use strict';

const gulp = require('gulp');
const gutil = require('gulp-util');

gulp.task('task1', (cb) => {
  gutil.log('task1 start');
  setTimeout(() => {
    gutil.log('task1 end');
    cb();
  }, 1000 * 10);
});

gulp.task('task2', (cb) => {
  gutil.log('task2 start');
  setTimeout(() => {
    gutil.log('task2 end');
    cb();
  }, 1000 * 10);
});

gulp.task('task3', ['task1', 'task2'], (cb) => {
  gutil.log('task3 start');
  setTimeout(() => {
    gutil.log('task3 end');
    cb();
  }, 1000 * 10);
});

gulp.task('default', ['task3']);
