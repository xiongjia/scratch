import { Button } from 'antd'

const TestButton = () => {
  return (
    <>
      <Button
        onClick={() => {
          console.log('test')
        }}
      >
        Test Button
      </Button>
    </>
  )
}

export default TestButton
