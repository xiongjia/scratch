import { ProCard, ProFormText } from '@ant-design/pro-components'
import { Space } from 'antd'
import React from 'react'

const IpRangeForm = (props: {
  beginAddrField?: string
  endAddrField?: string
  beginAddrInitValue?: string
  endAddrInitValue?: string
}) => {
  return (
    <ProCard
      title={
        <div className="inline-flex">
          <span className="font-bold">{'IP 范围'}</span>
        </div>
      }
      bordered
    >
      <Space direction="horizontal" key="3">
        <ProFormText
          width="sm"
          name={props?.beginAddrField ?? 'beginAddr'}
          label="起始 IP"
          initialValue={props?.beginAddrInitValue ?? ''}
          required={true}
        />
        <ProFormText
          width="sm"
          name={props?.endAddrField ?? 'endAddr'}
          label="终止 IP"
          initialValue={props?.endAddrInitValue ?? ''}
          required={true}
        />
      </Space>
    </ProCard>
  )
}

export default IpRangeForm
