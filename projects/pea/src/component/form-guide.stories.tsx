import { useState } from 'react'
import '@ant-design/v5-patch-for-react-19'
import { ProCard, ProFormText } from '@ant-design/pro-components'
import { Button, Form, Space } from 'antd'
import { GuideForm, GuideStepForm } from './'

type pgStep1Data = {
  username1: string
  password1: string
}

type pgStep2Data = {
  username2: string
  password2: string
}

type guideData = pgStep1Data | pgStep2Data

const pageStep1 = () => {
  return (
    <GuideStepForm<pgStep1Data>
      name="step1"
      initialValues={{
        username1: 'user1',
        password1: 'passowrd1',
      }}
    >
      <ProCard
        title={
          <div className="inline-flex">
            <span className="font-bold">Page1</span>
          </div>
        }
        bordered
      >
        <Space direction="vertical">
          <ProFormText
            width="sm"
            name="pg1Name"
            label="usernma1"
            rules={[{ required: true }]}
            required={true}
          />
          <ProFormText.Password
            width="sm"
            name="pg1Password"
            label="password1"
            required={true}
            rules={[{ required: true }]}
          />
        </Space>
      </ProCard>
    </GuideStepForm>
  )
}

const pageStep2 = () => {
  return (
    <GuideStepForm<pgStep2Data> name="step2">
      <ProCard
        title={
          <div className="inline-flex">
            <span className="font-bold">Page1</span>
          </div>
        }
        bordered
      >
        <Space direction="vertical">
          <ProFormText
            width="sm"
            name="pg2Name"
            label="username2"
            rules={[{ required: true }]}
            required={true}
          />
          <ProFormText.Password
            width="sm"
            name="pg2Password"
            label="password2"
            required={true}
            rules={[{ required: true }]}
          />
        </Space>
      </ProCard>
    </GuideStepForm>
  )
}

export const StoryFormGuide = () => {
  const [open, setOpen] = useState(false)
  const [form] = Form.useForm<guideData>()

  const showGuideForm = () => {
    setOpen(true)
  }

  const onGuideFormClose = () => {
    setOpen(false)
  }

  return (
    <>
      <Space>
        <Button onClick={showGuideForm}>Show Guide</Button>
      </Space>
      <GuideForm<guideData>
        closeAble={true}
        open={open}
        onClose={onGuideFormClose}
        form={form}
      >
        {pageStep1()}
        {pageStep2()}
      </GuideForm>
    </>
  )
}
