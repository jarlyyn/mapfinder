define(function(require) {
  var Vue = require("vue");
  var template = require("text!./app.html");
  var Navmenu = require("components/widgets/navmenu");
  var Errorpage = require("components/widgets/error");

  return {
    name: "app",
    template: template,
    props: ["user", "error","params"],
    components: {
      Navmenu: Navmenu,
      Errorpage: Errorpage
    }
  };
});
