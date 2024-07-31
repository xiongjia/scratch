import { Space } from 'antd'
import React, { useContext, useEffect } from 'react'
import { GuideContext } from './GuideContext'
import { actions } from './store/ModuleHosts'
import { useAppDispatch, useAppSelector } from './store/store'

const StepHostsSelector = () => {
  const ctx = useContext(GuideContext)
  const dispatch = useAppDispatch()
  const hosts = useAppSelector((state) => state.host.hosts)

  useEffect(() => {
    console.log('hosts = ' + JSON.stringify(hosts))
  }, [hosts])

  return (
    <Space direction="vertical">
      <p> step select host </p>
      <button
        type="button"
        onClick={() => {
          console.log('btn click')
          dispatch(actions.loadHostList())
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
