define(function(require) {
  return {
    name: "Menu",
    props: ["user"],
    template: require("text!./navmenu.html"),
    data: function() {
      return {};
    }
  };
});
