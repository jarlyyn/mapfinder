define(function(require) {
  return {
    name: "error",
    props: ["error"],
    template: require("text!./error.html"),
    data: function() {
      return {};
    }
  };
});
