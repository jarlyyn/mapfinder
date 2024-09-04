define(function(require) {
  var itemview=require("js/app/apis/map/view")
  var itemupdate=require("js/app/apis/map/update")
  var successurl="/list"
  return {
    name: "listmap",
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
      load: function() {
        var self = this;
        this.id=this.$route.params.id;
        if (this.id){
        itemview(this,function(){
        })
      }
      },
      onSubmit: function() {
        var self=this;
        this.errors = [];
        self.$refs["form"].validate(function(valid){
          if (valid){
            itemupdate(self,function(){
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
        id:"",
        Item:{},
        rules:{
        },
        errors: []
      };
    }
  };
});
