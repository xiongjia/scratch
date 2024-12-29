import { Form, Table, Input, Button } from 'antd'
import type { FormListFieldData, FormListOperation, FormProps } from 'antd'
import type { ColumnProps } from 'antd/es/table'

export type HostItem = {
  key: string
  name: string
  nodeName: string
  status: string
  outputLog: string
}

const HostTable = ({ hosts }: { hosts: HostItem[] }) => {
  const columns: ColumnProps<{
    field: FormListFieldData
    operation: FormListOperation
  }>[] = [
    {
      title: 'host name',
      dataIndex: 'name',
      render(value, { field }) {
        return (
          <Form.Item
            name={[field.name, 'name']}
            rules={[{ required: true, message: 'required' }]}
          >
            <Input />
          </Form.Item>
        )
      },
    },
  ]

  const onFinish: FormProps<FormData>['onFinish'] = (values) => {
    console.log(values)
  }

  return (
    <>
      <Form
        initialValues={{
          hosts: hosts,
        }}
        onFinish={onFinish}
      >
        <Form.Item label="hosts">
          <Form.List name="hosts">
            {(fields, operation) => {
              const dataSources = fields.map((field) => ({
                field,
                operation,
              }))
              console.log('dataSources', dataSources)
              return (
                <Table
                  size="small"
                  bordered
                  rowKey={(row) => row.field.key}
                  dataSource={dataSources}
                  pagination={false}
                  columns={columns}
                />
              )
            }}
          </Form.List>
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>
    </>
  )
}

export default HostTable
