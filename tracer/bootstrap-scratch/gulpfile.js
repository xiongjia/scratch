'use strict';

const argv = require('yargs').argv;
const gulp = require('gulp');
const del = require('del');
const seq = require('gulp-sequence');

const conf = {
  DEBUG: argv.debug
};

const dirs = {
  SRC: 'src',
  SRC_BOOTSTRAP_SASS: 'node_modules/bootstrap-sass',
  DEST: 'dist',
  DEST_CSS: 'dist/css',
  DEST_FONTS: 'dist/fonts'
};

gulp.task('build', [ 'fonts', 'html' ]);
gulp.task('default', seq('clean', 'build'));

gulp.task('clean:all', () => del([ dirs.DEST ]));
gulp.task('clean', [ 'clean:all' ]);

gulp.task('sass', () => {
  const sourcemaps = require('gulp-sourcemaps');
  const sass = require('gulp-sass');
  const gulpif = require('gulp-if');
  const minifyCss = require('gulp-minify-css');
  const rev = require('gulp-rev');

  const sassOpt = {
    outputStyle: 'nested',
    precison: 6,
    errLogToConsole: true,
    includePaths: [ dirs.SRC_BOOTSTRAP_SASS + '/assets/stylesheets' ]
  };

  return gulp.src([ dirs.SRC + '/**/*.scss' ])
    .pipe(gulpif(conf.DEBUG, sourcemaps.init()))
    .pipe(sass(sassOpt))
    .pipe(minifyCss())
    .pipe(gulpif(conf.DEBUG, sourcemaps.write(dirs.DEST_CSS)))
    .pipe(rev())
    .pipe(gulp.dest(dirs.DEST_CSS));
});

gulp.task('html', [ 'sass' ], () => {
  const inject = require('gulp-inject');
  const items = gulp.src([
    dirs.DEST + '/**/*.css',
    dirs.DEST + '/**/*.js'
  ], { read: false });

  return gulp.src([ dirs.SRC + '/**/*.html' ])
    .pipe(inject(items, { ignorePath: dirs.DEST, relative: true }))
    .pipe(gulp.dest(dirs.DEST));
});

gulp.task('fonts', () => {
  const src = [
    dirs.SRC_BOOTSTRAP_SASS + '/assets/fonts/**/*'
  ];
  return gulp.src(src).pipe(gulp.dest(dirs.DEST_FONTS));
});
