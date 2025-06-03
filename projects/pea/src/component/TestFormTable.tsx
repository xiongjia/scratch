import type { ProColumns } from '@ant-design/pro-components'
import {
  EditableProTable,
  ProForm,
  ProFormText,
} from '@ant-design/pro-components'

type TestTableItem = {
  id: number
  name: string
}

const testTableColumns: ProColumns<TestTableItem>[] = [
  {
    title: 'TestName',
    dataIndex: 'name',
    copyable: true,
    ellipsis: true,
  },
]

const testData: TestTableItem[] = [
  { id: 1, name: 'item1' },
  { id: 2, name: 'item2' },
]

type TestFormData = {
  username: string
}

const TestFormTable = () => {
  return (
    <>
      <ProForm<TestFormData>
        onFinish={async (val: TestFormData): Promise<boolean> => {
          console.log('finsh proform', val)
          return true
        }}
        initialValues={{
          username: '123',
        }}
      >
        <ProFormText name={'username'} />
      </ProForm>

      <br />

      <EditableProTable<TestTableItem>
        columns={testTableColumns}
        dataSource={testData}
        form={{
          initialValues: { id: 1, name: 'a' },
          onValuesChange: (changedValues, allValues) => {
            console.log('changedValues:', changedValues, allValues)
          },
        }}
      />
    </>
  )
}

export default TestFormTable
