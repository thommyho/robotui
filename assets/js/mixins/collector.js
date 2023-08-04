export default {
  methods: {
    // collect all target component properties from current instance
    collectProps: function (component) {
      let data = {};
      for (var p in component.props) {
        if (p in this) {
          data[p] = this[p];
        }
      }
      return data;
    },
  },
};
