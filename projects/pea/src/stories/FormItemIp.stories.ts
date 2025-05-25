import type { Meta, StoryObj } from '@storybook/react'
import { InputIp } from '&/component/input'

const meta: Meta<typeof InputIp> = {
  title: 'Example/Input IP',
  component: InputIp,
  argTypes: {
    placeholder: {
      name: 'placeholder lab',
      control: 'text',
      description: 'placeholder desc',
    },
  },
}
export default meta

type Story = StoryObj<typeof InputIp>

export const InputPrimary: Story = {
  name: 'input 1',
  args: {
    placeholder: 'placeholder21',
  },
  parameters: {
    status: { default: 'error' },
  },
}

export const InputWarr: Story = {
  name: 'input 2',
  args: {
    ...InputPrimary.args,
    status: 'error',
  },
}
