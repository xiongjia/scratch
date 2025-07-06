import { omit } from 'lodash-es'
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
  const formProp = omit(prop, ['children', 'open', 'formRef'])
  return (
    <BaseStepsForm<T> {...formProp} open={prop.open} formRef={prop.formRef}>
      {prop.children}
    </BaseStepsForm>
  )
}

export { GuideForm, GuideStepForm }
