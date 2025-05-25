import type { Meta, StoryObj } from '@storybook/react'
import { InputIp } from '&/component/input'

const inputStatus = {
  Error: 'error',
  Warn: 'warning',
  Normal: '',
}

const meta: Meta<typeof InputIp> = {
  title: 'Example/Input IP',
  component: InputIp,
  argTypes: {
    placeholder: {
      name: 'placeholder lab',
      control: 'text',
      description: 'placeholder desc',
    },
    status: {
      options: Object.keys(inputStatus),
      mapping: inputStatus,
      control: {
        type: 'select',
        labels: {
          Error: 'error',
          Warn: 'warning',
          Normal: '',
        },
      },
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
