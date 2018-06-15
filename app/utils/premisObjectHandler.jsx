import {PremisObjectStrings} from '../helpers/outputStrings'

class PremisObjectHandler {

  static categoryOptions = [
    PremisObjectStrings.categoryIntellectual,
    PremisObjectStrings.categoryRepresentation,
    PremisObjectStrings.categoryFile,
    PremisObjectStrings.categoryBitStream
  ]

  constructor () {

    this.object = {
      hash: undefined,
      category: undefined,
      format: undefined,
      propertyType: undefined,
      propertyValue: undefined
    }
  }

  reset () {
    this.object.hash = undefined
    this.object.category = undefined
    this.object.format = undefined
    this.object.propertyType = undefined
    this.object.propertyValue = undefined
  }

  checkSet () {
    if ((this.object.hash === undefined) ||
        (this.object.category === undefined) ||
        (this.object.format === undefined) ||
        (this.object.propertyType === undefined) ||
        (this.object.propertyValue === undefined)) {
      return false
    } else {
      return true
    }
  }

  setHash (_value) {
    this.object.hash = _value;
  }

  setCategory (_value) {
    this.object.category = _value
  }

  setFormat (_value) {
    this.object.format = _value
  }

  setPropertyType (_value) {
    this.object.propertyType = _value
  }

  setPropertyValue (_value) {
    this.object.propertyValue = _value
  }

  getHash () {
    return this.object.hash
  }

  getCategory () {
    return this.object.category
  }

  getFormat () {
    return this.object.format
  }

  getPropertyType () {
    return this.object.propertyType
  }

  getPropertyValue () {
    return this.object.propertyValue
  }

}

export default PremisObjectHandler
