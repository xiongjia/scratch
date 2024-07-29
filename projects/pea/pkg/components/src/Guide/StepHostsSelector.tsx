import { Space } from 'antd'
import React, { useContext } from 'react'
import { GuideContext } from './GuideContext'

const StepHostsSelector = () => {
  const ctx = useContext(GuideContext)

  return (
    <Space direction="vertical">
      <p> step select host </p>
      <button
        type="button"
        onClick={() => {
          console.log('state: ' + JSON.stringify(ctx))
        }}
      >
        Test btn
      </button>
    </Space>
  )
}

export default {
  name: 'stepHostsSelector',
  title: '选择主机',
  onFinish: async (): Promise<boolean> => {
    console.log('stepHostsSelector finished')
    return true
  },
  onLoad: () => {
    console.log('stepHostsSelector loaded')
  },
  content: <StepHostsSelector />,
}
