import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit'

function delayTest(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

export interface Host {
  ipAddress: string
  description: string
  checked: boolean
}

export interface SshAccount {
  username: string
  password: string
}

export interface HostState {
  account: SshAccount
  hosts: Host[]
}

const initialState: HostState = {
  account: {
    username: '',
    password: '',
  },
  hosts: [],
}

export const loadHostList = createAsyncThunk(
  'common/host',
  async (_, { dispatch }) => {
    console.log('begin get host list')
    await delayTest(1000 * 10)

    console.log('dispatch')

    const hosts: Host[] = []
    hosts.push({
      ipAddress: 'ip1',
      description: 'h1',
      checked: false,
    })
    hosts.push({
      ipAddress: 'ip2',
      description: 'h2',
      checked: false,
    })
    dispatch(actions.setHosts(hosts))
  },
)

export const hostSlice = createSlice({
  name: 'host',
  initialState,
  reducers: {
    setSshAccount: (state, action: PayloadAction<SshAccount>) => {
      state.account = action.payload
    },
    setHosts: (state, action: PayloadAction<Host[]>) => {
      state.hosts = action.payload
    },
  },
})

export const actions = { ...hostSlice.actions, loadHostList }
