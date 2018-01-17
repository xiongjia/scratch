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
  DEBUG: argv.debug,
  NAME: 'bootstrap-scratch',
  DESC: 'bootstrap scratch',
  VER: '0.1',
  DEV_NAME: 'xiong-jia.le',
  DEV_URL: 'https://github.com/xiongjia'
};

const dirs = {
  SRC: 'src',
  SRC_BOOTSTRAP_SASS: 'node_modules/bootstrap-sass',
  SRC_JQUERY: 'node_modules/jquery',
  DEST: 'dist',
  DEST_ASSETS: 'dist/assets',
  DEST_CSS: 'dist/css',
  DEST_CSS_MAP: '.',
  DEST_FONTS: 'dist/fonts',
  DEST_JS: 'dist/js',
  DEST_JS_MAP: '.'
};

gutil.log('bootstrap scratch');
gutil.log('conf = %j', conf);
gutil.log('dirs = %j', dirs);

gulp.task('build', seq('lint:js', [ 'fonts', 'assets:fav', 'index' ]));
gulp.task('default', seq('clean', 'lint:js', 'build'));

gulp.task('clean:all', () => del([ dirs.DEST ]));
gulp.task('clean:js', () => del([ dirs.DEST + '/js/**/*.{js,map}' ]));
gulp.task('clean:css', () => del([ dirs.DEST + '/css/**/*.{css,map}' ]));
gulp.task('clean:assets:fav', () => del([ dirs.DEST + '/assets/favicon*.{ico,png}' ]));
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
    .pipe(gulpif(conf.DEBUG, sourcemaps.write(dirs.DEST_CSS_MAP, {
      addComment: false
    })))
    .pipe(gulpif(!conf.DEBUG, rev()))
    .pipe(gulp.dest(dirs.DEST_CSS));
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
    dirs.DEST + '/**/jquery*.js',
    dirs.DEST + '/**/bootstrap*.js',
    dirs.DEST + '/**/bundle*.js'
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

gulp.task('js', seq('clean:js', [ 'js:libs', 'js:bundle' ]));

gulp.task('js:libs', () => {
  const libs = [
    dirs.SRC_JQUERY + '/dist/jquery.min.js',
    dirs.SRC_BOOTSTRAP_SASS + '/assets/javascripts/bootstrap.min.js'
  ];
  return gulp.src(libs)
    .pipe(gulp.dest(dirs.DEST_JS));
});

gulp.task('js:bundle', () => {
  const sourcemaps = require('gulp-sourcemaps');
  const source = require('vinyl-source-stream');
  const buffer = require('vinyl-buffer');
  const browserify = require('browserify');
  const uglify = require('gulp-uglify');
  const envify = require('envify/custom');

  const ent = {
    entries: [ 'src/main.js'],
    debug: conf.DEBUG
  };

  return browserify(ent)
    .transform('babelify', { presets: [ 'env' ] })
    .transform(envify({
      ENV_DEBUG: conf.DEBUG
    }))
    .bundle()
    .pipe(source('bundle.js'))
    .pipe(buffer())
    .pipe(gulpif(!conf.DEBUG, sourcemaps.init()))
    .pipe(gulpif(!conf.DEBUG, uglify()))
    .pipe(gulpif(!conf.DEBUG, sourcemaps.write(dirs.DEST_JS_MAP, {
      addComment: false
    })))
    .pipe(gulpif(!conf.DEBUG, rev()))
    .pipe(gulp.dest(dirs.DEST_JS));
});

gulp.task('assets:fav', [ 'clean:assets:fav' ], () => {
  const favicons = require('gulp-favicons');
  return gulp.src(dirs.SRC + '/assets/fav.png')
    .pipe(favicons({
      appName: conf.NAME,
      appDescription: conf.DESC,
      developerName: conf.DEV_NAME,
      developerURL: conf.DEV_URL,
      background: '#fff',
      theme_color: '#fff',
      path: '/',
      display: 'standalone',
      orientation: 'portrait',
      version: conf.VER,
      logging: false,
      online: false,
      pipeHTML: false,
      replace: false,
      icons: {
        android: false,
        appleIcon: false,
        appleStartup: false,
        coast: false,
        favicons: true,
        firefox: false,
        windows: false,
        yandex: false
      }
    }))
    .on('error', gutil.log)
    .pipe(gulp.dest(dirs.DEST_ASSETS));
});

gulp.task('serv', [ 'build' ], () => {
  browserSync.init({
    server: { baseDir: dirs.DEST },
    browser: conf.BROWSER
  });
  gulp.watch(dirs.SRC + '/**/*.html', [ 'html-watch' ]);
  gulp.watch(dirs.SRC + '/**/*.scss', [ 'sass-watch' ]);
  gulp.watch(dirs.SRC + '/assets/fav*.png', [ 'assets-watch:fav' ]);
});

gulp.task('assets-watch:fav', [ 'assets:fav' ], (done) => {
  browserSync.reload();
  done();
});

gulp.task('sass-watch', [ 'html-watch' ]);
gulp.task('html-watch', [ 'index' ], (done) => {
  browserSync.reload();
  done();
});

