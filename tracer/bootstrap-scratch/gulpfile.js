'use strict';

const os = require('os');
const argv = require('yargs').argv;
const gulp = require('gulp');
const gutil = require('gulp-util');
const del = require('del');
const seq = require('gulp-sequence');
const browserSync = require('browser-sync').create();
const path = require('path');

const conf = {
  BROWSER: (() => {
    if (argv.browser) {
      return argv.browser;
    }
    if (os.platform() === 'win32') {
      return 'chrome.exe';
    } else {
      return 'google chrome';
    }
  })(),
  DEBUG: argv.debug
};

const dirs = {
  SRC: 'src',
  SRC_BOOTSTRAP_SASS: 'node_modules/bootstrap-sass',
  DEST: 'dist',
  DEST_CSS: 'dist/css',
  DEST_FONTS: 'dist/fonts'
};

gutil.log('bootstrap scratch');
gutil.log('conf = %j', conf);
gutil.log('dirs = %j', dirs);

gulp.task('build', seq('lint:js', [ 'fonts', 'index' ]));
gulp.task('default', seq(['clean', 'lint:js'], 'build'));

gulp.task('clean:all', () => del([ dirs.DEST ]));
gulp.task('clean', [ 'clean:all' ]);

gulp.task('lint:js', () => {
  const eslint = require('gulp-eslint');
  return gulp.src([ '**/*.js', '!node_modules/**' ])
    .pipe(eslint()).pipe(eslint.format())
    .pipe(eslint.failAfterError());
});

gulp.task('sass', () => {
  const sourcemaps = require('gulp-sourcemaps');
  const sass = require('gulp-sass');
  const gulpif = require('gulp-if');
  const cleanCSS = require('gulp-clean-css');
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
    .pipe(cleanCSS({compatibility: 'ie8'}))
    .pipe(gulpif(conf.DEBUG, sourcemaps.write(dirs.DEST_CSS)))
    .pipe(rev()).pipe(gulp.dest(dirs.DEST_CSS));
});

gulp.task('index', [ 'sass' ], () => {
  const inject = require('gulp-inject');
  const posthtml = require('gulp-posthtml');
  const include = require('posthtml-include')({
    root: path.join(__dirname, dirs.SRC)
  });

  const items = gulp.src([
    dirs.DEST + '/**/*.css',
    dirs.DEST + '/**/*.js'
  ], { read: false });

  return gulp.src([ dirs.SRC + '/index.html' ])
    .pipe(inject(items, { ignorePath: dirs.DEST + '/', relative: false }))
    .pipe(posthtml(() => ({
      plugins: [ include ],
      options: {}
    })))
    .pipe(gulp.dest(dirs.DEST));
});

gulp.task('fonts', () => {
  const src = [
    dirs.SRC_BOOTSTRAP_SASS + '/assets/fonts/**/*'
  ];
  return gulp.src(src).pipe(gulp.dest(dirs.DEST_FONTS));
});

gulp.task('serv', [ 'build' ], () => {
  browserSync.init({
    server: { baseDir: dirs.DEST },
    browser: conf.BROWSER
  });
  gulp.watch(dirs.SRC + '/**/*.html', [ 'html-watch' ]);
  gulp.watch(dirs.SRC + '/**/*.sass', [ 'sass-watch' ]);
});

gulp.task('sass-watch', [ 'html-watch' ]);
gulp.task('html-watch', [ 'index' ], (done) => {
  browserSync.reload();
  done();
});
