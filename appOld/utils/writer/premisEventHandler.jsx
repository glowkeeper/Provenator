import { keccak256 } from 'js-sha3'

class PremisEventHandler {

  constructor () {

    this.event = {
      id: undefined,
      type: undefined,
      date: undefined
    }
  }

  reset () {
    this.event.id = undefined
    this.event.name = undefined
    this.event.date = undefined
  }

  checkSet () {
    if ((this.event.id === undefined) || (this.event.name === undefined) || (this.event.date === undefined)) {
      return false
    } else {
      return true
    }
  }

  setId (_account) {
    const now = Date.now().toString()
    const keyString = _account + now
    const key = keccak256(keyString)
    this.event.id = key
    this.event.date = now
  }

  setType (_value) {
    this.event.type = _value
  }

  getId () {
    return this.event.id;
  }

  getType () {
    return this.event.type
  }

  getDate () {
    return this.event.date
  }
}

export default PremisEventHandler
