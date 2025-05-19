import { StorybookConfig } from 'storybook-react-rsbuild'

const config: StorybookConfig = {
  stories: ['../src/**/*.mdx', '../src/**/*.stories.@(js|jsx|mjs|ts|tsx)'],
  framework: 'storybook-react-rsbuild',
}

export default config
