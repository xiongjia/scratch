import {
  BaseStepForm,
  BaseStepsForm,
  type GuideFormProps,
  type GuideStepFormProps,
} from './form-guide.type'

function GuideStepForm<T>({
  children,
  name,
  initialValues,
}: GuideStepFormProps) {
  return (
    <BaseStepForm<T> name={name} initialValues={initialValues}>
      {children}
    </BaseStepForm>
  )
}

function GuideForm<T>({
  children,
  closeAble,
  open,
  onClose,
  form,
}: GuideFormProps<T>) {
  return (
    <BaseStepsForm<T>
      open={open}
      onClose={onClose}
      closeAble={closeAble}
      onCurrentChange={(current: number) => {
        console.log('current step: ', current)
      }}
      onFinish={async (val: T): Promise<boolean> => {
        console.log('val: ', val)
        return true
      }}
      form={form}
    >
      {children}
    </BaseStepsForm>
  )
}

export { GuideForm, GuideStepForm }
