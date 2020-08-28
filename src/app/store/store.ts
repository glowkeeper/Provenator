import { combineReducers, Reducer, Store, createStore, applyMiddleware } from 'redux'
import ReduxThunk from 'redux-thunk'

import { ApplicationState, ActionProps } from './types'

import { reducer as chainReducer } from './app/reducers/blockchain/info/reducer'
import { reducer as contractReducer } from './app/reducers/blockchain/contracts/reducer'
import { reducer as infoReducer } from './app/reducers/info/reducer'
import { reducer as txReducer } from './app/reducers/tx/reducer'
import { reducer as dataReducer } from './app/reducers/get/data/reducer'

export const rootReducer: Reducer<ApplicationState, ActionProps> = combineReducers<ApplicationState, ActionProps>({
  chainInfo: chainReducer,
  chainContracts: contractReducer,
  info: infoReducer,
  tx: txReducer,
  data: dataReducer
})

export function configureStore(
  initialState: ApplicationState
): Store<ApplicationState, ActionProps> {

  // create the redux-saga middleware
  const store = createStore(
    rootReducer,
    initialState,
    applyMiddleware(ReduxThunk)
  )

  return store
}
