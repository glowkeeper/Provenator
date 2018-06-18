import { keccak256 } from 'js-sha3';

class PremisEventHandler {

  constructor () {

    this.event = {
      id: undefined,
      date: undefined
    }
  }

  reset () {
    this.event.id = undefined
    this.event.date = undefined
  }

  checkSet () {
    if ((this.event.id === undefined) || (this.event.date === undefined)) {
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
    this.event.date = now
  }

  getId () {
    return this.event.id;
  }

  getDate () {
    return this.event.date
  }
}

export default PremisEventHandler
