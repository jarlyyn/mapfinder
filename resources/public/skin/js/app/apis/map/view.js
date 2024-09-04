define(function(require) {
    var app = require("app");
    var $ = require("jquery");
    var parsers = require("parsers");
    return function(vm, cb) {
      var url = app.Host + app.APIList.mapview+vm.id;
      $.get(url, {})
        .done(function(body) {
          var data = parsers.parse200(body);
          vm.Item=data
          cb(data);
        })
        .fail(function(xhr) {});
    };
  });
  
