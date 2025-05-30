import { Form, FormItemProps, Input } from 'antd'

function FormItemIp<Values = any>(props: FormItemProps<Values>) {
  const { ...itemProps } = props

  return (
    <Form.Item {...itemProps}>
      <Input placeholder="test" status="" />
    </Form.Item>
  )
}

export default FormItemIp
