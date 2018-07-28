'use strict';

const path = require('path');
const webpack = require('webpack');

const CleanWebpackPlugin = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const BundleAnalyzerPlugin = require('webpack-bundle-analyzer')
  .BundleAnalyzerPlugin;

const buildTM = new Date();
const conf = {
  debug: process.env.NODE_ENV !== 'production',
  devSrvPort: 9502,
  apiSrvPort: 9501,
  analyzerReporter: !!process.env.BUILD_ANALYZER,
  buildTS: buildTM.valueOf(),
  buildTM: buildTM.toISOString(),
  buildOS: require('os').platform()
};

const dirs = {
  DIST: path.join(__dirname, 'dist'),
  SRC_ENTRY_JS: path.join(__dirname, './src/index.jsx'),
  SRC_ENTRY_PAGE: path.join(__dirname, './src/index-template.html')
};
 
exports = module.exports = {
  devtool: conf.debug ? 'inline-source-map' : undefined,
  entry: [ dirs.SRC_ENTRY_JS ],
  output: {
    path: dirs.DIST,
    filename: conf.debug ? 'js/bundle.js' : 'js/bundle.[hash:6].min.js'
  },
  module: {
    rules: [{
      test: /\.jsx?$/,
      exclude: /node_modules/,
      loader: 'babel-loader',
      query: { cacheDirectory: true }
    }]
  },
  plugins: (() => {
    const plugins = [];
    if (conf.debug) {
      plugins.push(new webpack.HotModuleReplacementPlugin());
    }
    plugins.push(new CleanWebpackPlugin([ dirs.DIST ], {
      verbose: true,
      dry: false
    }));
    plugins.push(new HtmlWebpackPlugin({
      inject: true,
      template: dirs.SRC_ENTRY_PAGE,
      minify: conf.debug ? {} : {
        collapseWhitespace: true,
        removeComments: true,
        removeRedundantAttributes: true
      }
    }));
    plugins.push(new ExtractTextPlugin(
      conf.debug ? 'css/bundle.css' : 'css/bundle.[hash:6].css'));
    if (!conf.debug) {
      plugins.push(new webpack.IgnorePlugin(/^\.\/locale$/, /moment$/));
      plugins.push(new UglifyJsPlugin({
        parallel: 8,
        sourceMap: false,
        uglifyOptions: {
          output: {
            comments: false,
            beautify: false
          }
        }
      }));
    }
    if (conf.analyzerReporter) {
      plugins.push(new BundleAnalyzerPlugin({
        openAnalyzer: false,
        analyzerMode: 'server'
      }));
    }
    return plugins;
  })(),
  devServer: {
    contentBase: dirs.DIST,
    compress: true,
    port: conf.devSrvPort,
    proxy: {
      '/api/*': {
        target: `http://localhost:${conf.apiSrvPort}`
      }
    }
  }
}

