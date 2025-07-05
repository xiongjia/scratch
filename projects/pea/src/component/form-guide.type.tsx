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
  closeAble?: boolean
  form?: FormInstance<T>
  children: ReactNode
}

export interface GuideStepFormProps {
  name: string
  initialValues?: Store
  children?: ReactNode
}

interface BaseStepsFormProp<T> {
  open?: boolean
  onClose?: (e: React.MouseEvent | React.KeyboardEvent) => void
  closeAble?: boolean

  form?: FormInstance<T>
  children: ReactNode
}

interface BaseStepFormProp {
  children?: ReactNode
  initialValues?: Store
}

function BaseStepForm<T>(prop: StepFormProps<T> & BaseStepFormProp) {
  const { children, initialValues } = prop
  const stepFormProp = omit(prop, ['children'])

  return (
    <StepsForm.StepForm<T>
      {...stepFormProp}
      colProps={{
        span: 8,
      }}
      layout="horizontal"
      labelCol={{
        style: { width: 200, whiteSpace: 'normal', textAlign: 'right' },
      }}
      initialValues={initialValues}
    >
      {children}
    </StepsForm.StepForm>
  )
}

function BaseStepsForm<T>(prop: StepsFormProps<T> & BaseStepsFormProp<T>) {
  const { children, closeAble, onClose, open, form } = prop
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
            <Form
              className="w-60 relative *:first:sticky *:first:top-0"
              form={form}
            >
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
            title={
              <Space size="large" direction="horizontal">
                Test
              </Space>
            }
            maskClosable={false}
            destroyOnHidden={true}
            closeIcon={closeAble}
            size="large"
            width="100%"
            style={{ minWidth: '1100px' }}
            placement="left"
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
            {dom}
          </Drawer>
        )
      }}
    >
      {children}
    </StepsForm>
  )
}

export { BaseStepsForm, BaseStepForm }
