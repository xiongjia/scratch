import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components'
import { configureStore, createSlice } from '@reduxjs/toolkit'
import React, { useEffect, useRef } from 'react'
import { Provider, connect, useSelector, useStore } from 'react-redux'

type HostType = {
  ipAddr: string
  hostname: string
}

const counterSlice = createSlice({
  name: 'counter',
  initialState: {
    value: 0,
  },
  reducers: {
    incremented: (state) => {
      console.log('inc ')
      state.value += 1
    },
    decremented: (state) => {
      console.log('dec ')
      state.value -= 1
    },
  },
})

const store = configureStore({
  reducer: counterSlice.reducer,
})

export const { incremented, decremented } = counterSlice.actions
type RootState = ReturnType<typeof store.getState>

const HostColumns: ProColumns<HostType>[] = [
  {
    dataIndex: 'ipAddr',
    title: 'IP Address',
    editable: false,
  },
  {
    dataIndex: 'hostname',
    title: 'Hostname',
    width: 350,
  },
  {
    title: 'option',
    valueType: 'option',
    key: 'option',
    width: 100,
    fixed: 'right',
    render: (text, record, idx, action) => [
      <a
        key="remove"
        onClick={() => {
          console.log('remove ' + record.ipAddr)
          action?.startEditable(record.ipAddr)
        }}
      >
        REMOVE
      </a>,
    ],
  },
]

const testSrc: HostType[] = [
  {
    ipAddr: '1.1.1.1111111',
    hostname: 'host1',
  },
  {
    ipAddr: '1.1.1.2',
    hostname: 'host1',
  },
  {
    ipAddr: '1.1.1.3',
    hostname: 'host1',
  },
]

const HostButtons = connect(null, { incremented })(() => {
  const val1 = useSelector<RootState>((state) => state.value)
  const storelocal = useStore()
  const testFunc = () => {
    console.log('test buttons 1' + JSON.stringify(storelocal.getState()))
    incremented()
  }

  return (
    <>
      <p> v1 = {`${val1}`}</p>
      <button type="button" onClick={testFunc}>
        test2
      </button>
    </>
  )
})

const HostTable = () => {
  const actionRef = useRef<ActionType>()
  const inputRef = useRef<HTMLInputElement>(null)

  const testFunc2 = () => {
    console.log('func 2, before dispatch inc' + store.getState().value)
    store.dispatch(incremented())
    console.log('func 2, after dispatch inc' + store.getState().value)
  }

  const testFunc = () => {
    console.log('val = ' + inputRef.current?.value)
    if (inputRef.current) {
      inputRef.current.value = 'clear'
    }
  }

  useEffect(() => {
    store.subscribe(() => {
      console.log(
        'subcribe event , new state: ' + JSON.stringify(store.getState()),
      )
    })
  }, [])

  return (
    <Provider store={store}>
      <ProTable<HostType>
        columns={HostColumns}
        actionRef={actionRef}
        search={false}
        options={false}
        pagination={false}
        rowKey={(record) => record.ipAddr}
        scroll={{ x: 'max-content' }}
        request={async () => {
          return {
            data: testSrc,
            success: true,
            total: testSrc.length,
          }
        }}
        editable={{
          type: 'single',
          actionRender: () => {
            return []
          },
        }}
        rowSelection={{
          defaultSelectedRowKeys: ['ipAddr'],
          onChange: (selectedRowKeys) => {
            console.log('select keys = ' + selectedRowKeys)
          },
        }}
        tableAlertRender={false}
      />

      <button type="button" onClick={testFunc2}>
        test
      </button>
      <input ref={inputRef} onClick={testFunc} />

      <br />
      <HostButtons />
    </Provider>
  )
}

export default HostTable
