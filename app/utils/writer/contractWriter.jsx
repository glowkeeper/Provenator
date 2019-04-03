import Web3Handler from '../web3Handler'
import ContractHandler from '../contractHandler'

import PremisObjectHandler from './premisObjectHandler'
import PremisRightsHandler from './premisRightsHandler'
import PremisEventHandler from './premisEventHandler'
import PremisAgentHandler from './premisAgentHandler'

import {PremisEventStrings} from '../../helpers/strings'

class ContractWriter {

  constructor (_web3Handler, _contractHandler) {
    //console.log(_web3)
    this.web3Handler = _web3Handler
    const contractHandler = _contractHandler
    this.premisObject = contractHandler.getPremisObject()
    this.premisEvent = contractHandler.getPremisEvent()
    this.premisAgent = contractHandler.getPremisAgent()
    this.premisRights = contractHandler.getPremisRights()

    this.premisObjectData = undefined
    this.premisEventData = undefined
    this.premisAgentData = undefined
    this.premisRightsData = undefined

    this.numTransactions = 17
    this.numWrites = 0
  }

  getNumTransactions () {
    return this.numTransactions
  }

  getNumWrites () {
    return this.numWrites
  }

  _checkData () {
    if ((this.premisObjectData === undefined) ||
        (this.premisEventData === undefined) ||
        (this.premisAgentData === undefined) ||
        (this.premisRightsData === undefined)) {
      return false
    } else {
      return true
    }
  }

  _resetNumWrites () {
    this.numWrites = 0
  }

  _writeCounter (_self, _result) {
    //console.log(_result)
    _self.numWrites += 1
  }

  // Premis Object Stuff

