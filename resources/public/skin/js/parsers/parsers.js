define(function(require) {
  return {
    parse200: function(data) {
      return data;
    },
    parseItems: function(data) {
      return data.Items;
    },
    parseItemsCount: function(data) {
      return data.Count;
    },
    parse422: function(data) {
      return data;
    }
  };
});
