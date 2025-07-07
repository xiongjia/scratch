import { omit } from 'lodash-es'
import {
  type ForwardedRef,
  forwardRef,
  type ReactElement,
  useEffect,
  useImperativeHandle,
} from 'react'
import {
  BaseStepForm,
  BaseStepsForm,
  type GuideFormProps,
  type GuideStepFormProps,
  type GuideStepFormRef,
} from './form-guide.type'

export const GuideStepForm = forwardRef(
  <T,>(prop: GuideStepFormProps<T>, ref: ForwardedRef<GuideStepFormRef>) => {
    const formProp = omit(prop, ['children', 'formRef', 'ref', 'onStepShow'])
    useImperativeHandle(ref, () => ({
      actionStepShow: () => {
        prop?.onStepShow?.()
      },
    }))
    return (
      <BaseStepForm<T> {...formProp} formRef={prop.formRef}>
        {prop.children}
      </BaseStepForm>
    )
  },
) as <T>(
  props: GuideStepFormProps<T> & { ref?: ForwardedRef<GuideStepFormRef> },
) => ReactElement

export const GuideForm = <T,>(prop: GuideFormProps<T>) => {
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
