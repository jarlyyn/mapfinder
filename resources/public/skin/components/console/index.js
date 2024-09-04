define(function(require) {
  var search = require("js/app/apis/search");
  return {
    name: "console",
    template: require("text!./index.html"),
    props: ["user"],
    methods: {
      onSearch:function(){
        var self=this;
        search(this.Map,function(result){
          self.Searched=true
          self.Result=result
        })
      }
    },
    data: function() {
      return {
        Map:"",
        Searched:false,
        Result:null,
      };
    }
  };
});
