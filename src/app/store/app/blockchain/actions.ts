import React from 'react'

import Web3 from 'web3'
import { ethers } from 'ethers'
//import { JsonRpcProvider } from 'ethers/providers/json-rpc-provider'

import { Contracts, Config } from '../../../config/contracts'
import { Entities } from '../../../config/typechain/Entities'
import { Artefacts } from '../../../config/typechain/Artefacts'

import { write } from '../../actions'
import {
    AppDispatch,
    ContractProps,
    ChainInfoProps,
    ChainContractActionTypes,
    ChainInfoActionTypes,
    TransactionActionTypes,
    TxData,
    Author,
    CopyrightHolder,
    Publisher,
    EntityTypes,
    CreativeWorks
} from '../../types'

import { Transaction } from '../../../config'

const txDispatch = (tx: any) => {
    return async (dispatch: AppDispatch, getState: Function) => {

        const state = getState()
        const provider = state.chainInfo.data.Provider

        let d = new Date(Date.now())
        let txData: TxData = {
            key:  tx.hash,
            summary: `${Transaction.pending}`,
            time: d.toString()
        }
        dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_PENDING))

        await provider.waitForTransaction(tx.hash)
        d = new Date(Date.now())
        txData = {
            key:  tx.hash,
            summary: `${Transaction.success}`,
            time: d.toString()
        }
        dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_SUCCESS))
    }
}

export const init = () => {
    return async (dispatch: AppDispatch) => {

        //let blockchainProvider: JsonRpcProvider
        let ethereum = (window as any).ethereum

        if (ethereum) {

          //let web3 = (window as any).web3
          await ethereum.enable()
          let web3 = new Web3(ethereum)
          let blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider as any)
          const signer = await blockchainProvider.getSigner()
          const account = await signer.getAddress()
          const chainObj = await blockchainProvider.getNetwork()
          let ENSAddress = ""
          if (typeof chainObj.ensAddress != 'undefined') {
            ENSAddress = chainObj.ensAddress
          }
          const infoData: ChainInfoProps = {
            data: {
              Name: chainObj.name,
              ChainId: chainObj.chainId,
              ENS: ENSAddress,
              Provider: blockchainProvider,
              Account: account
            }
          }

         //console.log(infoData)

          await dispatch(write({data: infoData})(ChainInfoActionTypes.ADD_DATA))

          const contractData: ContractProps = {
            data: {
              contracts: {
                entities: new ethers.Contract(
                    Contracts.entitiesAddress,
                    Contracts.entitiesABI,
                    signer
                ) as Entities,
                artefacts: new ethers.Contract(
                    Contracts.artefactsAddress,
                    Contracts.artefactsABI,
                    signer
                ) as Artefacts
              }
            }
          }

          dispatch(write({data: contractData.data})(ChainContractActionTypes.ADD_CONTRACTS))
      }
  }
}

export const addAuthor = (authorProps: Author) => {
  return async (dispatch: AppDispatch, getState: Function) => {

     const state = getState()
     const entitiesContract = state.chainContracts.data.contracts.entities

     try {

        const tx = await entitiesContract.addEntity(authorProps, EntityTypes.Author)
        dispatch(txDispatch(tx))

      } catch (error) {

        const d = new Date(Date.now())
        let txData: TxData  = {
            key:  "-1",
            summary: `${Transaction.failure}`,
            time: d.toString()
        }
        dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_FAILURE))
    }

  }
}

export const addCopyrightHolder = (copyrightProps: CopyrightHolder) => {
  return async (dispatch: AppDispatch, getState: Function) => {

     const state = getState()
     const entitiesContract = state.chainContracts.data.contracts.entities

     try {

        const tx = await entitiesContract.addEntity(copyrightProps, EntityTypes.CopyrightHolder)
        dispatch(txDispatch(tx))

      } catch (error) {

        const d = new Date(Date.now())
        let txData: TxData  = {
            key:  "-1",
            summary: `${Transaction.failure}`,
            time: d.toString()
        }
        dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_FAILURE))
    }

  }
}

export const addPublisher = (publisherProps: Publisher) => {
  return async (dispatch: AppDispatch, getState: Function) => {

     const state = getState()
     const entitiesContract = state.chainContracts.data.contracts.entities

     try {

        const tx = await entitiesContract.addEntity(publisherProps, EntityTypes.Publisher)
        dispatch(txDispatch(tx))

      } catch (error) {

        const d = new Date(Date.now())
        let txData: TxData  = {
            key:  "-1",
            summary: `${Transaction.failure}`,
            time: d.toString()
        }
        dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_FAILURE))
    }

  }
}

export const addFile = (props: CreativeWorks) => {
  return async (dispatch: AppDispatch, getState: Function) => {

      const state = getState()
      const artefactsContract = state.chainContracts.data.contracts.artefacts

      //console.log("file props: ", props)

      try {

         const tx = await artefactsContract.addWork(props)
         dispatch(txDispatch(tx))

       } catch (error) {

         const d = new Date(Date.now())
         let txData: TxData  = {
             key:  "-1",
             summary: `${Transaction.failure}`,
             time: d.toString()
         }
         dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_FAILURE))
     }

  }
}

export const addFileAuthor = (fileId: string, authorId: string) => {
  return async (dispatch: AppDispatch, getState: Function) => {

      const state = getState()
      const artefactsContract = state.chainContracts.data.contracts.artefacts

      //console.log("file props: ", props)

      try {

         const tx = await artefactsContract.addWorkAuthor(fileId, authorId)
         dispatch(txDispatch(tx))

       } catch (error) {

         const d = new Date(Date.now())
         let txData: TxData  = {
             key:  "-1",
             summary: `${Transaction.failure}`,
             time: d.toString()
         }
         dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_FAILURE))
     }

  }
}

export const addFileCopyrightHolder = (fileId: string, copyrightHolderId: string) => {
  return async (dispatch: AppDispatch, getState: Function) => {

      const state = getState()
      const artefactsContract = state.chainContracts.data.contracts.artefacts

      //console.log("file props: ", props)

      try {

         const tx = await artefactsContract.addWorkCopyrightHolder(fileId, copyrightHolderId)
         dispatch(txDispatch(tx))

       } catch (error) {

         const d = new Date(Date.now())
         let txData: TxData  = {
             key:  "-1",
             summary: `${Transaction.failure}`,
             time: d.toString()
         }
         dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_FAILURE))
     }

  }
}

export const addFilePublisher = (fileId: string, publisherId: string) => {
  return async (dispatch: AppDispatch, getState: Function) => {

      const state = getState()
      const artefactsContract = state.chainContracts.data.contracts.artefacts

      //console.log("file props: ", props)

      try {

         const tx = await artefactsContract.addWorkPublisher(fileId, publisherId)
         dispatch(txDispatch(tx))

       } catch (error) {

         const d = new Date(Date.now())
         let txData: TxData  = {
             key:  "-1",
             summary: `${Transaction.failure}`,
             time: d.toString()
         }
         dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_FAILURE))
     }

  }
}
