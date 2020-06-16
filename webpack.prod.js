const webpack = require('webpack');
const uglifyJSPlugin = require('uglifyjs-webpack-plugin');
const merge = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = merge(common, {
  plugins: [
    new webpack.DefinePlugin({
     'process.env': {
       'NODE_ENV': JSON.stringify('production')
      }
    }),
    new uglifyJSPlugin({
      "test": /\.js$/i,
      "extractComments": false,
      "sourceMap": false,
      "cache": true,
      "parallel": true,
      "uglifyOptions": {
        "output": {
          "ascii_only": true,
          "comments": false
        },
        "ecma": 5,
        "warnings": false,
        "ie8": false,
        "mangle": true,
        "compress": {
          sequences: true,
      		dead_code: true,
      		conditionals: true,
      		booleans: true,
      		unused: true,
      		if_return: true,
      		join_vars: true,
      		drop_console: true
        }
      }
    })
  ]
});
