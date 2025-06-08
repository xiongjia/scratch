import { map } from 'lodash-es'
import { Form, Input, Button } from 'antd'
import type { FormListFieldData, FormListOperation, FormProps } from 'antd'
import type { ColumnProps } from 'antd/es/table'
import { ProTable } from '@ant-design/pro-components'
import type { ProColumns } from '@ant-design/pro-components'
import { useEffect, useState } from 'react'

export type HostItem = {
  key: string
  name: string
  nodeName: string
  status: string
  outputLog: string
}

const HostTable = ({ hosts }: { hosts: HostItem[] }) => {
  const [editableKeys, setEditableKeys] = useState<string[]>([])
  useEffect(() => {
    setEditableKeys(map<HostItem, string>(hosts, (itr: HostItem) => itr.key))
  }, [hosts])
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
    {
      title: 'node name',
      dataIndex: 'nodeName',
      render(value, { field }) {
        return (
          <Form.Item
            name={[field.name, 'nodeName']}
            rules={[{ required: true, message: 'required' }]}
          >
            <Input />
          </Form.Item>
        )
      },
    },
  ]

  const hostCols: ProColumns<HostItem>[] = [
    {
      title: '主机名',
      key: 'name',
      width: 250,
      dataIndex: 'name',
      fixed: 'left',
      editable: false,
    },
    {
      title: '状态',
      key: 'status',
      dataIndex: 'status',
      editable: false,
    },
    {
      title: '节点名',
      key: 'nodeName',
      dataIndex: 'nodeName',
      renderFormItem: (_, config, data) => {
        console.log(_, config, data)
        return (
          <Form.Item
            name={'nodeName'}
            initialValue={config.record?.nodeName}
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
                <ProTable<HostItem>
                  size="small"
                  bordered
                  rowKey={(rcd) => rcd.key}
                  dataSource={hosts}
                  pagination={false}
                  columns={hostCols}
                  editable={{
                    editableKeys: editableKeys,
                  }}
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
