import { EmptyProps } from 'antd'
import React, { createContext, PropsWithChildren } from 'react'
import { Updater, useImmer } from 'use-immer'

export interface GuideContextType {
  sshUsername?: string
  sshPassword?: string
  beginAddr?: string
  endAddr?: string
}

const DefaultGuideContext: () => GuideContextType = () => ({
  sshUsername: 'root',
  sshPassword: 'root123',
  beginAddr: '1.1.1.1',
  endAddr: '1.1.1.10',
})

export const GuideContext = createContext<
  GuideContextType & { setState?: Updater<GuideContextType> }
>(DefaultGuideContext())

export default function GuideContextProvider({
  children,
}: PropsWithChildren<EmptyProps>) {
  const [state, setState] = useImmer(DefaultGuideContext())
  return (
    <GuideContext.Provider value={{ ...state, setState }}>
      {children}
    </GuideContext.Provider>
  )
}
