import { Input, type InputProps } from 'antd'

export interface InputIpProps extends InputProps {}

export const InputIp = (props: InputIpProps) => {
  const { ...restProps } = props
  return <Input {...restProps} status="" />
}
