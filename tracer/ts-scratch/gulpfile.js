'use strict';

const gulp = require('gulp');
const ts = require('gulp-typescript');
const tslint = require('gulp-tslint');

const tsProject = ts.createProject('tsconfig.json');

gulp.task('compile', ['tslint'], () => tsProject.src()
          .pipe(tsProject()).js.pipe(gulp.dest('dist')));

gulp.task('tslint', () => tsProject.src()
          .pipe(tslint({
            configuration: './tslint.json',
            formatter: 'prose'
          })).pipe(tslint.report({ emitError: true })));

gulp.task('default', ['compile']);
