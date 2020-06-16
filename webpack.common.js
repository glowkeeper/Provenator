const path = require('path')
const fs = require('fs')
const webpack = require('webpack')
const htmlWebpackPlugin = require('html-webpack-plugin')
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const HTMLInlineCSSWebpackPlugin = require("html-inline-css-webpack-plugin").default;
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
    new MiniCssExtractPlugin({
      filename: "[name].css",
      chunkFilename: "[id].css"
    }),
    new htmlWebpackPlugin({
      template: './app/index.html',
      inject: 'body',
      inlineSource: '.(js|css)$'
    }),
    new HTMLInlineCSSWebpackPlugin(),
  ],
  module: {
    rules: [
      {
        test: /\.css$/,
        use: [
          MiniCssExtractPlugin.loader,
          "css-loader"
        ]
    },
      {
        test: /\.less$/,
        use: [
          {
            loader: "style-loader" // creates style nodes from JS strings
          },
          {
              loader: 'less-loader',
              options: {
                lessOptions: {
                    javascriptEnabled: true,
                    modifyVars: themeVariables
                }
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
          test: /\.js$/,
          use: ["source-map-loader"],
          exclude: /node_modules/,
          enforce: "pre"
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
