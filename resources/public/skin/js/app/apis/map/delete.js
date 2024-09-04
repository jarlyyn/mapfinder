define(function(require) {
    var app = require("app");
    var $ = require("jquery");
    var parsers = require("parsers");
    return function(vm, id,cb) {
      var url = app.Host + app.APIList.mapdelete + id;
      $.post(url)
        .done(function(body) {
          var data = parsers.parse200(body);
          cb(data);
        })
        .fail(function(xhr) {
          if (xhr.status === 422) {
            vm.errors = parsers.parse422(xhr.responseJSON);
          }
        });
    };
  });
  