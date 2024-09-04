define(function(require) {
  var app = require("app");
  var $ = require("jquery");
  var parsers = require("parsers");
  return function(vm, cb) {
    var url = app.Host + app.APIList.login;
    $.post(
      url,
      JSON.stringify({ Username: vm.Username, Password: vm.Password })
    )
      .done(function(body) {
        cb(body);
      })
      .fail(function(xhr) {
        if (xhr.status === 422) {
          vm.errors = parsers.parse422(xhr.responseJSON);
        }
      });
  };
});
