import { Space } from 'antd'
import React, { useContext } from 'react'
import useSWR from 'swr'
import { GuideContext } from './GuideContext'
import IpRangeForm from './IpRangeForm'
import SshAccountForm from './SshAccountForm'

function delayTest(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

type MockData = {
  name: string
  data: string
}

const fetcher = async (args: string): Promise<MockData> => {
  console.log('load data ', args)
  await delayTest(1000 * 15)
  console.log('return data ', args)
  return { name: args, data: 'b' }
}

const StepHostInput = () => {
  const ctx = useContext(GuideContext)
  const { data, error, isLoading, isValidating } = useSWR<MockData>(
    '/api/user/1211',
    fetcher,
    {
      keepPreviousData: false,
    },
  )

  console.log('data', data)
  console.log('error', error)
  console.log('isLoading', isLoading)
  console.log('isValidating', isValidating)

  return (
    <Space direction="vertical">
      <IpRangeForm
        beginAddrInitValue={ctx?.beginAddr}
        endAddrInitValue={ctx?.endAddr}
      />
      <SshAccountForm
        sshUsernameInitValue={ctx?.sshUsername}
        sshPasswordInitValue={ctx?.sshPassword}
      />

      {isLoading ?? <div>loading...</div>}

      <div>
        hello [{data?.name}, {data?.data}]
      </div>
    </Space>
  )
}

export default {
  name: 'stepHostsInput',
  title: '扫描主机',
  onFinish: async (): Promise<boolean> => {
    console.log('stepHostsInput finished')
    return true
  },
  onLoad: () => {
    console.log('stepHostsInput loaded')
  },
  content: <StepHostInput />,
}
