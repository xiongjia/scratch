import { useRef } from 'react'
import type { ProFormInstance } from '@ant-design/pro-components'
import { ProCard, StepsForm, ProFormText } from '@ant-design/pro-components'

type Step1FormData = {
  step1Usrename: string
}

type Step2FormData = {
  steup2Username: string
}

type CompleteFormData = Step1FormData & Step2FormData

const TestSteps = () => {
  const formRef = useRef<ProFormInstance>(null)

  return (
    <>
      <ProCard>
        <StepsForm<CompleteFormData>
          stepsProps={{
            direction: 'vertical',
          }}
          formRef={formRef}
          onFinish={async (val: CompleteFormData): Promise<boolean> => {
            console.log('steps finsh', val)
            return true
          }}
        >
          <StepsForm.StepForm<Step1FormData>
            title="step1"
            stepProps={{
              description: 'step1 desc',
            }}
            initialValues={{
              step1Usrename: 'step1 init',
            }}
            onFinish={async (val: Step1FormData): Promise<boolean> => {
              console.log('step1 finish', val)
              return true
            }}
          >
            <ProFormText
              name="step1Usrename"
              label="实验名称"
              width="md"
              tooltip="最长为 24 位，用于标定的唯一 id"
              placeholder="请输入名称"
              rules={[{ required: true }]}
            />
          </StepsForm.StepForm>
          <StepsForm.StepForm<Step2FormData>
            title="step2"
            stepProps={{
              description: 'step2 desc',
            }}
            initialValues={{
              steup2Username: 'step2 init',
            }}
            onFinish={async (val: Step2FormData): Promise<boolean> => {
              console.log('step 2 finsh', val)
              return true
            }}
          >
            <ProFormText
              name="steup2Username"
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
