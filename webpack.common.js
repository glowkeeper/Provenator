const path = require('path')
const fs  = require('fs')
const webpack = require('webpack')
const htmlWebpackPlugin = require('html-webpack-plugin');
const htmlWebpackInlineSourcePlugin = require('html-webpack-inline-source-plugin')
const { CleanWebpackPlugin } = require('clean-webpack-plugin')
const lessToJs = require('less-vars-to-js');
const themeVariables = lessToJs(fs.readFileSync(path.join(__dirname, './app/stylesheets/themeVariables.less'), 'utf8'))
themeVariables["@icon-url"] = 'https://fonts.googleapis.com/css?family=Source+Sans+Pro:regular,bold,italic&subset=latin,latin-ext'

var config = {
  node: {
    console: true,
    fs: 'empty',
    net: 'empty',
    tls: 'empty'
  },
  entry: {
    app: [
      '@babel/polyfill',
      './app/index.jsx'
    ]
  },
  output: {
    path: path.resolve(__dirname, 'build')
  },
  resolve: {
    extensions: ['.js', '.jsx']
  },
  plugins: [
    new webpack.ProgressPlugin(),
    new CleanWebpackPlugin(),
    new htmlWebpackPlugin({
      template: './app/index.html',
      inject: 'body',
      inlineSource: '.(js|css)$'
    }),
    new htmlWebpackInlineSourcePlugin(),
  ],
  module: {
    rules: [
      {
        test: /\.less$/,
        use: [
          {
            loader: "style-loader" // creates style nodes from JS strings
          },
          {
            loader: "css-loader" // translates CSS into CommonJS
          },
          {
            loader: "less-loader", // compiles Less to CSS
            options: {
              javascriptEnabled: true,
              modifyVars: themeVariables
            }
          }
        ]
      },
      {
        test: /\.(jp(e*)g|png|gif|svg|pdf|ico)$/i,
        use: [
          {
            loader: 'url-loader',
            options: {
              limit: 8192,
              name: '[sha512:hash:base64:7].[ext]'
            }
          }
        ]
      },
      {
        test: /\.jsx?/,
        include: path.resolve(__dirname, 'app'),
        exclude: /node_modules/,
        loader: 'babel-loader'
      }
    ]
  }
}

module.exports = config
