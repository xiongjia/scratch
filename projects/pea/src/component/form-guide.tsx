import { omit } from 'lodash-es'
import { useEffect } from 'react'
import {
  BaseStepForm,
  BaseStepsForm,
  type GuideFormProps,
  type GuideStepFormProps,
} from './form-guide.type'

function GuideStepForm<T>(prop: GuideStepFormProps<T>) {
  const formProp = omit(prop, ['children', 'formRef'])
  return (
    <BaseStepForm<T> {...formProp} formRef={prop.formRef}>
      {prop.children}
    </BaseStepForm>
  )
}

function GuideForm<T>(prop: GuideFormProps<T>) {
  const { currentStepNum, onCurrentChange } = prop
  const formProp = omit(prop, [
    'children',
    'open',
    'currentStepNum',
    'onCurrentChange',
  ])

  useEffect(() => {
    onCurrentChange?.(currentStepNum ?? 0)
  }, [currentStepNum, onCurrentChange])

  return (
    <BaseStepsForm<T>
      {...formProp}
      open={prop.open}
      current={currentStepNum}
      onCurrentChange={(current: number) => {
        onCurrentChange?.(current)
      }}
    >
      {prop.children}
    </BaseStepsForm>
  )
}

export { GuideForm, GuideStepForm }
