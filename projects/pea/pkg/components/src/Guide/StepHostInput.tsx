import { Space } from 'antd'
import React, { useContext } from 'react'
import { GuideContext } from './GuideContext'
import IpRangeForm from './IpRangeForm'
import SshAccountForm from './SshAccountForm'

const StepHostInput = () => {
  const ctx = useContext(GuideContext)

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
