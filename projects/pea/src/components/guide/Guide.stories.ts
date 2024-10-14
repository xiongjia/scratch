import type { Meta, StoryObj } from '@storybook/react'

import { Guide } from './Guide'

const meta = {
  title: 'Example/Guide',
  component: Guide,
  parameters: {},
  tags: ['autodocs'],
  argTypes: {},
  args: {},
} satisfies Meta<typeof Guide>

export default meta
type Story = StoryObj<typeof meta>

export const Secondary: Story = {
  args: {
    title: 'secondary',
  },
}

export const Primary: Story = {
  args: {
    title: 'primary',
  },
}
