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

    this.defaultTO = {gas: 300000}

    this.numTransactions = 17
    this.numWrites = 0
    this.batchWrites = false //can't get batch writes working :()
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

  _writeCounter (_self) {
    _self.numWrites += 1
    if ( (_self.numWrites == _self.numTransactions ) &&
         (_self.batchWrites) )  {
      _self.web3Handler.executeBatch()
    }
  }

  // Premis Object Stuff

  _setPropType(_self, _exists) {
    // console.log('Checking new property type exists with index ' + _propTypeIndex)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const propType = _self.premisObjectData.getPropertyType()
      const params = [propType, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.setPropertyType, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _setObject(_self, _exists) {
    if (_exists) {
      _self.numWrites += 1
    } else {
      const hash = _self.premisObjectData.getHash()
      const category = _self.premisObjectData.getCategory()
      const format = _self.premisObjectData.getFormat()
      const params = [hash, category, format, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.setObject, params, _self._writeCounter, _self.batchWrites)
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
      const params = [hash, propType, propValue, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.setProperties, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _setObjectEvent(_self, _exists) {
    //console.log('Setting Object event ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const hash = _self.premisObjectData.getHash()
      const eventId = _self.premisEventData.getId()
      const params = [hash, eventId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.setEvent, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _setObjectAgent(_self, _exists) {
    console.log('Setting Object agent ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const hash = _self.premisObjectData.getHash()
      const agentId = _self.premisAgentData.getId()
      const params = [hash, agentId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.setAgent, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _setObjectRights(_self, _exists) {
    console.log('Object Rights Index ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const hash = _self.premisObjectData.getHash()
      const rightsId = _self.premisRightsData.getId()
      const params = [hash, rightsId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.setRights, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _objectSubmit () {
    const hash = this.premisObjectData.getHash()
    const propType = this.premisObjectData.getPropertyType()
    const propValue = this.premisObjectData.getPropertyValue()
    const agentId = this.premisAgentData.getId()
    const eventId = this.premisEventData.getId()
    const rightsId = this.premisRightsData.getId()
    let params = [propType]
    this.web3Handler.callParamHandler(this, this.premisObject.getPropertyTypeExists, params, this._setPropType, false)
    params = [hash]
    this.web3Handler.callParamHandler(this, this.premisObject.getObjectExists, params, this._setObject, false)
    params = [hash, propType, propValue, this.defaultTO]
    this.web3Handler.callParamHandler(this, this.premisObject.getObjectPropertiesExist, params, this._setObjectProperties, false)
    params = [hash, eventId]
    this.web3Handler.callParamHandler(this, this.premisObject.getObjectEventExists, params, this._setObjectEvent, false)
    params = [hash, agentId]
    this.web3Handler.callParamHandler(this, this.premisObject.getObjectAgentExists, params, this._setObjectAgent, false)
    params = [hash, rightsId]
    console.log(rightsId)
    this.web3Handler.callParamHandler(this, this.premisObject.getObjectRightsExists, params, this._setObjectRights, false)
  }

  // Premis Event Stuff

  _setEventType(_self, _exists) {
    console.log('Setting event type ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const eventTypeName = _self.premisEventData.getType()
      const params = [eventTypeName, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisEvent.setEventType, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _setEvent(_self, _exists) {
    console.log('Setting Event ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const eventId = _self.premisEventData.getId()
      const eventTypeName = _self.premisEventData.getType()
      const date = _self.premisEventData.getDate()
      const hash = _self.premisObjectData.getHash()
      const agentId = _self.premisAgentData.getId()
      const params = [eventId, eventTypeName, date, hash, agentId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisEvent.setEvent, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _eventSubmit () {
    this.premisEventData.setType(PremisEventStrings.objectRecord)
    //console.log(PremisEventStrings.objectRecord)
    const eventId = this.premisEventData.getId()
    let params = [PremisEventStrings.objectRecord]
    //console.log(PremisEventStrings.objectRecord)
    this.web3Handler.callParamHandler(this, this.premisEvent.getEventTypeExists, params, this._setEventType, false)
    params = [eventId]
    this.web3Handler.callParamHandler(this, this.premisEvent.getEventExists, params, this._setEvent, false)
  }

  // Premis Agent Stuff

  _setAgent(_self, _exists) {
    console.log('Setting Agent ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const agentId = _self.premisAgentData.getId()
      const name = _self.premisAgentData.getName()
      const agentType = _self.premisAgentData.getType()
      const params = [agentId, name, agentType, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.setAgent, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _setAgentObject(_self, _exists) {
    console.log('Setting Agent Object ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const agentId = _self.premisAgentData.getId()
      const hash = _self.premisObjectData.getHash()
      const params = [agentId, hash, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.setObject, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _setAgentEvent(_self, _exists) {
    console.log('Setting Agent event ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const agentId = _self.premisAgentData.getId()
      const eventId = _self.premisEventData.getId()
      const params = [agentId, eventId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.setEvent, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _setAgentRights(_self, _exists) {
    console.log('Setting Agent rights ' + _exists)
    if ( _exists ) {
      _self.numWrites += 1
    } else {
      const agentId = _self.premisAgentData.getId()
      const rightsId = _self.premisRightsData.getId()
      const params = [agentId, rightsId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.setRights, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _agentSubmit () {
    const agentId = this.premisAgentData.getId()
    const hash = this.premisObjectData.getHash()
    const eventId = this.premisEventData.getId()
    const rightsId = this.premisRightsData.getId()
    let params = [agentId]
    this.web3Handler.callParamHandler(this, this.premisAgent.getAgentExists, params, this._setAgent, false)
    params = [agentId, hash]
    this.web3Handler.callParamHandler(this, this.premisAgent.getAgentObjectExists, params, this._setAgentObject, false)
    params = [agentId, eventId]
    this.web3Handler.callParamHandler(this, this.premisAgent.getAgentEventExists, params, this._setAgentEvent, false)
    params = [agentId, rightsId]
    this.web3Handler.callParamHandler(this, this.premisAgent.getAgentRightsExists, params, this._setAgentRights, false)
  }

  // Premis Rights Stuff

  _setRights(_self, _exists) {
    console.log('Setting rights ' + _exists)
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
      let params = [rightsId, hash, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.setObject, params, _self._writeCounter, _self.batchWrites)
      params = [rightsId, basis, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.setBasis, params, _self._writeCounter, _self.batchWrites)
      params = [rightsId, status, code, date, note, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.setCopyrightInfo, params, _self._writeCounter, _self.batchWrites)
      params = [rightsId, act, restriction, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.setRightsGranted, params, _self._writeCounter, _self.batchWrites)
      params = [rightsId, agentId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.setAgent, params, _self._writeCounter, _self.batchWrites)
    }
  }

  _rightsSubmit () {
    const rightsId = this.premisRightsData.getId()
    let params = [rightsId]
    this.web3Handler.callParamHandler(this, this.premisRights.getRightsExists, params, this._setRights, false)
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
      this.web3Handler.destroyBatch()
      this.web3Handler.createBatch()
      this._resetNumWrites()
      this._objectSubmit()
      this._eventSubmit()
      this._agentSubmit()
      this._rightsSubmit()
    }
  }
}

export default ContractWriter
