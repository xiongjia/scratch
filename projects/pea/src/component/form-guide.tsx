import {
  BaseStepForm,
  BaseStepsForm,
  type GuideFormProps,
  type GuideStepFormProps,
} from './form-guide.type'

function GuideStepForm<T>({ children, name }: GuideStepFormProps) {
  return <BaseStepForm<T> name={name}>{children}</BaseStepForm>
}

function GuideForm<T>({ children, closeAble, open, onClose }: GuideFormProps) {
  return (
    <BaseStepsForm<T>
      open={open}
      onClose={onClose}
      closeAble={closeAble}
      onCurrentChange={(current: number) => {
        console.log('current step: ', current)
      }}
    >
      {children}
    </BaseStepsForm>
  )
}

export { GuideForm, GuideStepForm }
