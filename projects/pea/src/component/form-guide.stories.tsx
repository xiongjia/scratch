import {
  forwardRef,
  useEffect,
  useImperativeHandle,
  useRef,
  useState,
} from 'react'
import '@ant-design/v5-patch-for-react-19'
import {
  ProCard,
  type ProFormInstance,
  ProFormText,
} from '@ant-design/pro-components'
import { Button, Space } from 'antd'
import { GuideForm, GuideStepForm } from './'

type pgStep1Data = {
  pg1Name: string
  pg1Password: string
}

type pgStep2Data = {
  pg2Name: string
  pg2Password: string
}

type guideData = pgStep1Data | pgStep2Data

interface PageRef {
  onCurrentStep: () => void
}

const PageStep1 = forwardRef<PageRef>((prop, ref) => {
  const formRef = useRef<ProFormInstance<pgStep1Data>>(null)
  useImperativeHandle(ref, () => ({
    onCurrentStep: () => {
      console.log('on page1 step', prop)
    },
  }))

  return (
    <GuideStepForm<pgStep1Data>
      formRef={formRef}
      name="step1"
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
  const formRef = useRef<ProFormInstance<pgStep2Data>>(null)
  useImperativeHandle(ref, () => ({
    onCurrentStep: () => {
      formRef.current?.setFieldValue('pg2Name', 'init')
      console.log('on page2 setp', prop, formRef.current?.getFieldsValue())
    },
  }))

  useEffect(() => {
    console.log('effect pg2', formRef.current?.getFieldsValue())
  })

  const onBtnClick = () => {
    console.log('on click', formRef.current?.getFieldsValue())
    formRef.current?.setFieldValue('pg2Name', 'data1')
  }

  return (
    <GuideStepForm<pgStep2Data>
      name="step2"
      formRef={formRef}
      onFinish={async (formData: pgStep2Data): Promise<boolean> => {
        console.log(formData)
        return true
      }}
    >
      <Button onClick={onBtnClick}>Test</Button>
      <ProCard
        title={
          <div className="inline-flex">
            <span className="font-bold">Page2</span>
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
    // setOpen(false)
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
      >
        <PageStep1 ref={pgStep1Ref} />
        <PageStep2 ref={pgStep2Ref} />
      </GuideForm>
    </>
  )
}
