import React from 'react'
import Guide, { GuideStep } from './Guide'

const makeSteps = (): GuideStep[] => {
  const steps: GuideStep[] = []

  steps.push({
    key: 'step1',
    title: 'step1 - title',
    name: 'step1 - name',
    content: <p>step 1</p>,
    onFinish: async (): Promise<boolean> => {
      console.log('setp1 finished')
      return true
    },
  })

  steps.push({
    key: 'step2',
    title: 'step2 - title',
    name: 'step2 - name',
    content: <p>step 2</p>,
    onFinish: async (): Promise<boolean> => {
      console.log('setp2 finished')
      return true
    },
  })

  steps.push({
    key: 'step3',
    title: 'step3 - title',
    name: 'step3 - name',
    content: <p>step 3</p>,
    onFinish: async (): Promise<boolean> => {
      console.log('setp3 finished')
      return true
    },
  })

  return steps
}

const TestGuide = () => {
  return (
    <Guide
      onFinish={async () => {
        console.log('guide finish')
        return true
      }}
      steps={makeSteps()}
    />
  )
}

export default TestGuide
