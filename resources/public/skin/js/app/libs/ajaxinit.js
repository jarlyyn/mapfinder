define(function(require) {
  var app = require("app");
  var $ = require("jquery");
  var router = require("js/app/router");
  var Cookies = require("Cookies");
  return function() {
    $(function() {
      $(document).ajaxStart(function() {
        app.Vue.Loading = true;
      });
      $(document).ajaxStop(function() {
        app.Vue.Loading = false;
      });
    });
    var csrftoken = Cookies.get("herb-csrf-token");
    $.ajaxSetup({
      dataType: "json",
      headers: {
        "X-CSRF-TOKEN": csrftoken
      },
      statusCode: {
        400: function() {
          location.reload(true);
        },
        401: function() {
          localStorage.setItem("returnurl", JSON.stringify(app.RouterEntering));
          router.push("/login");
        },
        403: function() {
          app.Vue.Error = "403";
        },
        404: function() {
          app.Vue.Error = "404";
        },
        500: function() {
          app.Vue.Error = "500";
        }
      }
    });
  };
});
