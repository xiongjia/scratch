'use strict';

const gulp = require('gulp');
const sass = require('gulp-sass');
const inject = require('gulp-inject');
const wiredep = require('wiredep').stream;

gulp.task('clean', (cb) => require('del')(['dist'], cb));

gulp.task('styles', () => {
  const injFiles = gulp.src('src/styles/*.scss', { read: false });
  const injOpts = {
    transform: (filepath) => ('@import "' + filepath + '";'),
    addRootSlash: false,
    starttag: '// inject:app',
    endtag: '// endinject'
  };

  return gulp.src('src/main.scss')
    .pipe(wiredep())
    .pipe(inject(injFiles, injOpts))
    .pipe(sass())
    .pipe(gulp.dest('dist'));
});

gulp.task('html', ['styles'], () => {
  const injectFiles = gulp.src(['dist/main.css']);
  const injectOptions = {
    addRootSlash: false,
    ignorePath: ['src', 'dist']
  };

  return gulp.src('src/index.html')
    .pipe(inject(injectFiles, injectOptions))
    .pipe(gulp.dest('dist'));
});

