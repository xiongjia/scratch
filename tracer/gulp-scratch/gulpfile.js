'use strict';

const gulp = require('gulp');
const gutil = require('gulp-util');
const posthtml = require('gulp-posthtml');

gulp.task('default', ['task3']);

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

gulp.task('html', () => {
  const plugins = [];
  plugins.push(require('posthtml-doctype')({
    doctype: 'HTML 5'
  }));
  const options = {};
  gulp.src(['./public/**/*.html'])
    .pipe(posthtml(plugins, options))
    .pipe(gulp.dest('dist'));
});
