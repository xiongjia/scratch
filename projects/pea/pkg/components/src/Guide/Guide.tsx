import { StepsForm, StepsFormProps } from '@ant-design/pro-components'
import { Drawer } from 'antd'
import React, { ReactNode } from 'react'

export type GuideStep = {
  key: string
  title: string
  name: string
  content: ReactNode
  onLoad?: () => void
  onFinish?: () => Promise<boolean | void>
}

function Guide<T>(
  props: StepsFormProps<T> & {
    steps: Array<GuideStep>
  },
) {
  const { steps } = props

  return (
    <StepsForm<{ name: string }>
      containerStyle={{
        margin: 0,
        minHeight: '100%',
      }}
      stepsProps={{
        direction: 'vertical',
        className: 'sticky top-0',
        style: {
          minHeight: 'calc(100vh - 160px)',
          maxHeight: 'calc(100vh - 160px)',
        },
      }}
      formProps={{ labelAlign: 'left' }}
      onCurrentChange={(current: number) => {
        steps[current]?.onLoad?.()
      }}
      stepsFormRender={(dom, submitter) => {
        return (
          <Drawer
            title={'安装 xxx '}
            maskClosable={false}
            size="default"
            width="100%"
            placement="left"
            destroyOnClose={true}
            closeIcon={false}
            open={true}
            footer={
              <div
                style={{
                  display: 'flex',
                  gap: '8px',
                  justifyContent: 'end',
                  minWidth: '520px',
                }}
              >
                {submitter}
              </div>
            }
          >
            {dom}
          </Drawer>
        )
      }}
    >
      {steps?.map((step) => {
        return (
          <StepsForm.StepForm<T>
            key={step.key}
            name={step.name}
            title={step.title}
            onFinish={step.onFinish}
          >
            {step.content}
          </StepsForm.StepForm>
        )
      })}
    </StepsForm>
  )
}

export default Guide
