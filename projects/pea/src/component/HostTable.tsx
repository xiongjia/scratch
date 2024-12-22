import type { ProColumns } from '@ant-design/pro-components'
import { ProTable } from '@ant-design/pro-components'

type HostItem = {
  key: string
  name: string
}

const hostCols: ProColumns<HostItem>[] = [
  {
    title: 'name',
    key: 'name',
    width: 250,
    dataIndex: 'name',
    fixed: 'left',
  },
]

const items: HostItem[] = [
  { key: 'host1', name: 'host1' },
  { key: 'host2', name: 'host2' },
]

const HostTable = () => {
  return (
    <>
      <ProTable<HostItem>
        headerTitle="Hosts"
        size="large"
        columns={hostCols}
        rowKey={(rcd) => rcd.key}
        scroll={{ x: 'max-content' }}
        dataSource={items}
      />
    </>
  )
}

export default HostTable
