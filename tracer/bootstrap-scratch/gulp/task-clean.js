'use strict';

const gulp = require('gulp');
const del = require('del');

exports = module.exports = (conf, dirs) => {
  gulp.task('clean:all', () => del([ dirs.DEST ]));
  gulp.task('clean:js', () => del([ dirs.DEST + '/js/**/*.{js,map}' ]));
  gulp.task('clean:css', () => del([ dirs.DEST + '/css/**/*.{css,map}' ]));
  gulp.task('clean:assets:img', () => del([ dirs.DEST + '/assets/img/**/*' ]));
  gulp.task('clean:assets:fav', () => {
    return del([ dirs.DEST + '/assets/favicon*.{ico,png}' ]);
  });
  gulp.task('clean', [ 'clean:all' ]);
};
