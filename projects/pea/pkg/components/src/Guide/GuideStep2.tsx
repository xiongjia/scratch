import { ProFormText, StepsForm } from '@ant-design/pro-components'
import React from 'react'

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
      </StepsForm.StepForm>
    </>
  )
}

export default GuideStep2
