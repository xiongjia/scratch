import { Space, Button } from 'antd'
import { HostTable, type HostItem } from '&/component'

const hosts: HostItem[] = [
  {
    key: 'host1',
    name: 'host1',
    nodeName: 'n1',
    status: 'running',
    outputLog: 'log123',
  },
  {
    key: 'host2',
    name: 'host2',
    nodeName: 'n2',
    status: 'stopped',
    outputLog: 'log123',
  },
]

const Playground = () => {
  return (
    <>
      <Space direction="vertical" size="large">
        <p>Playground</p>
        <HostTable hosts={hosts} />

        <Button
          size="large"
          onClick={() => {
            console.log('test')
          }}
        >
          Test
        </Button>
      </Space>
    </>
  )
}

export default Playground
