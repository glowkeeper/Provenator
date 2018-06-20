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

    this.maxIndex = Math.pow(2, 256) - 1
    // console.log('Max index is ' + this.maxIndex)

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

    this.eventTypeCallee = undefined
    this.eventTypeCb = undefined

    this.propTypeCallee = undefined
    this.propTypeCb = undefined

    this.defaultTO = {gas: 300000}

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
    this.numWrites = 0;
  }

  _blockchainWriter (_self, _result) {
    _self.numWrites += 1;
    // console.log(_self.numWrites)
  }

  _isNewIndex (_index) {
    // console.log('In index checker ' + _index)
    if ( _index == this.maxIndex ) {
      // console.log('Index ' + _index + ' Matches ' + this.maxIndex)
      return true
    } else {
      // console.log('No match')
      return false
    }
  }

  // Premis Object Stuff

  _setPropType(_self, _propTypeIndex) {
    // console.log('Checking new property type exists with index ' + _propTypeIndex)
    // console.log('Setting Object proptype ' + _propTypeIndex)
    if ( _self._isNewIndex(_propTypeIndex) ) {
      const propType = _self.premisObjectData.getPropertyType()
      const params = [propType, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.setPropertyType, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 1
      // console.log(_self.numWrites)
    }
  }

  _setObject(_self, _objectIndex) {
    // console.log('Object index ' + _objectIndex)
    if ( _self._isNewIndex(_objectIndex) ) {
      const hash = _self.premisObjectData.getHash()
      const category = _self.premisObjectData.getCategory()
      const format = _self.premisObjectData.getFormat()
      const params = [hash, category, format, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.setObject, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 1
      // console.log(_self.numWrites)
    }
  }

  _setObjectEvent(_self, _eventIndex) {
    // console.log('Setting Object event ' + _eventIndex)
    if ( _self._isNewIndex(_eventIndex) ) {
      const hash = _self.premisObjectData.getHash()
      const eventId = _self.premisEventData.getId()
      const params = [hash, eventId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.setEvent, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 1
      // console.log(_self.numWrites)
    }
  }

  _setObjectAgent(_self, _agentIndex) {
    // console.log('Setting Object agent ' + _agentIndex)
    if ( _self._isNewIndex(_agentIndex) ) {
      const hash = _self.premisObjectData.getHash()
      const agentId = _self.premisAgentData.getId()
      const params = [hash, agentId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.setAgent, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 1
      // console.log(_self.numWrites)
    }
  }

  _setObjectRights(_self, _rightsIndex) {
    // console.log('Object Rights Index ' + _rightsIndex)
    if ( _self._isNewIndex(_rightsIndex) ) {
      const hash = _self.premisObjectData.getHash()
      const rightsId = _self.premisRightsData.getId()
      const params = [hash, rightsId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisObject.setRights, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 1
      // console.log(_self.numWrites)
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
    this.web3Handler.callParamHandler(this, this.premisObject.setProperties, params, this._blockchainWriter, false)
    params = [hash, eventId]
    this.web3Handler.callParamHandler(this, this.premisObject.getObjectEventExists, params, this._setObjectEvent, false)
    params = [hash, agentId]
    this.web3Handler.callParamHandler(this, this.premisObject.getObjectAgentExists, params, this._setObjectAgent, false)
    params = [hash, rightsId]
    this.web3Handler.callParamHandler(this, this.premisObject.getObjectRightsExists, params, this._setObjectRights, false)
  }

  // Premis Event Stuff

  _setEventType(_self, _eventTypeIndex) {
    // console.log('Setting event type ' + _eventTypeIndex)
    if ( _self._isNewIndex(_eventTypeIndex) ) {
      const eventTypeName = _self.premisEventData.getType()
      const params = [eventTypeName, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisEvent.setEventType, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 1
      // console.log(_self.numWrites)
    }
  }

  _setEvent(_self, _eventIndex) {
    // console.log('Setting Event ' + _eventIndex)
    if ( _self._isNewIndex(_eventIndex) ) {
      const eventId = _self.premisEventData.getId()
      const eventTypeName = _self.premisEventData.getType()
      const date = _self.premisEventData.getDate()
      const hash = _self.premisObjectData.getHash()
      const agentId = _self.premisAgentData.getId()
      const params = [eventId, eventTypeName, date, hash, agentId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisEvent.setEvent, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 1
      // console.log(_self.numWrites)
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

  _setAgent(_self, _agentIndex) {
    // console.log('Setting Agent ' + _agentIndex)
    if ( _self._isNewIndex(_agentIndex) ) {
      const agentId = _self.premisAgentData.getId()
      const name = _self.premisAgentData.getName()
      const agentType = _self.premisAgentData.getType()
      const params = [agentId, name, agentType, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.setAgent, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 1
      // console.log(_self.numWrites)
    }
  }

  _setAgentObject(_self, _objectIndex) {
    // console.log('Setting Agent Object ' + _objectIndex)
    if ( _self._isNewIndex(_objectIndex) ) {
      const agentId = _self.premisAgentData.getId()
      const hash = _self.premisObjectData.getHash()
      const params = [agentId, hash, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.setObject, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 1
      // console.log(_self.numWrites)
    }
  }

  _setAgentEvent(_self, _eventIndex) {
    // console.log('Setting Agent event ' + _eventIndex)
    if ( _self._isNewIndex(_eventIndex) ) {
      const agentId = _self.premisAgentData.getId()
      const eventId = _self.premisEventData.getId()
      const params = [agentId, eventId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.setEvent, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 1
      // console.log(_self.numWrites)
    }
  }

  _setAgentRights(_self, _rightsIndex) {
    // console.log('Setting Agent rights ' + _rightsIndex)
    if ( _self._isNewIndex(_rightsIndex) ) {
      const agentId = _self.premisAgentData.getId()
      const rightsId = _self.premisRightsData.getId()
      const params = [agentId, rightsId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisAgent.setRights, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 1
      // console.log(_self.numWrites)
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

  _setRights(_self, _rightsIndex) {
    // console.log('Setting rights ' + _rightsIndex)
    if ( _self._isNewIndex(_rightsIndex) ) {
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
      _self.web3Handler.callParamHandler(_self, _self.premisRights.setObject, params, _self._blockchainWriter, false)
      params = [rightsId, basis, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.setBasis, params, _self._blockchainWriter, false)
      params = [rightsId, status, code, date, note, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.setCopyrightInfo, params, _self._blockchainWriter, false)
      params = [rightsId, act, restriction, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.setRightsGranted, params, _self._blockchainWriter, false)
      params = [rightsId, agentId, _self.defaultTO]
      _self.web3Handler.callParamHandler(_self, _self.premisRights.setAgent, params, _self._blockchainWriter, false)
    } else {
      _self.numWrites += 5
      // console.log(_self.numWrites)
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
    // this.web3Handler.createBatch()
    if ( this._checkData() ) {
      this._resetNumWrites()
      this._objectSubmit()
      this._eventSubmit()
      this._agentSubmit()
      this._rightsSubmit()
    }
    // this.web3Handler.executeBatch()
  }

  // Random functions to fill selection types for various 'writer' fields
  _getPropTypes (_self, _numPropTypes) {
    // console.log('num props: ' + _numPropTypes)
    for (let i = 0; i < _numPropTypes; i++) {
      _self.web3Handler.callParamHandler(_self.propTypeCallee, _self.premisObject.getPropertyTypeName, [i], _self.propTypeCb, false)
    }
  }

  getPropTypes(_callee, _cb) {
    this.propTypeCallee = _callee
    this.propTypeCb = _cb
    this.web3Handler.callHandler(this, this.premisObject.getNumPropTypes, this._getPropTypes, false)
  }

  _getEventTypes (_self, _numEventTypes) {
    // console.log('Num event types ' + _numEventTypes)
    for (let i = 0; i < _numEventTypes; i++) {
      _self.web3Handler.callParamHandler(_self.eventTypeCallee, _self.premisEvent.getEventType, [i], _self.eventTypeCb, false)
    }
  }

  getEventTypes(_callee, _cb) {
    this.eventTypeCallee = _callee
    this.eventTypeCb = _cb
    this.web3Handler.callHandler(this, this.premisEvent.getNumEventTypes, this._getEventTypes, false)
  }
}

export default ContractWriter
