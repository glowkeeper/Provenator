import { Action } from 'redux'
import { ThunkDispatch } from 'redux-thunk'
import { NavLink } from 'react-router-dom'

// Store stuff
export interface ApplicationState {
  chainInfo: ChainInfoProps,
  chainContracts: ContractProps
  info: InfoPageProps
  tx: TransactionProps
}

export interface PayloadProps {
  data: object
}

export interface ActionProps extends Action {
  type: string
  payload: PayloadProps
}

export type AppDispatch = ThunkDispatch<ApplicationState, any, ActionProps>

// Blockchain stuff
export interface ChainInfoProps extends PayloadProps {
  data: {
    Name: string
    ChainId: number
    ENS: string
    Provider: object,
    Account: string
  }
}

export interface ContractProps extends PayloadProps {
  data: {
    contracts: {
      entities: object
      artefacts: object
    }
  }
}

// Info (about etc.) stuff
export const enum InfoTypes {
  HOME = "home",
  ABOUT = "about",
  HELP = "help",
  FAQ = "faq",
  CONTACT = "contact"
}

export interface InfoPageProps extends PayloadProps {
  data: InfoData
}

export interface InfoProps {
  title: string
  data: string
}

export interface InfoData {
  home: InfoProps
  about: InfoProps
  help: InfoProps
  faq: InfoProps
  contact: InfoProps
}

//Tx stuff
export interface TxData {
  key: string
  summary: string
  time: string
}

export interface TransactionProps extends PayloadProps {
  data: TxData
}

// Creative Works stuff
export enum EntityTypes {
    Author = 1,
    CopyrightHolder,
    Publisher
}

export interface Entity {
    name: string
    email: string
    url: string
}

export type Author = Entity & Id
export type CopyrightHolder = Entity & Id
export type Publisher = Entity & Id

export interface Works {
    workType: number
    license: number
    dateCreated: string
    dateModified: string
    url: string
    name: string
    description: string
    author: Author
    copyrightHolder: CopyrightHolder
    publisher: Publisher
}

// Reference stuff
export interface Id {
     id: string
}

export type CreativeWorks = Works & Id

export interface CreativeWorksProps {
    fileInfo: Array<CreativeWorks>
}

// Action types
export const enum TransactionActionTypes {
  TRANSACTION_INIT = '@@TransactionActionTypes/TRANSACTION_INIT',
  TRANSACTION_PENDING = '@@TransactionActionTypes/TRANSACTION_PENDING',
  TRANSACTION_SUCCESS = '@@TransactionActionTypes/TRANSACTION_SUCCESS',
  TRANSACTION_FAILURE = '@@TransactionActionTypes/TRANSACTION_FAILURE'
}


export const enum ChainContractActionTypes {
  ADD_CONTRACTS = '@@ChainContractAction/ADD_CONTRACTS'
}

export const enum ChainInfoActionTypes {
  ADD_DATA = '@@ChainInfoAction/ADD_DATA'
}
