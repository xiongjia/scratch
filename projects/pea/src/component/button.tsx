import { Button } from 'antd'

const TestButton = () => {
  return (
    <Button
      className="bg-blue-500 text-white px-4 py-2 rounded"
      onClick={() => console.log('test')}
    >
      Test
    </Button>
  )
}

export default TestButton
