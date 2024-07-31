import { configureStore } from '@reduxjs/toolkit'
import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux'
import { hostSlice } from './ModuleHosts'

const store = configureStore({
  reducer: {
    host: hostSlice.reducer,
  },
})

type AppDispatch = typeof store.dispatch
type RootState = ReturnType<typeof store.getState>

export default store
export const useAppDispatch: () => AppDispatch = useDispatch
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector
