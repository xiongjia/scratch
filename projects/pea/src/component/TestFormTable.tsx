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

type TestSrcData = {
  srcData: TestTableItem[]
}

const testData: TestTableItem[] = [
  { id: 1, name: 'item-1' },
  { id: 2, name: 'item-2' },
  { id: 3, name: 'item-3' },
]

const TestFormTable = () => {
  const [editableKeys, setEditableRowKeys] = useState<React.Key[]>(() =>
    testData.map((i) => i.id),
  )

  const formRef = useRef<ProFormInstance<TestSrcData>>()
  const editorFormRef = useRef<EditableFormInstance<TestTableItem>>()
  // const [controlled, setControlled] = useState<boolean>(false)

  const testTableColumns: ProColumns<TestTableItem>[] = [
    {
      title: 'TestName',
      dataIndex: 'name',
      formItemProps: () => {
        return {
          rules: [{ required: true, message: 'must hava a value' }],
        }
      },
    },
    {
      title: 'TestId',
      dataIndex: 'id',
    },
    {
      title: 'Opts',
      valueType: 'option',
      render: (text, record, _, action) => [
        <a
          key="editable"
          onClick={() => {
            console.log('edit')
            action?.startEditable?.(record.id, record)
          }}
        >
          editOpt
        </a>,
        <a
          key="delete"
          onClick={() => {
            console.log('delete ', record.id)
            const dataSrc = formRef.current?.getFieldValue(
              'srcData',
            ) as TestTableItem[]
            formRef.current?.setFieldsValue({
              srcData: dataSrc.filter((item) => item.id !== record.id),
            })
          }}
        >
          delOpt
        </a>,
      ],
    },
  ]

  return (
    <>
      <ProForm<TestSrcData>
        formRef={formRef}
        initialValues={{ srcData: testData }}
        validateTrigger="onBlur"
      >
        <EditableProTable<TestTableItem>
          headerTitle="hdr title"
          name="srcData"
          rowKey={'id'}
          scroll={{ x: 560 }}
          columns={testTableColumns}
          editableFormRef={editorFormRef}
          form={{
            onValuesChange: (changedValues, allValues) => {
              console.log('changedValues:', changedValues, allValues)
            },
          }}
          editable={{
            type: 'multiple',
            editableKeys,
            onChange: (keys, rows) => {
              console.log('edit-keys', keys)
              console.log('edit-rows', rows)
              setEditableRowKeys(keys)
            },
            actionRender: (row, config, defaultDom) => {
              return [defaultDom.save, defaultDom.delete, defaultDom.cancel]
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
