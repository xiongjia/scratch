import { Button } from 'antd'
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from '@tanstack/react-query'

const queryClient = new QueryClient()

const wait = (sec: number) => {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log('Done waiting')
      resolve(sec)
    }, sec * 1000)
  })
}

const TestData = () => {
  const { error, data, isFetching } = useQuery({
    queryKey: ['todos'],
    queryFn: async () => {
      const testData = {
        a: 1,
      }
      return new Promise((resolve, reject) => {
        setTimeout(() => {
          console.log('Done waiting')
          resolve(JSON.stringify(testData))
        }, 10 * 1000)
      })
    },
  })

  if (isFetching) {
    console.log('isFetching')
  }
  if (error) {
    console.log('err', error)
  }
  return (
    <>
      <p>{isFetching ? 'fetching' : data}</p>
      <br />
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

const TestButton = () => {
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <TestData />
      </QueryClientProvider>
    </>
  )
}

export default TestButton
