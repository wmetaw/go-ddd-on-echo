const gulp = require('gulp');
const aglio = require('gulp-aglio');
const aglioConfig = require('./aglioconfig.json');

// mdファイルからoutputディレクトリにhtmlを出力
gulp.task('output', function() {
  gulp.src('docs/**')
    .pipe(aglio(aglioConfig))
    .pipe(gulp.dest('output'));
});

// mdファイルを監視
gulp.task('watch', function() {
  gulp.watch('docs/**', ['output']);
});

// デフォルトで実行されるタスクを指定
gulp.task('default', ['output', 'watch']);
