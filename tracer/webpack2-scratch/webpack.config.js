'use strict';

const path = require('path');

const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

const prod = (process.env.NODE_ENV === 'production');

const plugins = (() => {
  let webpackPlugins = [];
  if (prod) {
    webpackPlugins.push(new webpack.optimize.UglifyJsPlugin());
    webpackPlugins.push(new ExtractTextPlugin('style-[contenthash:10].css'));
  }
  webpackPlugins.push(new HtmlWebpackPlugin({
    inject: true,
    minify: (() => {
      if (prod) {
        return {
          collapseWhitespace: true,
          removeComments: true,
          removeRedundantAttributes: true
        };
      }
      return {};
    })(),
    template: path.join(__dirname, 'public/index-template.html')
  }));
  webpackPlugins.push(new webpack.DefinePlugin({
    PRODUCTION: prod
  }));
  return webpackPlugins;
})();

exports = module.exports = {
  devtool: 'source-map',
  entry: './src/index.js',
  plugins: plugins,
  output: {
    path: path.join(__dirname, 'dist'),
    filename: prod ? 'bundle.[hash:12].min.js' : 'bundle.js'
  },
  module: {
    loaders: [{
      enforce: 'pre',
      test: /\.(js|jsx)$/,
      exclude: /node_modules/,
      loader: 'eslint-loader'
    }, {
      test: /\.(js|jsx)$/,
      exclude: /node_modules/,
      loaders: ['babel-loader']
    }, {
      test: /\.(png|jpg|gif)$/,
      loaders: ['url-loader?limit=10000&name=images/[hash:12].[ext]'],
      exclude: /node_modules/
    }, {
      test: /\.css$/,
      loaders: (() => {
        if (prod) {
          return ExtractTextPlugin.extract({
            loader: 'css-loader?minimize&localIdentName=[hash:base64:10]'
          });
        }
        return [
          'style-loader',
          'css-loader?localIdentName=[path][name]---[local]'
        ];
      })(),
      exclude: /node_modules/
    }]
  }
};
