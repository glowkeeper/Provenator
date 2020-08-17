import React from 'react'

import Web3 from 'web3'
import { ethers } from 'ethers'
import { JsonRpcProvider } from 'ethers/providers/json-rpc-provider'

import { write } from '../../actions'
import { AppDispatch, ChainDataProps, ChainDataActionTypes, TransactionActionTypes, FileProps } from '../../types'

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
          const infoData: ChainDataProps = {
            data: {
              Name: chainObj.name,
              ChainId: chainObj.chainId,
              ENS: ENSAddress,
              provider: blockchainProvider,
              account: account
            }
          }

          await dispatch(write({data: infoData})(ChainDataActionTypes.ADD_DATA))
      }
  }
}

export const addContract = () => {
    return async (dispatch: AppDispatch) => {

        /*const contractData: ContractProps = {
          data: {
            contracts: {
              jobs: new ethers.Contract(
                  Contracts.jobsAddress,
                  Contracts.jobsABI,
                  signer
              ) as Jobs,
              spons: new ethers.Contract(
                  Contracts.sponsAddress,
                  Contracts.sponsABI,
                  signer
              ) as Spons,
              subbies: new ethers.Contract(
                  Contracts.subbiesAddress,
                  Contracts.subbiesABI,
                  signer
              ) as Subbies
            }
          }
        }

        dispatch(write({data: chainData.data})(ChainDataActionTypes.ADD_DATA))*/
    }
}

export const addFile = (props: FileProps) => {
  return async (dispatch: AppDispatch) => {


  }
}
