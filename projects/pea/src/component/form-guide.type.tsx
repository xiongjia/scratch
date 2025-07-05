import {
  type StepFormProps,
  StepsForm,
  type StepsFormProps,
} from '@ant-design/pro-components'
import { Drawer, Form, Space } from 'antd'
import { omit } from 'lodash-es'
import type { ReactNode } from 'react'

export interface GuideFormProps {
  open?: boolean
  onClose?: (e: React.MouseEvent | React.KeyboardEvent) => void
  closeAble?: boolean

  children: ReactNode
}
export interface GuideStepFormProps {
  name: string
  children?: ReactNode
}

interface BaseStepsFormProp {
  open?: boolean
  onClose?: (e: React.MouseEvent | React.KeyboardEvent) => void
  closeAble?: boolean

  children: ReactNode
}

interface BaseStepFormProp {
  children?: ReactNode
}

function BaseStepForm<T>(prop: StepFormProps<T> & BaseStepFormProp) {
  const { children } = prop
  const stepFormProp = omit(prop, ['children'])

  return (
    <StepsForm.StepForm<T> {...stepFormProp}>{children}</StepsForm.StepForm>
  )
}

function BaseStepsForm<T>(prop: StepsFormProps<T> & BaseStepsFormProp) {
  const { children, closeAble, onClose, open } = prop
  const stepsFormProp = omit(prop, ['children'])
  return (
    <StepsForm<T>
      {...stepsFormProp}
      stepsProps={{
        direction: 'vertical',
        className: 'sticky top-0',
        style: {
          minHeight: 'calc(100vh - 160px)',
          maxHeight: 'calc(100vh - 160px)',
        },
      }}
      containerStyle={{
        display: 'grid',
        margin: 0,
        minHeight: '100%',
        minWidth: '800px',
      }}
      layoutRender={(layoutDom) => {
        return (
          <div className="flex">
            <Form className="w-60 relative *:first:sticky *:first:top-0">
              {layoutDom.stepsDom}
            </Form>
            <div className="flex-1">{layoutDom.formDom}</div>
          </div>
        )
      }}
      formProps={{ labelAlign: 'left' }}
      stepsFormRender={(dom, submitter) => {
        return (
          <Drawer
            title={<Space size="large" direction="horizontal"></Space>}
            maskClosable={false}
            size="large"
            width="100%"
            style={{ minWidth: '1100px' }}
            placement="left"
            closeIcon={closeAble}
            open={open}
            onClose={onClose}
            footer={
              <div
                style={{
                  display: 'flex',
                  gap: '8px',
                  justifyContent: 'end',
                }}
              >
                {submitter}
              </div>
            }
          >
            <div style={{}}>{dom}</div>
          </Drawer>
        )
      }}
    >
      {children}
    </StepsForm>
  )
}

export { BaseStepsForm, BaseStepForm }
