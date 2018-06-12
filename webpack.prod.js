const webpack = require('webpack');
const uglifyJSPlugin = require('uglifyjs-webpack-plugin');
const merge = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = merge(common, {
  devtool: 'cheap-module-source-map',
  plugins: [
    new uglifyJSPlugin({
      "test": /\.js$/i,
      "extractComments": false,
      "sourceMap": true,
      "cache": false,
      "parallel": false,
      "uglifyOptions": {
        "output": {
          "ascii_only": true,
          "comments": false
        },
        "ecma": 5,
        "warnings": false,
        "ie8": false,
        "mangle": true,
        "compress": {}
      }
    }),
    new webpack.DefinePlugin({
     'process.env': {
       'NODE_ENV': JSON.stringify('production')
     }
    })
  ]
});
