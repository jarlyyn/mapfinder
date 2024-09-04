define(function(require) {
  var app = {};
  app.Host = "";
  app.APIList = {
    // current: "/api/current",
    // logout: "/api/logout",
    // login: "/login"
    search:"/api/search",
    maplist:"/api/map/list",
    mapview:"/api/map/view/"
  };
  return app;
});
