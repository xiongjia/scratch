'use strict';

const gulp = require('gulp');
const browserSync = require('browser-sync').create();

exports = module.exports = (conf, dirs) => {
  gulp.task('serv', [ 'build' ], () => {
    browserSync.init({
      server: { baseDir: dirs.DEST },
      browser: conf.BROWSER
    });
    gulp.watch(dirs.SRC + '/**/*.html', [ 'html-watch' ]);
    gulp.watch(dirs.SRC + '/**/*.scss', [ 'html-watch' ]);
    gulp.watch(dirs.SRC + '/assets/fav*.png', [ 'assets-watch:fav' ]);
    gulp.watch(dirs.SRC + '/assets/img/**/*', [ 'html-watch' ]);
    gulp.watch(dirs.SRC + '/**/*.js', [ 'html-watch' ]);
  });

  gulp.task('assets-watch:fav', [ 'assets:fav' ], (done) => {
    browserSync.reload();
    done();
  });

  gulp.task('html-watch', [ 'index' ], (done) => {
    browserSync.reload();
    done();
  });
};
