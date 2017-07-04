'use strict';

const gulp = require('gulp');
const ts = require('gulp-typescript');
const tslint = require('gulp-tslint');
const eslint = require('gulp-eslint');

const tsProject = ts.createProject('tsconfig.json');

gulp.task('compile', ['tslint', 'eslint'], () => tsProject.src()
  .pipe(tsProject()).js.pipe(gulp.dest('dist')));

gulp.task('tslint', () => tsProject.src()
  .pipe(tslint({
    configuration: './tslint.json',
    formatter: 'prose'
  })).pipe(tslint.report({ emitError: true })));

gulp.task('eslint', () => gulp.src(['gulpfile.js'])
  .pipe(eslint('.eslintrc.json'))
  .pipe(eslint.format())
  .pipe(eslint.failAfterError()));

gulp.task('default', ['compile']);
