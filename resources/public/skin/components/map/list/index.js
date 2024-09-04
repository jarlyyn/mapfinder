define(function(require) {
  var itemlist = require("js/app/apis/map/list");
  //var itemedelete = require("js/app/apis/itemedelete");
  var lodash = require("lodash");
  return {
    name: "mapist",
    watch: {
      $route: function(to, from) {
        this.load();
      }
    },
    mounted: function() {
      this.load();
    },
    methods: {
      load: function() {
        var self = this;
        itemlist(self, function() {});
      },
      handleDelete: function(item) {
        var self = this;
        self.errors = [];
        this.$confirm("此操作将永久删除该对象, 是否继续?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        })
          .then(function() {
            itemedelete(self, item.ID, function() {
              self.load();
            });
          })
          .catch(function() {});
      },
      onEdit: function(id) {
        this.$router.push({ name: "itemeditname", params: { id: id } });
      },
      onView: function(id) {
        this.$router.push({ name: "mapview", params: { id: id } });
      },
    },
    template: require("text!./index.html"),
    data: function() {
      return {
        Items: [],
        errors: [],
        Count: null,
      };
    }
  };
});
