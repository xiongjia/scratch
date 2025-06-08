import { ProForm, ProFormText } from '@ant-design/pro-components'
import { Button } from 'antd'

type TestFormData = {
  username: string
}

const FormTest = () => {
  return (
    <>
      <ProForm<TestFormData>
        onFinish={async (val: TestFormData): Promise<boolean> => {
          console.log('finsh proform', val)
          return true
        }}
        onValuesChange={(val: TestFormData) => {
          console.log('proform on change', val)
        }}
        initialValues={{
          username: '123',
        }}
        layout="inline"
        submitter={{
          resetButtonProps: { style: { display: 'none' } },
          submitButtonProps: {},
          render: (props): React.ReactNode[] => {
            return [
              <Button key="submit" onClick={() => props.form?.submit?.()}>
                Ok submit 1
              </Button>,
            ]
          },
        }}
      >
        <ProFormText name={'username'} />
      </ProForm>
    </>
  )
}

export default FormTest
