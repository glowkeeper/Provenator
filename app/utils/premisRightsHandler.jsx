import { keccak256 } from 'js-sha3'
import {PremisRightsStrings} from '../helpers/outputStrings'

class PremisRightsHandler {

  static actOptions = [
    PremisRightsStrings.actReplicate,
    PremisRightsStrings.actMigrate,
    PremisRightsStrings.actModify,
    PremisRightsStrings.actUse,
    PremisRightsStrings.actDisseminate,
    PremisRightsStrings.actDelete
  ]

  static restrictionsOptions = [
    PremisRightsStrings.restrictionsAllow,
    PremisRightsStrings.restrictionsDisallow,
    PremisRightsStrings.restrictionsConditional,
  ]

  static basisOptions = [
    PremisRightsStrings.basisCopyright,
    PremisRightsStrings.basisStatute,
    PremisRightsStrings.basisLicense,
    PremisRightsStrings.basisDonor,
    PremisRightsStrings.basisPolicy,
    PremisRightsStrings.basisOther,
    PremisRightsStrings.basisGrants
  ]

  constructor () {

    this.rights = {
      id: undefined,
      basis: undefined,
      status: undefined,
      countryCode: undefined,
      determinationDate: undefined,
      note: undefined,
      act: undefined,
      restriction: undefined
    }
  }

  reset () {
    this.rights.id = undefined
    this.rights.basis = undefined
    this.rights.status = undefined
    this.rights.countryCode = undefined
    this.rights.determinationDate = undefined
    this.rights.note = undefined
    this.rights.act = undefined
    this.rights.restriction = undefined
  }

  checkSet () {
    if ((this.rights.id === undefined) ||
        (this.rights.basis === undefined) ||
        (this.rights.status === undefined) ||
        (this.rights.countryCode === undefined) ||
        (this.rights.determinationDate === undefined) ||
        (this.rights.note === undefined) ||
        (this.rights.act === undefined) ||
        (this.rights.restriction === undefined)) {
      return false
    } else {
      return true
    }
  }

  setId () {
    const keyString = this.rights.id +
                      this.rights.basis +
                      this.rights.status +
                      this.rights.countryCode +
                      this.rights.determinationDate +
                      this.rights.act + this.rights.restriction
    const key = keccak256(keyString)
    this.rights.id = key;
  }

  setBasis (_basis) {
    this.rights.basis = _basis
  }

  setStatus (_status) {
    this.rights.status = _status
  }

  setCountryCode (_code) {
    this.rights.countryCode = _code
  }

  setDeterminationDate (_date) {
    this.rights.determinationDate = _date
  }

  setNote (_note) {
    this.rights.note = _note
  }

  setAct (_act) {
    this.rights.act = _act
  }

  setRestriction (_restriction) {
    this.rights.restriction = _restriction
  }

  getId () {
    return this.rights.id;
  }

  getBasis () {
    return this.rights.basis
  }

  getStatus () {
    return this.rights.status
  }

  getCountryCode () {
    return this.rights.countryCode
  }

  getDeterminationdate () {
    return this.rights.determinationDate
  }

  getNote () {
    return this.rights.note
  }

  getAct () {
    return this.rights.act
  }

  getRestriction () {
    return this.rights.restriction
  }

}

export default PremisRightsHandler