  _setPropType(_self, _exists) {
    // console.log('Checking new property type exists with index ' + _propTypeIndex)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const propType = _self.premisObjectData.getPropertyType()
      const account = _self.web3Handler.getAccount()
      const params = [propType, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.methods.setPropertyType, params, _self._writeCounter, false)
    }
  }

  _setObject(_self, _exists) {
    if (_exists) {
      _self.numWrites += 1
    } else {
      const hash = _self.premisObjectData.getHash()
      const category = _self.premisObjectData.getCategory()
      const format = _self.premisObjectData.getFormat()
      const account = _self.web3Handler.getAccount()
      const params = [hash, category, format, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.methods.setObject, params, _self._writeCounter, false)
    }
  }

  _setObjectProperties(_self, _exists) {
    // console.log('Checking new property type exists with index ' + _propTypeIndex)
    if (_exists ) {
      _self.numWrites += 1
    } else {
      const hash = _self.premisObjectData.getHash()
      const propType = _self.premisObjectData.getPropertyType()
      const propValue = _self.premisObjectData.getPropertyValue()
      const account = _self.web3Handler.getAccount()
      const params = [hash, propType, propValue, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.methods.setProperties, params, _self._writeCounter, false)
    }
  }

  _setObjectEvent(_self, _exists) {
    //console.log('Setting Object event ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const hash = _self.premisObjectData.getHash()
      const eventId = _self.premisEventData.getId()
      const account = _self.web3Handler.getAccount()
      const params = [hash, eventId, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.methods.setEvent, params, _self._writeCounter, false)
    }
  }

  _setObjectAgent(_self, _exists) {
    //console.log('Setting Object agent ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const hash = _self.premisObjectData.getHash()
      const agentId = _self.premisAgentData.getId()
      const account = _self.web3Handler.getAccount()
      const params = [hash, agentId, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.methods.setAgent, params, _self._writeCounter, false)
    }
  }

  _setObjectRights(_self, _exists) {
    //console.log('Object Rights Index ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const hash = _self.premisObjectData.getHash()
      const rightsId = _self.premisRightsData.getId()
      const account = _self.web3Handler.getAccount()
      const params = [hash, rightsId, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.methods.setRights, params, _self._writeCounter, false)
    }
  }

  _objectSubmit () {
    const hash = this.premisObjectData.getHash()
    const propType = this.premisObjectData.getPropertyType()
    const propValue = this.premisObjectData.getPropertyValue()
    const agentId = this.premisAgentData.getId()
    const eventId = this.premisEventData.getId()
    const rightsId = this.premisRightsData.getId()
    const account = this.web3Handler.getAccount()
    let params = [propType, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisObject.methods.getPropertyTypeExists, params, this._setPropType, true)
    params = [hash, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisObject.methods.getObjectExists, params, this._setObject, true)
    params = [hash, propType, propValue, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisObject.methods.getObjectPropertiesExist, params, this._setObjectProperties, true)
    params = [hash, eventId, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisObject.methods.getObjectEventExists, params, this._setObjectEvent, true)
    params = [hash, agentId, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisObject.methods.getObjectAgentExists, params, this._setObjectAgent, true)
    params = [hash, rightsId, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisObject.methods.getObjectRightsExists, params, this._setObjectRights, true)
  }

  // Premis Event Stuff

  _setEventType(_self, _exists) {
    //console.log('Setting event type ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const account = _self.web3Handler.getAccount()
      const eventTypeName = _self.premisEventData.getType()
      const params = [eventTypeName, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisEvent.methods.setEventType, params, _self._writeCounter, false)
    }
  }

  _setEvent(_self, _exists) {
    //console.log('Setting Event ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const eventId = _self.premisEventData.getId()
      const eventTypeName = _self.premisEventData.getType()
      const date = _self.premisEventData.getDate()
      const hash = _self.premisObjectData.getHash()
      const agentId = _self.premisAgentData.getId()
      const account = _self.web3Handler.getAccount()
      const params = [eventId, eventTypeName, date, hash, agentId, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisEvent.methods.setEvent, params, _self._writeCounter, false)
    }
  }

  _eventSubmit () {
    this.premisEventData.setType(PremisEventStrings.objectRecord)
    //console.log(PremisEventStrings.objectRecord)
    const eventId = this.premisEventData.getId()
    const account = this.web3Handler.getAccount()
    let params = [PremisEventStrings.objectRecord, {from: account}]
    //console.log(PremisEventStrings.objectRecord)
    this.web3Handler.callParamHandler(this, this.premisEvent.methods.getEventTypeExists, params, this._setEventType, true)
    params = [eventId, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisEvent.methods.getEventExists, params, this._setEvent, true)
  }

  // Premis Agent Stuff

  _setAgent(_self, _exists) {
    //console.log('Setting Agent ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const agentId = _self.premisAgentData.getId()
      const name = _self.premisAgentData.getName()
      const agentType = _self.premisAgentData.getType()
      const account = _self.web3Handler.getAccount()
      const params = [agentId, name, agentType, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.methods.setAgent, params, _self._writeCounter, false)
    }
  }

  _setAgentObject(_self, _exists) {
    //console.log('Setting Agent Object ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const agentId = _self.premisAgentData.getId()
      const hash = _self.premisObjectData.getHash()
      const account = _self.web3Handler.getAccount()
      const params = [agentId, hash, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.methods.setObject, params, _self._writeCounter, false)
    }
  }

  _setAgentEvent(_self, _exists) {
    //console.log('Setting Agent event ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const agentId = _self.premisAgentData.getId()
      const eventId = _self.premisEventData.getId()
      const account = _self.web3Handler.getAccount()
      const params = [agentId, eventId, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.methods.setEvent, params, _self._writeCounter, false)
    }
  }

  _setAgentRights(_self, _exists) {
    //console.log('Setting Agent rights ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const agentId = _self.premisAgentData.getId()
      const rightsId = _self.premisRightsData.getId()
      const account = _self.web3Handler.getAccount()
      const params = [agentId, rightsId, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.methods.setRights, params, _self._writeCounter, false)
    }
  }

  _agentSubmit () {
    const agentId = this.premisAgentData.getId()
    const hash = this.premisObjectData.getHash()
    const eventId = this.premisEventData.getId()
    const rightsId = this.premisRightsData.getId()
    const account = this.web3Handler.getAccount()
    let params = [agentId, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisAgent.methods.getAgentExists, params, this._setAgent, true)
    params = [agentId, hash, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisAgent.methods.getAgentObjectExists, params, this._setAgentObject, true)
    params = [agentId, eventId, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisAgent.methods.getAgentEventExists, params, this._setAgentEvent, true)
    params = [agentId, rightsId, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisAgent.methods.getAgentRightsExists, params, this._setAgentRights, true)
  }

  // Premis Rights Stuff

  _setRights(_self, _exists) {
    //console.log('Setting rights ' + _exists)
    if ( _exists ) {
      _self.numWrites += 5
    } else {
      // since the rights index key is made up of every field, we need to update everything here
      const rightsId = _self.premisRightsData.getId()
      const basis = _self.premisRightsData.getBasis()
      const status = _self.premisRightsData.getStatus()
      const code = _self.premisRightsData.getCountryCode()
      const date = _self.premisRightsData.getDeterminationdate()
      const note = _self.premisRightsData.getNote()
      const act = _self.premisRightsData.getAct()
      const restriction = _self.premisRightsData.getRestriction()
      const hash = _self.premisObjectData.getHash()
      const agentId = _self.premisAgentData.getId()
      const account = _self.web3Handler.getAccount()
      let params = [rightsId, hash, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.methods.setObject, params, _self._writeCounter, false)
      params = [rightsId, basis, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.methods.setBasis, params, _self._writeCounter, false)
      params = [rightsId, status, code, date, note, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.methods.setCopyrightInfo, params, _self._writeCounter, false)
      params = [rightsId, act, restriction, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.methods.setRightsGranted, params, _self._writeCounter, false)
      params = [rightsId, agentId, {from: account, gas: 300000}]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.methods.setAgent, params, _self._writeCounter, false)
    }
  }

  _rightsSubmit () {
    const rightsId = this.premisRightsData.getId()
    const account = this.web3Handler.getAccount()
    let params = [rightsId, {from: account}]
    this.web3Handler.callParamHandler(this, this.premisRights.methods.getRightsExists, params, this._setRights, true)
  }

  setData (_premisObjectHandler, _premisEventHandler, _premisAgentHandler, _premisRightsHandler) {
    this.premisObjectData = _premisObjectHandler
    this.premisEventData = _premisEventHandler
    this.premisAgentData = _premisAgentHandler
    this.premisRightsData = _premisRightsHandler
  }

  resetData () {
    this.premisObjectData = undefined
    this.premisEventData = undefined
    this.premisAgentData = undefined
    this.premisRightsData = undefined
  }

  dispatcher () {
    if ( this._checkData() ) {
      this._resetNumWrites()
      this._objectSubmit()
      this._eventSubmit()
      this._agentSubmit()
      this._rightsSubmit()
    }
  }
}

export default ContractWriter
