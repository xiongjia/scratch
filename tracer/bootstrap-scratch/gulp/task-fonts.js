'use strict';

const gulp = require('gulp');

exports = module.exports = (conf, dirs) => {
  gulp.task('fonts', () => {
    const src = [
      dirs.SRC_BOOTSTRAP_SASS + '/assets/fonts/**/*'
    ];
    return gulp.src(src).pipe(gulp.dest(dirs.DEST_FONTS));
  });
};
