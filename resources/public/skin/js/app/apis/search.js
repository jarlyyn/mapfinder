define(function(require) {
    var app = require("app");
    var $ = require("jquery");
    var parsers = require("parsers");
    return function(data,cb) {
      var url = app.Host + app.APIList.search;
      $.post(url, data)
      .done(function(body) {
        cb(body);
      });
    };
  });
  