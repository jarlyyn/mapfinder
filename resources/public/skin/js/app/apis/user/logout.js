define(function(require) {
  var app = require("app");
  var $ = require("jquery");

  return function(vm, cb) {
    var url = app.Host + app.APIList.logout;
    $.post(url).done(function() {
      app.Vue.CurrentUser = null;
      cb();
    });
  };
});
