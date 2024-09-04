define(function(require) {
  var Vue = require("vue");
  var template = require("text!./notfound.html");
  var Errorpage = require("components/widgets/error");
  return {
    name: "notfound",
    template: template,
    components: {
      Errorpage: Errorpage
    },
    data: function() {
      return {};
    }
  };
});
