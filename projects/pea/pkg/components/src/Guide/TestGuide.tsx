import React from 'react'
import Guide, { GuideStep } from './Guide'
import GuideContextProvider from './GuideContext'
import StepHostInput from './StepHostInput'
import StepHostsSelector from './StepHostsSelector'

const makeSteps = (): GuideStep[] => {
  const steps: GuideStep[] = []
  steps.push(StepHostInput)
  steps.push(StepHostsSelector)

  steps.push({
    name: 'step3 - name',
    title: 'step3 - title',
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
    <GuideContextProvider>
      <Guide
        onFinish={async () => {
          console.log('guide finish')
          return true
        }}
        steps={makeSteps()}
      />
    </GuideContextProvider>
  )
}

export default TestGuide
