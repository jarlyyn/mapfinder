define(function(require) {
  var logout = require("js/app/apis/user/logout");
  var router = require("js/app/router");
  return {
    name: "logout",
    template: require("text!./logout.html"),
    mounted: function() {
      logout(this, function() {
        router.push("/login");
      });
    },
    data() {
      return {};
    }
  };
});
