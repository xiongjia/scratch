import { map } from 'lodash-es'
import { useEffect, useState } from 'react'
import type { ProColumns } from '@ant-design/pro-components'
import { ProTable } from '@ant-design/pro-components'

export type HostItem = {
  key: string
  name: string
  nodeName: string
  status: string
  outputLog: string
}

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
  },
]

const HostTable = ({ hosts }: { hosts: HostItem[] }) => {
  const [editableKeys, setEditableKeys] = useState<string[]>([])

  useEffect(() => {
    setEditableKeys(map<HostItem, string>(hosts, (itr: HostItem) => itr.key))
  }, [hosts])

  return (
    <>
      <ProTable<HostItem>
        headerTitle="Hosts"
        size="large"
        search={false}
        options={false}
        pagination={false}
        columns={hostCols}
        rowKey={(rcd) => rcd.key}
        scroll={{ x: 'max-content' }}
        dataSource={hosts}
        editable={{
          editableKeys: editableKeys,
        }}
      />
    </>
  )
}

export default HostTable
