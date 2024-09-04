define(function(require) {
  var vueloader = function(name) {
    return function(resolve, reject) {
      require([name], function(component) {
        resolve(component);
      });
    };
  };
  var Vue = require("vue");
  var Router = require("vue-router");
  var app = require("app");
  var CurrentUser = require("js/app/apis/user/current");
  Vue.use(Router);
  var approuter = new Router({
    routes: [
      {
        path: "/",
        name: "console",
        component: vueloader("components/console/index"),
        meta: {
          anonymous: false
        }
      },
      {
        path: "/list",
        name: "list",
        component: vueloader("components/map/list/index"),
        meta: {
          anonymous: false
        }
      },
      {
        path: "/mapview/:id",
        name: "mapview",
        component: vueloader("components/map/view/index"),
        meta: {
          anonymous: false
        }
      },
      // {
      //   path: "/login",
      //   name: "login",
      //   component: vueloader("components/pages/login"),
      //   meta: {
      //     anonymous: true
      //   }
      // },
      // {
      //   path: "/logout",
      //   name: "logout",
      //   component: vueloader("components/pages/logout"),
      //   meta: {
      //     anonymous: true
      //   }
      // },

      {
        path: "*",
        name: "notfound",
        component: vueloader("components/pages/notfound"),
        meta: {
          anonymous: true
        }
      }
    ]
  });
  // approuter.beforeEach(function(to, from, next) {
  //   app.RouterEntering = to.fullPath;
  //   if (app.Vue) {
  //     app.Vue.Error = "";
  //   }
  //   if (
  //     to.meta.anonymous === undefined ||
  //     to.meta.anonymous === null ||
  //     to.meta.anonymous === false
  //   ) {
  //     if (
  //       approuter.app.CurrentUser === null ||
  //       approuter.app.CurrentUser === undefined
  //     ) {
  //       next(false);
  //       CurrentUser(function() {
  //         if (
  //           approuter.app.CurrentUser === null ||
  //           approuter.app.CurrentUser === undefined
  //         ) {
  //           approuter.push("/login");
  //         } else {
  //           approuter.push(to.fullPath);
  //         }
  //       });
  //       return;
  //     }
  //   }
  //   next();
  // });
  return approuter;
});
