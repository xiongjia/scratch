import { ProCard, ProFormText } from '@ant-design/pro-components'
import { Space } from 'antd'

export const Guide = ({ title }: { title: string }) => {
  return (
    <ProCard
      title={
        <div className="inline-flex">
          <span className="font-bold">{title}</span>
        </div>
      }
      bordered
    >
      <Space direction="horizontal" key="3">
        <ProFormText
          width="sm"
          name={'beginAddr'}
          label="IP Address"
          initialValue={'1'}
          required={true}
        />
        <ProFormText
          width="sm"
          name={'endAddr'}
          label="IP Address"
          initialValue={'2'}
          required={true}
        />
      </Space>
    </ProCard>
  )
}
