'use strict'

const htmlmin = require('html-minifier')

module.exports = (cfg) => {
  // cfg.addTransform('htmlmin', (content) => {
  //   if (!this.page.outputPath) {
  //     return content
  //   }
  //   if (!this.page.outputPath.endsWith('.html')) {
  //     return content
  //   }
  //   return htmlmin.minify(content, {
  //     collapseWhitespace: true,
  //     useShortDoctype: true,
  //   })
  // })

  return {
    dir: {
      input: 'src',
      output: 'dist',
    },
    // templateFormats: ['md', 'html'],
    // markdownTemplateEngine: 'njk',
    // dataTemplateEngine: 'njk',
    // htmlTemplateEngine: 'njk',
  }
}
