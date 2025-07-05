import { useState } from 'react'
import '@ant-design/v5-patch-for-react-19'
import { Button, Space } from 'antd'

import { GuideForm, GuideStepForm } from './'

export const StoryFormGuide = () => {
  const [open, setOpen] = useState(false)

  const showGuideForm = () => {
    setOpen(true)
  }

  const onGuideFormClose = () => {
    setOpen(false)
  }

  return (
    <>
      <Space>
        <Button onClick={showGuideForm}>Show Guide</Button>
      </Space>
      <GuideForm closeAble={true} open={open} onClose={onGuideFormClose}>
        <GuideStepForm name="step1" />
        <GuideStepForm name="step2" />
      </GuideForm>
    </>
  )
}
