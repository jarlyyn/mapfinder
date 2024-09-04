define(function(require) {
  var login = require("js/app/apis/user/login");
  var router = require("js/app/router");
  return {
    name: "login",
    template: require("text!./login.html"),
    data: function() {
      return {
        Username: "",
        Password: "",
        errors: []
      };
    },
    mounted: function() {
      var self = this;
    },
    methods: {
      login: function(event) {
        login(this, function(user) {
          var returnUrlJSON = localStorage.getItem("returnurl");
          var returnUrl = "";
          if (returnUrlJSON !== "") {
            returnUrl = JSON.parse(returnUrlJSON);
          }
          localStorage.setItem("returnurl", "");
          if (returnUrl) {
            router.push(router.resolve(returnUrl).normalizedTo);
          } else {
            router.push("/");
          }
        });
      }
    }
  };
});
