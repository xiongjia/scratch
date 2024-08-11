import { ProColumns, ProTable } from '@ant-design/pro-components'
import { Modal, Space } from 'antd'
import React, { useContext, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { GuideContext } from './GuideContext'
import { useAppDispatch, useAppSelector } from './store/store'

export type TableListItem = {
  key: number
  name: string
  progress: number
}

const ProcessMap = {
  close: 'normal',
  running: 'active',
  online: 'success',
  error: 'exception',
} as const

const columns: ProColumns<TableListItem>[] = [
  {
    title: '应用名称',
    width: 120,
    dataIndex: 'name',
    fixed: 'left',
    render: (_) => <a>{_}</a>,
  },
  {
    title: '执行进度',
    dataIndex: 'progress',
    valueType: (item) => ({
      type: 'progress',
      status: 'success',
      showSymbol: false,
      showColor: false,
    }),
  },
]

const tableListDataSource: TableListItem[] = []
tableListDataSource.push({
  key: 1,
  name: 'AppName-1',
  progress: 50,
})

tableListDataSource.push({
  key: 2,
  name: 'AppName-2',
  progress: 100,
})

const StepHostsSelector = () => {
  const ctx = useContext(GuideContext)
  const dispatch = useAppDispatch()
  const hosts = useAppSelector((state) => state.host.hosts)
  const { t, i18n } = useTranslation('common')

  const [modal, contextHolder] = Modal.useModal()

  useEffect(() => {
    i18n.changeLanguage('en')
    console.log('hosts = ' + JSON.stringify(hosts))
  }, [hosts])

  return (
    <Space direction="vertical">
      <p>{t('title', { a: 'a', b: 'b' })}</p>
      <p> step select host </p>
      <button
        type="button"
        onClick={() => {
          console.log('btn click')
          // dispatch(actions.loadHostList())
          modal.success({
            title: 'title',
            content: 'msg',
          })
        }}
      >
        Test btn
      </button>

      <ProTable<TableListItem>
        columns={columns}
        rowKey="key"
        dataSource={tableListDataSource}
      />

      {contextHolder}
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
