import { omit } from 'lodash-es'
import {
  BaseStepForm,
  BaseStepsForm,
  type GuideFormProps,
  type GuideStepFormProps,
} from './form-guide.type'

function GuideStepForm<T>(prop: GuideStepFormProps<T>) {
  const formProp = omit(prop, ['children'])
  return <BaseStepForm<T> {...formProp}>{prop.children}</BaseStepForm>
}

function GuideForm<T>(prop: GuideFormProps<T>) {
  const formProp = omit(prop, ['children', 'open'])
  return (
    <BaseStepsForm<T> {...formProp} open={prop.open}>
      {prop.children}
    </BaseStepsForm>
  )
}

export { GuideForm, GuideStepForm }
