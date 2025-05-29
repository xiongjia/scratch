import type { Meta, StoryObj } from '@storybook/react'
import { TestSteps } from '&/component'

const meta: Meta<typeof TestSteps> = {
  title: 'Example/Test Steps',
  component: TestSteps,
  argTypes: {},
}
export default meta

type Story = StoryObj<typeof TestSteps>

export const InputPrimary: Story = {
  name: 'Test 1',
  args: {},
}
