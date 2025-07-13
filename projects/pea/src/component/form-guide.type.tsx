import {
  type StepFormProps,
  StepsForm,
  type StepsFormProps,
} from '@ant-design/pro-components'
import { Drawer, Form, type FormInstance, Space } from 'antd'
import type { Store } from 'antd/es/form/interface'
import { omit } from 'lodash-es'
import type { ReactNode } from 'react'

export interface GuideFormProps<T> {
  open?: boolean
  onClose?: (e: React.MouseEvent | React.KeyboardEvent) => void
  onFinish?: (formData: T) => Promise<boolean>
  onCurrentChange?: (current: number) => void
  closeAble?: boolean
  children: ReactNode
  currentStepNum?: number
}

export interface GuideStepFormProps<T> {
  children?: ReactNode

  name: string
  title: string
  formRef?: React.RefObject<FormInstance<T> | null>

  onFinish?: (formData: T) => Promise<boolean>
  onStepShow?: () => void
}

export interface GuideStepFormRef {
  actionStepShow: () => void
}

interface BaseStepsFormProp<T> {
  open?: boolean
  onClose?: (e: React.MouseEvent | React.KeyboardEvent) => void
  closeAble?: boolean
  onFinish?: (formData: T) => Promise<boolean>
  children: ReactNode
  initialValues?: Store
}

interface BaseStepFormProp<T> {
  children?: ReactNode
  formRef?: React.RefObject<FormInstance<T> | null>
  title?: string
}

function BaseStepForm<T>(prop: StepFormProps<T> & BaseStepFormProp<T>) {
  const stepFormProp = omit(prop, ['children', 'formRef'])
  return (
    <StepsForm.StepForm<T>
      {...stepFormProp}
      formRef={prop.formRef}
      colProps={{
        span: 8,
      }}
      layout="horizontal"
      labelCol={{
        style: { width: 200, whiteSpace: 'normal', textAlign: 'right' },
      }}
    >
      {prop.children}
    </StepsForm.StepForm>
  )
}

function BaseStepsForm<T>(prop: StepsFormProps<T> & BaseStepsFormProp<T>) {
  const stepsFormProp = omit(prop, [
    'children',
    'initialValues',
    'open',
    'onClose',
    'closeAble',
  ])

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
      formProps={{ labelAlign: 'left', initialValues: prop.initialValues }}
      stepsFormRender={(dom, submitter) => {
        return (
          <Drawer
            title={
              <Space size="large" direction="horizontal">
                Test
              </Space>
            }
            closeIcon={prop.closeAble}
            onClose={prop.onClose}
            open={prop.open}
            maskClosable={false}
            destroyOnHidden={true}
            size="large"
            width="100%"
            style={{ minWidth: '1100px' }}
            placement="left"
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
            {dom}
          </Drawer>
        )
      }}
    >
      {prop.children}
    </StepsForm>
  )
}

export { BaseStepsForm, BaseStepForm }
