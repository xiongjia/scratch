'use strict'

const htmlmin = require('html-minifier')
const hjson = require('hjson')
const eleventySass = require('eleventy-sass')

module.exports = (cfg) => {
  if (process.env.REACT_APP === 'production') {
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
  }

  cfg.addDataExtension('hjson', (content) => hjson.parse(content))
  cfg.addLayoutAlias('base', 'base.njk')
  cfg.addPlugin(eleventySass)

  return {
    dir: {
      input: 'src',
      output: 'dist',
      includes: 'includes',
      data: 'data',
    },
    templateFormats: ['md', 'njk'],
    markdownTemplateEngine: 'njk',
    dataTemplateEngine: 'njk',
    htmlTemplateEngine: 'njk',
  }
}
