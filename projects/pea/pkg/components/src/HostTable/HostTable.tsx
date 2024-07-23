import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components'
import React, { useRef } from 'react'

type HostType = {
  ipAddr: string
  hostname: string
}

const HostColumns: ProColumns<HostType>[] = [
  {
    dataIndex: 'ipAddr',
    title: 'IP Address',
  },
  {
    dataIndex: 'hostname',
    title: 'Hostname',
  },
  {
    title: 'option',
    valueType: 'option',
    key: 'option',
    render: (text, record) => [
      <a
        key="remove"
        onClick={() => {
          console.log('remove ' + record.ipAddr)
        }}
      >
        REMOVE
      </a>,
    ],
  },
]

const testSrc: HostType[] = [
  {
    ipAddr: '1.1.1.1',
    hostname: 'host1',
  },
  {
    ipAddr: '1.1.1.2',
    hostname: 'host1',
  },
  {
    ipAddr: '1.1.1.3',
    hostname: 'host1',
  },
]

const HostTable = () => {
  const actionRef = useRef<ActionType>()

  return (
    <>
      <ProTable<HostType>
        columns={HostColumns}
        actionRef={actionRef}
        search={false}
        options={false}
        pagination={false}
        rowKey={(record) => record.ipAddr}
        request={async () => {
          return {
            data: testSrc,
            success: true,
            total: testSrc.length,
          }
        }}
        editable={{
          type: 'multiple',
        }}
        rowSelection={{
          defaultSelectedRowKeys: ['ipAddr'],
          onChange: (selectedRowKeys) => {
            console.log('select keys = ' + selectedRowKeys)
          },
        }}
        tableAlertRender={false}
      />
    </>
  )
}

export default HostTable
