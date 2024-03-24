'use strict'

const htmlmin = require('html-minifier')

module.exports = (cfg) => {
  cfg.addTransform('async-htmlmin', async (content, outputPath) => {
    if (outputPath.toLowerCase().endsWith('.html')) {
      return htmlmin.minify(content, {
        useShortDoctype: true,
        removeComments: true,
        collapseWhitespace: true,
      })
    } else {
      return content
    }
  })

  return {
    dir: {
      input: 'src',
      output: 'dist',
    },
    templateFormats: ['md', 'html'],
    markdownTemplateEngine: 'njk',
    dataTemplateEngine: 'njk',
    htmlTemplateEngine: 'njk',
  }
}
