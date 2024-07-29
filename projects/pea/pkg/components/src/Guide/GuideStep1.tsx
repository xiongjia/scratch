import { ProFormText, StepsForm } from '@ant-design/pro-components'
import { Button } from 'antd'
import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { RootState, Todo, TodoActionsTypes, actions } from './store'

type Dispatch = <TReturnType>(action: TodoActionsTypes) => TReturnType
const useTypedDispatch = () => useDispatch<Dispatch>()

const GuideStep1 = () => {
  const srcData = useSelector((state: RootState) => state.todo.data)
  useEffect(() => {}, [srcData])
  const dispatch = useTypedDispatch()

  const test1Click = () => {
    console.log('test click')
    dispatch(
      actions.createTodo({ id: 1, description: 'test1', checked: false }),
    )
  }

  return (
    <>
      <StepsForm.StepForm<{
        name: string
      }>
        name="base"
        title="Step 1"
        stepProps={{
          description: '这是步骤1',
        }}
        onFinish={async () => {
          console.log('finish')
          return true
        }}
      >
        <ProFormText
          name="name"
          label="实验名称"
          width="md"
          tooltip="最长为 24 位，用于标定的唯一 id"
          placeholder="请输入名称"
          rules={[{ required: true }]}
        />

        <ul>
          {srcData.map((todo: Todo) => (
            <li key={todo.id}>
              <div>
                <span>{`${todo.id} ${todo.checked} - ${todo.description}`}</span>
              </div>
            </li>
          ))}
        </ul>

        <Button onClick={test1Click}>Test1</Button>
      </StepsForm.StepForm>
    </>
  )
}

export default GuideStep1
