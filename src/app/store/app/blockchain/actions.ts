import React from 'react'

import Web3 from 'web3'
import { ethers } from 'ethers'
import { JsonRpcProvider } from 'ethers/providers/json-rpc-provider'

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
    EntityTypes,
    CreativeWorks
} from '../../types'

import { Transaction } from '../../../config'

export const init = () => {
    return async (dispatch: AppDispatch) => {

        let blockchainProvider: JsonRpcProvider
        let ethereum = (window as any).ethereum
        let web3 = (window as any).web3

        if (ethereum) {

          await ethereum.enable()
          web3 = new Web3(ethereum)
          blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
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

export const addAuthor = (props: Author) => {
  return async (dispatch: AppDispatch, getState: Function) => {

     const state = getState()
     const entitiesContract = state.chainContracts.data.contracts.entities
     const provider = state.chainInfo.data.Provider

     const entityType = EntityTypes.Author

     let d = new Date(Date.now())
     let txData: TxData  = {
        key: "",
        summary: "",
        time:  ""
      }

     try {

        const tx = await entitiesContract.addEntity(props, entityType)
        txData = {
            key:  tx.hash,
            summary: `${Transaction.pending}`,
            time: d.toString()
        }
        dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_PENDING))

        const receipt = await provider.waitForTransaction(tx.hash)
        d = new Date(Date.now())
        txData = {
            key:  tx.hash,
            summary: `${Transaction.success}`,
            time: d.toString()
        }
        dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_SUCCESS))

      } catch (error) {

        d = new Date(Date.now())
        txData = {
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

      console.log(props, state.chainContracts)

      const artefactsContract = state.chainContracts.data.contracts.artefacts
      const provider = state.chainInfo.data.provider

     /*const createInfoWriter: ExtraInfoWriterProps = {
          fileHash: props.fileHash,
          fileName: props.fileName,
          fileUrl: props.fileUrl,
          description: props.description
      }
      const jobRef = setBytes32(props.ref)
      const workRef = setBytes32(props.workRef)
      const infoRef = setBytes32(props.infoRef)

      let txData = {
          transaction:  -1,
          summary: '',
          info: {
              time: ''
          }
      }

     let d = new Date(Date.now())
     try {

        //console.log("dispatching write: ", jobRef, workRef, infoRef, createInfoWriter)
        const tx = await jobsContract.addInfo(jobRef, workRef, infoRef, createInfoWriter)
        txData = {
            transaction:  tx.hash,
            summary: `${Transaction.pending}`,
            info: {
                time: d.toString()
            }
        }
        dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_PENDING))

        const receipt = await provider.waitForTransaction(tx.hash)
        d = new Date(Date.now())
        txData = {
            transaction:  tx.hash,
            summary: `${Transaction.success}`,
            info: {
                time: d.toString()
            }
        }
        dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_SUCCESS))

      } catch (error) {

        d = new Date(Date.now())
        txData = {
            transaction:  -1,
            summary: `${Transaction.failure}`,
            info: {
                time: d.toString()
            }
        }
        dispatch(write({data: txData})(TransactionActionTypes.TRANSACTION_FAILURE))
    }*/

  }
}
