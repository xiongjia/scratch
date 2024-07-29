import { ProCard, ProFormText } from '@ant-design/pro-components'
import { Space } from 'antd'
import React from 'react'

const SshAccountForm = (props: {
  sshUsernameField?: string
  sshPasswordField?: string
  sshUsernameInitValue?: string
  sshPasswordInitValue?: string
}) => {
  return (
    <ProCard
      title={
        <div className="inline-flex">
          <span className="font-bold">{'连接系统的用户名和密码'}</span>
        </div>
      }
      bordered
    >
      <Space direction="vertical">
        <ProFormText
          width="sm"
          name={props?.sshUsernameField ?? 'sshUsername'}
          label="用户名"
          initialValue={props?.sshUsernameInitValue ?? ''}
          required={true}
        />
        <ProFormText.Password
          width="sm"
          name={props?.sshPasswordField ?? 'sshPassword'}
          label="密码"
          initialValue={props?.sshPasswordInitValue ?? ''}
          required={true}
        />
      </Space>
    </ProCard>
  )
}

export default SshAccountForm
