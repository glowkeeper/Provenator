import { keccak256 } from 'js-sha3'
import {PremisAgentStrings} from '../helpers/outputStrings'

class PremisAgentHandler {

  static typeOptions = [
    PremisAgentStrings.agentIndividual,
    PremisAgentStrings.agentCompany
  ]

  constructor () {

    this.agent = {
      id: undefined,
      name: undefined,
      type: undefined
    }
  }

  reset () {
    this.agent.id = undefined
    this.agent.name = undefined
    this.agent.type = undefined
  }

  checkSet () {
    if ((this.agent.id === undefined) || (this.agent.name === undefined) || (this.agent.type === undefined)) {
      return false
    } else {
      return true
    }
  }

  setId (_account) {
    const keyString = _account +
                      this.agent.name +
                      this.agent.type
    const key = keccak256(keyString)
    this.agent.id = key;
  }

  setName (_value) {
    this.agent.name = _value
  }

  setType (_value) {
    this.agent.type = _value
  }

  getId () {
    return this.agent.id;
  }

  getName () {
    return this.agent.name
  }

  getType () {
    return this.agent.type
  }
}

export default PremisAgentHandler
