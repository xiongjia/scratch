import { StepsForm, StepsFormProps } from '@ant-design/pro-components'
import { Drawer } from 'antd'
import i18next from 'i18next'
import React, { ReactNode } from 'react'
import { I18nextProvider } from 'react-i18next'

i18next.init({
  interpolation: { escapeValue: false },
  lng: 'chs',
  resources: {
    en: {
      common: {
        title: 'en-title {{ a }}, {{ b }}',
      },
    },
    chs: {
      common: {
        title: 'chs-title {{ b }}, {{ a }}',
      },
    },
  },
})
i18next.setDefaultNamespace('common')

export type GuideStep = {
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
    <I18nextProvider i18n={i18next}>
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
              key={step.name}
              name={step.name}
              title={step.title}
              onFinish={step.onFinish}
            >
              {step.content}
            </StepsForm.StepForm>
          )
        })}
      </StepsForm>
    </I18nextProvider>
  )
}

export default Guide
