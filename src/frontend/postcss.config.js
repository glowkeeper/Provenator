const reactToolboxVariables = {};

module.exports = {
  plugins: {
    'postcss-import': {},
    'postcss-cssnext': {
      features: {
        customProperties: {
          variables: reactToolboxVariables,
        },
      }
    },
    'postcss-modules-values': {}
  }
};
