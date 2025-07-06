import { forwardRef, useImperativeHandle, useRef, useState } from 'react'
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

interface PageRef {
  onCurrentStep: () => void
}

const PageStep1 = forwardRef<PageRef>((prop, ref) => {
  // const formRef = ProForm.useFormInstance()

  useImperativeHandle(ref, () => ({
    onCurrentStep: () => {
      console.log('on page1 setp', prop)
    },
  }))

  return (
    <GuideStepForm<pgStep1Data>
      name="step1"
      initialValues={{
        pg1Name: 'u1',
      }}
      onFinish={async (formData: pgStep1Data): Promise<boolean> => {
        console.log(formData)
        return true
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
})

const PageStep2 = forwardRef<PageRef>((prop, ref) => {
  useImperativeHandle(ref, () => ({
    onCurrentStep: () => {
      console.log('on page2 setp', prop)
    },
  }))
  return (
    <GuideStepForm<pgStep2Data>
      name="step2"
      initialValues={{
        pg2Name: 'u2',
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
})

export const StoryFormGuide = () => {
  const [open, setOpen] = useState(false)
  const [form] = Form.useForm<guideData>()
  const pgStep1Ref = useRef<PageRef>(null)
  const pgStep2Ref = useRef<PageRef>(null)

  const showGuideForm = () => {
    setOpen(true)
  }

  const onGuideFormClose = () => {
    setOpen(false)
  }

  const onSetpsFinish = async (formData: guideData): Promise<boolean> => {
    console.log('steps finish', formData)
    setOpen(false)
    return true
  }

  const onStepChange = (currentStep: number) => {
    console.log('current setp', currentStep)
    if (currentStep === 0) {
      pgStep1Ref?.current?.onCurrentStep?.()
    } else if (currentStep === 1) {
      pgStep2Ref?.current?.onCurrentStep?.()
    }
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
        onFinish={onSetpsFinish}
        onCurrentChange={onStepChange}
        form={form}
        initialValues={{
          pg1Name: 'u1',
          pg2Name: 'u2',
        }}
      >
        <PageStep1 ref={pgStep1Ref} />
        <PageStep2 ref={pgStep2Ref} />
      </GuideForm>
    </>
  )
}
