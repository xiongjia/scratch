import {
  ProColumns,
  ProFormText,
  ProTable,
  StepsForm,
} from '@ant-design/pro-components'
import React from 'react'

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
      status: 'active',
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

const GuideStep2 = () => {
  return (
    <>
      <StepsForm.StepForm<{
        name: string
      }>
        name="base"
        title="Step 1"
        stepProps={{
          description: '这是步骤1',
        }}
        onFinish={async () => {
          console.log('finish')
          return true
        }}
      >
        <ProFormText
          name="name"
          label="实验名称"
          width="md"
          tooltip="最长为 24 位，用于标定的唯一 id"
          placeholder="请输入名称"
          rules={[{ required: true }]}
        />

        <ProTable<TableListItem>
          columns={columns}
          rowKey="key"
          dataSource={tableListDataSource}
        />
      </StepsForm.StepForm>
    </>
  )
}

export default GuideStep2
