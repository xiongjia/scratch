import { useRef } from 'react'
import type { ProFormInstance } from '@ant-design/pro-components'
import { ProCard, StepsForm, ProFormText } from '@ant-design/pro-components'

const TestSteps = () => {
  const formRef = useRef<ProFormInstance>(null)

  return (
    <>
      <ProCard>
        <StepsForm
          stepsProps={{
            direction: 'vertical',
          }}
          formRef={formRef}
        >
          <StepsForm.StepForm
            title="step1"
            stepProps={{
              description: 'step1 desc',
            }}
            onFinish={async () => {
              console.log(formRef.current?.getFieldsValue())
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
          <StepsForm.StepForm
            title="step2"
            stepProps={{
              description: 'step2 desc',
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
        </StepsForm>
      </ProCard>
    </>
  )
}

export default TestSteps
