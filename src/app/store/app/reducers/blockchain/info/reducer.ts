import { ActionProps, ChainInfoActionTypes, ChainInfoProps } from '../../../../types'

const initialInfoState: ChainInfoProps = {
  data: {
    Name: '',
    ChainId: 0,
    ENS: '',
    Provider: {},
    Account: ''
  }
}

export const reducer = (state: ChainInfoProps = initialInfoState, action: ActionProps): ChainInfoProps => {
  //console.log('blockchain info: ', action.type, action.payload)
  if ( action.type == ChainInfoActionTypes.ADD_DATA ) {
    //console.log('Chain info: ', action.payload)
    return Object.assign({}, state, action.payload.data)
  } else {
    return state
  }
}
