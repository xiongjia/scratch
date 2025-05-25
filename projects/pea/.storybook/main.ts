import type { StorybookConfig } from 'storybook-react-rsbuild'

const config: StorybookConfig = {
  stories: ['../src/**/*.mdx', '../src/**/*.stories.@(js|jsx|mjs|ts|tsx)'],
  framework: 'storybook-react-rsbuild',
  addons: [
    '@storybook/addon-essentials',
    '@storybook/addon-actions',
    '@storybook/addon-links',
    '@storybook/addon-docs',
    '@storybook/addon-toolbars',
    '@storybook/addon-measure',
  ],
}

export default config
