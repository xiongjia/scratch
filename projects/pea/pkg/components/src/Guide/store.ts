import { combineReducers, configureStore } from '@reduxjs/toolkit'

export interface Todo {
  id: number
  description: string
  checked: boolean
}

export interface TodoState {
  data: Todo[]
}

export const CREATE_TODO_REQUEST = '@todo/CREATE_TODO_REQUEST'
export const UPDATE_TODO_REQUEST = '@todo/UPDATE_TODO_REQUEST'
export const DELETE_TODO_REQUEST = '@todo/DELETE_TODO_REQUEST'

interface CreateTodoRequest {
  type: typeof CREATE_TODO_REQUEST
  payload: { todo: Todo }
}

interface UpdateTodoRequest {
  type: typeof UPDATE_TODO_REQUEST
  payload: { todo: Todo }
}

interface DeleteTodoRequest {
  type: typeof DELETE_TODO_REQUEST
  payload: { todo: Todo }
}

export type TodoActionsTypes =
  | CreateTodoRequest
  | UpdateTodoRequest
  | DeleteTodoRequest

export const actions = {
  createTodo: (todo: Todo): TodoActionsTypes => {
    return {
      type: CREATE_TODO_REQUEST,
      payload: { todo },
    }
  },
  updateTodo: (todo: Todo): TodoActionsTypes => {
    return {
      type: UPDATE_TODO_REQUEST,
      payload: { todo },
    }
  },

  deleteTodo: (todo: Todo): TodoActionsTypes => {
    return {
      type: DELETE_TODO_REQUEST,
      payload: { todo },
    }
  },
}

const initialState: TodoState = {
  data: [],
}

export const todoReducer = (
  state = initialState,
  action: TodoActionsTypes,
): TodoState => {
  switch (action.type) {
    case CREATE_TODO_REQUEST:
      return {
        data: [...state.data, action.payload.todo],
      }
    case UPDATE_TODO_REQUEST: {
      const data = state.data.map((todo) =>
        todo.id === action.payload.todo.id ? action.payload.todo : todo,
      )

      return { data }
    }
    case DELETE_TODO_REQUEST: {
      const data = state.data.filter(
        (todo) => todo.id !== action.payload.todo.id,
      )

      return { data }
    }
    default:
      return state
  }
}

export const rootReducer = combineReducers({
  todo: todoReducer,
})

export type RootState = ReturnType<typeof rootReducer>

export const store = configureStore({
  reducer: rootReducer,
})
