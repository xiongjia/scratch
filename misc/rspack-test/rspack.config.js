const rspack = require('@rspack/core');

const config = {
  entry: './src/index.js',
  plugins: [
    new rspack.HtmlRspackPlugin({
      template: './public/index.html',
    }),
  ],
};
module.exports = config;
