import type { Meta, StoryObj } from '@storybook/react'
import { TestFormTable } from '&/component'

const meta: Meta<typeof TestFormTable> = {
  title: 'Example/Test Table',
  component: TestFormTable,
  argTypes: {},
}
export default meta

type Story = StoryObj<typeof TestFormTable>

export const InputPrimary: Story = {
  name: 'Test 2',
  args: {},
}
