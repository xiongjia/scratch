import { Space, Button, Form, Input } from 'antd'
import { HostTable, type HostItem, TestSteps } from '&/component'

import {
  useQuery,
  useQueryClient,
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'

const queryClient = new QueryClient()

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

const Todos = () => {
  const queryClient = useQueryClient()
  const query = useQuery({ queryKey: ['todos'], queryFn: getTodos })
  return (
    <>
      <div>Add Todo</div>
    </>
  )
}

const Playground = () => {
  const [form] = Form.useForm()

  const onFinish = (val: { user: string }) => {
    console.log('Received user:', val)
    form.setFieldValue('user', 'update1')
  }

  return (
    <>
      <Space direction="vertical" size="large">
        <p>Playground</p>

        <Form
          form={form}
          onFinish={onFinish}
          name="basic"
          initialValues={{ user: '123' }}
          layout="vertical"
        >
          <Form.Item
            label="user"
            name="user"
            rules={[{ required: true, message: 'Please input your username!' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit">
              Submit
            </Button>
          </Form.Item>
        </Form>

        <HostTable hosts={hosts} />

        <QueryClientProvider client={queryClient}>
          <Todos />
        </QueryClientProvider>

        <Button
          size="large"
          onClick={() => {
            console.log('test')
          }}
        >
          Test
        </Button>

        <TestSteps />
      </Space>
    </>
  )
}

export default Playground
