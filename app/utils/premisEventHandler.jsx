import { keccak256 } from 'js-sha3';

class PremisEventHandler {

  constructor () {

    this.event = {
      id: undefined,
      name: undefined,
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
    this.event.id = key;
  }

  setName (_value) {
    this.event.name = _value
  }

  setDate (_value) {
    this.event.date = _value
  }

  getId () {
    return this.event.id;
  }

  getName () {
    return this.event.name
  }

  getDate () {
    return this.event.date
  }
}

export default PremisEventHandler
