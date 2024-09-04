define(function(require) {
  Vue = require("vue");
  function format(n) {
    if (n < 10) {
      return "0" + n;
    }
    return n;
  }
  return function() {
    Vue.filter("datetime", function(ts) {
      if (ts == 0) {
        return "-";
      }
      var date = new Date(ts * 1000);
      return (
        date.getFullYear() +
        "-" +
        format(date.getMonth() + 1) +
        "-" +
        format(date.getDate()) +
        " " +
        format(date.getHours()) +
        ":" +
        format(date.getMinutes()) +
        ":" +
        format(date.getSeconds())
      );
    });
  };
});
