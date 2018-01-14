'use strict';

const os = require('os');
const argv = require('yargs').argv;
const gulp = require('gulp');
const gutil = require('gulp-util');
const gulpif = require('gulp-if');
const del = require('del');
const seq = require('gulp-sequence');
const browserSync = require('browser-sync').create();
const path = require('path');
const rev = require('gulp-rev');

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
  DEST_CSS_MAP: '.',
  DEST_FONTS: 'dist/fonts',
  DEST_JS: 'dist/js',
  DEST_JS_MAP: '.'
};

gutil.log('bootstrap scratch');
gutil.log('conf = %j', conf);
gutil.log('dirs = %j', dirs);

gulp.task('build', seq('lint:js', [ 'fonts', 'index' ]));
gulp.task('default', seq('clean', 'lint:js', 'build'));

gulp.task('clean:all', () => del([ dirs.DEST ]));
gulp.task('clean:js', () => del([ dirs.DEST + '/**/*.js' ]));
gulp.task('clean:css', () => del([ dirs.DEST + '/**/*.css' ]));
gulp.task('clean', [ 'clean:all' ]);

gulp.task('lint:js', () => {
  const eslint = require('gulp-eslint');
  return gulp.src([ '**/*.js', '!node_modules/**', '!dist/**' ])
    .pipe(eslint()).pipe(eslint.format())
    .pipe(eslint.failAfterError());
});

gulp.task('sass', [ 'clean:css' ], () => {
  const sourcemaps = require('gulp-sourcemaps');
  const sass = require('gulp-sass');
  const cleanCSS = require('gulp-clean-css');

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
    .pipe(gulpif(conf.DEBUG, sourcemaps.write(dirs.DEST_CSS_MAP)))
    .pipe(rev()).pipe(gulp.dest(dirs.DEST_CSS));
});

gulp.task('index', [ 'sass', 'js' ], () => {
  const inject = require('gulp-inject');
  const htmlbeautify = require('gulp-html-beautify');
  const posthtml = require('gulp-posthtml');
  const htmlmin = require('gulp-htmlmin');
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
    .pipe(gulpif(conf.DEBUG, htmlbeautify({ indentSize: 2 })))
    .pipe(gulpif(!conf.DEBUG, htmlmin({ collapseWhitespace: true })))
    .pipe(gulp.dest(dirs.DEST));
});

gulp.task('fonts', () => {
  const src = [
    dirs.SRC_BOOTSTRAP_SASS + '/assets/fonts/**/*'
  ];
  return gulp.src(src).pipe(gulp.dest(dirs.DEST_FONTS));
});

gulp.task('js', [ 'clean:js' ], () => {
  const sourcemaps = require('gulp-sourcemaps');
  const source = require('vinyl-source-stream');
  const buffer = require('vinyl-buffer');
  const browserify = require('browserify');
  const uglify = require('gulp-uglify');

  const ent = {
    entries: [ 'src/main.js'],
    debug: true
  };

  return browserify(ent)
    .transform('babelify', { presets: [ 'env' ] })
    .bundle()
    .pipe(source('bundle.js'))
    .pipe(buffer())
    .pipe(gulpif(conf.DEBUG, sourcemaps.init()))
    .pipe(uglify())
    .pipe(gulpif(conf.DEBUG, sourcemaps.write(dirs.DEST_JS_MAP)))
    .pipe(rev()).pipe(gulp.dest(dirs.DEST_JS));
});

gulp.task('serv', [ 'build' ], () => {
  browserSync.init({
    server: { baseDir: dirs.DEST },
    browser: conf.BROWSER
  });
  gulp.watch(dirs.SRC + '/**/*.html', [ 'html-watch' ]);
  gulp.watch(dirs.SRC + '/**/*.scss', [ 'sass-watch' ]);
});

gulp.task('sass-watch', [ 'html-watch' ]);
gulp.task('html-watch', [ 'index' ], (done) => {
  browserSync.reload();
  done();
});
