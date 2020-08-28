import { ethers } from 'ethers'

import { write } from '../../actions'
import {
    AppDispatch,
    TransactionActionTypes,
    GetActionTypes,
    TxData
} from '../../types'

import { Transaction } from '../../../config'

export const initialise = () => {
  return async (dispatch: AppDispatch, getState: Function) => {
    await dispatch(write({data: []})(GetActionTypes.GET_INIT))
  }
}

export const getData = (url: string) => {
  return async (dispatch: AppDispatch, getState: Function) => {

       fetch(url, {
         method: 'GET',
       })
       .then(response => {
         if (!response.ok) {
            const statusText = response.statusText
             return response.json()
             .then(data => {
                 throw new Error(statusText)
             })
         }
         return response.json()
       })
       .then(data => {
           dispatch(write({data: Array.of(data)})(GetActionTypes.GET_SUCCESS))
       })
      .catch(error => {
          const d = new Date(Date.now())
          const dateText = d.toString()
          const txData: TxData = {
              key: '-1',
              summary: `${Transaction.errorGettingData}: ${error.message} at ${dateText}`,
              time: dateText
           }
           //console.log(`${Transaction.errorGettingData}: ${error.message} at ${dateText}`)
           dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_FAILURE))
      })
  }
}
