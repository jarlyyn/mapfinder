define(function(require) {
    var app = require("app");
    var $ = require("jquery");
    var parsers = require("parsers");
    return function(vm, cb) {
      var url = app.Host + app.APIList.maplist;
      $.get(url, {page:vm.CurrentPage})
        .done(function(body) {
          vm.Items=parsers.parse200(body)
          cb(vm.Items);
        })
        .fail(function(xhr) {});
    };
  });
  
