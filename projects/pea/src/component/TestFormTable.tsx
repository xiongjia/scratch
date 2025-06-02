import type { ProColumns } from '@ant-design/pro-components'
import { ProTable } from '@ant-design/pro-components'

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

const TestFormTable = () => {
  return (
    <>
      <ProTable<TestTableItem>
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
