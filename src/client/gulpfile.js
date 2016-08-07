//'use strict';
var gulp = require('gulp');
var webserver = require('gulp-webserver');

//https://www.npmjs.com/package/gulp-webserver
gulp.task('webserver', function() {
  gulp.src('app')
    .pipe(webserver({
      host: '0.0.0.0',
      port: 9000,
      livereload: true,
      directoryListing: false,
      open: true
    }));
});

//https://github.com/andresvia/go-angular-drone/blob/master/.drone.yml
gulp.task('default', ['webserver'], function() {
  // place code for your default task here
});
