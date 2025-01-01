import { useState } from 'react'

export declare type StoreBaseValue = string | number | boolean
export declare type StoreValue = StoreBaseValue | Store | StoreBaseValue[]
export interface Store {
  [name: string]: StoreValue
}

export interface UseFormConfig {
  defaultFormValues?: Store | (() => Promise<Store> | Store)
  submit?: (formValues: Store) => void
}

export interface UseStepsFormConfig extends UseFormConfig {
  defaultCurrent?: number
  total?: number
  isBackValidate?: boolean
}

export const useStepsForm = (config: UseStepsFormConfig) => {
  const { defaultCurrent } = config || ({} as UseStepsFormConfig)
  const [current, setCurrent] = useState(defaultCurrent)

  return {}
}
