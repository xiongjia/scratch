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
  DEST_CSS_MAP: '.'
};

gulp.task('build', [ 'html' ]);
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
    precison: 3,
    errLogToConsole: true,
    includePaths: [ dirs.SRC_BOOTSTRAP_SASS + '/assets/stylesheets' ]
  };

  return gulp.src([ dirs.SRC + '/**/*.scss' ])
    .pipe(gulpif(conf.DEBUG, sourcemaps.init()))
    .pipe(sass(sassOpt))
    .pipe(minifyCss())
    .pipe(gulpif(conf.DEBUG, sourcemaps.write(dirs.DEST_CSS_MAP)))
    .pipe(rev())
    .pipe(gulp.dest(dirs.DEST));
});

gulp.task('html', [ 'sass' ], () => {
  const inject = require('gulp-inject');
  const items = gulp.src([
    dirs.DEST + '/**/*.css',
    dirs.DEST + '/**/*.js'
  ], { read: false });

  return gulp.src([ dirs.SRC + '/**/*.html' ])
    .pipe(inject(items), { relative: true })
    .pipe(gulp.dest(dirs.DEST));
});

gulp.task('fonts', () => {
  const src = ['node_modules/bootstrap-sass/assets/fonts/**/*'];
  return gulp.src(src)
    .pipe(gulp.dest('dist'));
});
