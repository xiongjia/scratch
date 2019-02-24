'use strict';

module.exports = {
  root: true,
  parserOptions: {
    parser: 'babel-eslint'
  },
  env: {
    browser: true,
  },
  extends: [ 'plugin:vue/essential',  'standard' ],
  plugins: [ 'vue' ],
  rules: {
    'generator-star-spacing': 'off',
    'semi': [2, 'always'],
    'vue/no-parsing-error': [2, { 'x-invalid-end-tag': false }],
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off'
  }
};
