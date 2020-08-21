import { ActionProps, ChainContractActionTypes, ContractProps } from '../../../../types'

 const initialContractState: ContractProps = {
    data: {
      contracts: {
        entities: {},
        artefacts: {}
      }
    }
  }

export const reducer = (state: ContractProps = initialContractState, action: ActionProps): ContractProps => {
  //console.log('Account info: ', action.type, action.payload)
  if ( action.type == ChainContractActionTypes.ADD_CONTRACTS ) {
    return Object.assign({}, state, action.payload)
  } else {
    return state
  }
}
