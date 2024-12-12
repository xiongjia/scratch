const path = require('path');
const rspack = require('@rspack/core');

const config = {
  entry: './src/index.ts',
  resolve: {
    tsConfig: {
      configFile: path.resolve(__dirname, 'tsconfig.json'),
    },
    extensions: ['...', '.ts'],
  },
  plugins: [
    new rspack.HtmlRspackPlugin({
      template: './public/index.html',
    }),
  ],
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: [{
          loader: 'builtin:swc-loader',
          options: { jsc: { parser: { syntax: 'typescript' } } }
        }],
      },
    ],
  },
};

module.exports = config;
