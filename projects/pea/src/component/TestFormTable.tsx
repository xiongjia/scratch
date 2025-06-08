import type {
  ProColumns,
  EditableFormInstance,
  ProFormInstance,
} from '@ant-design/pro-components'
import {
  ProCard,
  ProForm,
  EditableProTable,
  ProFormDependency,
  ProFormField,
} from '@ant-design/pro-components'
import { useRef, useState } from 'react'

// Colume Data type
type TestTableItem = {
  id: number
  name: string
}

const testTableColumns: ProColumns<TestTableItem>[] = [
  {
    title: 'TestName',
    dataIndex: 'name',
    ellipsis: true,
  },
  {
    title: 'TestId',
    dataIndex: 'id',
  },
]

const testData: TestTableItem[] = [
  { id: 1, name: 'item-1' },
  { id: 2, name: 'item-2' },
]

const TestFormTable = () => {
  const formRef = useRef<ProFormInstance<TestTableItem>>()
  const editorFormRef = useRef<EditableFormInstance<TestTableItem>>()
  // const [controlled, setControlled] = useState<boolean>(false)
  return (
    <>
      <ProForm<{ srcData: TestTableItem[] }>
        formRef={formRef}
        initialValues={{ srcData: testData }}
        validateTrigger="onBlur"
      >
        <EditableProTable<TestTableItem>
          name="srcData"
          scroll={{ x: 560 }}
          columns={testTableColumns}
          editableFormRef={editorFormRef}
          form={{
            onValuesChange: (changedValues, allValues) => {
              console.log('changedValues:', changedValues, allValues)
            },
          }}
        />

        <ProForm.Item>
          <ProCard
            title="Table Data sets"
            headerBordered
            collapsible
            defaultCollapsed
          >
            <ProFormDependency name={['srcData']}>
              {(srcData) => {
                return (
                  <>
                    <ProFormField
                      ignoreFormItem
                      fieldProps={{
                        style: {
                          width: '100%',
                        },
                      }}
                      mode="read"
                      valueType="jsonCode"
                      text={JSON.stringify(srcData)}
                    />
                  </>
                )
              }}
            </ProFormDependency>
          </ProCard>
        </ProForm.Item>
      </ProForm>
    </>
  )
}

export default TestFormTable
