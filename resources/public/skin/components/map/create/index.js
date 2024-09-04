define(function(require) {
  var itemcreate=require("js/app/apis/itemcreate")
  var successurl="successurl"
  return {
    name: "componentname",
    template: require("text!./index.html"),
    watch: {
      $route: function(to, from) {
        this.load();
      }
    },
    mounted: function() {
      this.load();
    },
    methods: {
      load: function() {},
      onSubmit: function() {
        var self=this;
        this.errors = [];
        self.$refs["form"].validate(function(valid){
        if (valid){
          itemcreate(self,function(){
          if (self.errors.length==0){
            self.$router.push(successurl);
          }
          })
        }else{
          return false
        }
        })
      }
    },
    data: function() {
      return {
        Item: {},
        rules:{
        },
        errors: []
      };
    }
  };
});
