import { Action } from 'redux'
import { ThunkDispatch } from 'redux-thunk'
import { NavLink } from 'react-router-dom'

// Store stuff
export interface ApplicationState {
  chainInfo: ChainInfoProps,
  chainContracts: ContractProps
  info: InfoPageProps
  tx: TransactionProps
  data: GetProps
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

// Get stuff
export type DataProps = Entities

export interface GetProps extends PayloadProps {
    data: Array<DataProps>
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

//Creators
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
export type Entities = Author | CopyrightHolder  | Publisher

export interface EntityProps {
    entities: Array<Entities>
}

// the works they create
export interface CreativeWorks {
    workType: number
    license: number
    id: string
    authorId: string
    dateCreated: string
    dateModified: string
    url: string
    name: string
    description: string
}

//export type CreativeWorks = Works & Id

export interface CreativeWorksProps {
    fileInfo: Array<CreativeWorks>
}

// Reference stuff - unique Ids
export interface Id {
     id: string
}

// Action types

// Get action InfoTypes// Actiontypes stuff
export const enum GetActionTypes {
  GET_INIT = '@@GetActionTypes/GET_INIT',
  GET_SUCCESS = '@@GetActionTypes/GET_SUCCESS',
  GET_FAILURE = '@@GetActionTypes/GET_FAILURE'
}

// put action types
export const enum TransactionActionTypes {
  TRANSACTION_INIT = '@@TransactionActionTypes/TRANSACTION_INIT',
  TRANSACTION_PENDING = '@@TransactionActionTypes/TRANSACTION_PENDING',
  TRANSACTION_SUCCESS = '@@TransactionActionTypes/TRANSACTION_SUCCESS',
  TRANSACTION_FAILURE = '@@TransactionActionTypes/TRANSACTION_FAILURE'
}

// blockchain contracts
export const enum ChainContractActionTypes {
  ADD_CONTRACTS = '@@ChainContractAction/ADD_CONTRACTS'
}

// blockchain info
export const enum ChainInfoActionTypes {
  ADD_DATA = '@@ChainInfoAction/ADD_DATA'
}
