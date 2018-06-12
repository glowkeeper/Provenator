import Web3Handler from './web3Handler'
import ContractHandler from '../utils/contractHandler'

import {ObjectReaderStrings} from '../helpers/outputStrings'

class ContractReader {

  constructor (_web3Handler, _contractHandler) {
    //console.log(_web3)

    this.web3Handler = _web3Handler
    const thisWeb3 = this.web3Handler.getWeb3()

    const contractHandler = _contractHandler
    this.premisObject = contractHandler.getPremisObject()
    this.premisEvent = contractHandler.getPremisEvent()
    this.premisAgent = contractHandler.getPremisAgent()
    this.premisRights = contractHandler.getPremisRights()

    this.reader = undefined
    this.readKeyCb = undefined
    this.readNoKeyCb = undefined
    this.ipfsHash = undefined
  }

  // Functions for fetching data fromm the blockchain for the reader output
  _getCategory (_self, _category) {
    _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.OBJECT.value, _category[0], ObjectReaderStrings.categoryLabel, _category[1])
   }

   _getFormat (_self, _format) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.OBJECT.value, _format[0], ObjectReaderStrings.formatLabel, _format[1])
   }

   _getNumProperties (_self, _numProperties) {
     //console.log(_numProperties)
     _self.readNoKeyCb(_self.reader, _self.reader.OUTPUTDATA.OBJECT.value, ObjectReaderStrings.numPropertiesLabel, _numProperties)
     for (let i = 0; i < _numProperties; i++) {
       const params = [_self.ipfsHash, i]
       _self.web3Handler.callParamHandler(_self, _self.premisObject.getProperties, params, _self._getProperties, false)
     }
   }

   _getProperties (_self, _properties) {
     const props = _properties[1] + ' - ' + _properties[2]
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.OBJECT.value, _properties[0], ObjectReaderStrings.propertiesLabel, props)
   }

   _getObjectData () {
     // console.log(this.premisObject)
     // console.log(this.ipfsHash)
     const params = [this.ipfsHash]
     this.web3Handler.callParamHandler(this, this.premisObject.getCategory, params, this._getCategory, false)
     this.web3Handler.callParamHandler(this, this.premisObject.getFormat, params, this._getFormat, false)
     this.web3Handler.callParamHandler(this, this.premisObject.getNumProperties, params, this._getNumProperties, false)
   }

   _getNumEvents (_self, _numEvents) {
     _self.readNoKeyCb(_self.reader, _self.reader.OUTPUTDATA.EVENT.value, ObjectReaderStrings.numEventsLabel, _numEvents)
     for (let i = 0; i < _numEvents; i++) {
       const params = [_self.ipfsHash, i]
       _self.web3Handler.callParamHandler(_self, _self.premisObject.getEvent, params, _self._getEvent, false)
     }
   }

   _getEvent (_self, _eventId) {
     //console.log(_eventId)
     const key = _eventId[0]
     const id =  _eventId[1]
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.EVENT.value, key, ObjectReaderStrings.eventLabel, id)
     const params = [id]
     _self.web3Handler.callParamHandler(_self, _self.premisEvent.getObject, params, _self._getEventObject, false)
     _self.web3Handler.callParamHandler(_self, _self.premisEvent.getType, params, _self._getEventType, false)
     _self.web3Handler.callParamHandler(_self, _self.premisEvent.getAgent, params, _self._getEventAgent, false)
     _self.web3Handler.callParamHandler(_self, _self.premisEvent.getTime, params, _self._getEventTime, false)
   }

   _getEventObject (_self, _eventObject) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.EVENT.value, _eventObject[0], ObjectReaderStrings.eventObjectLabel, _eventObject[1])
   }

   _getEventType (_self, _eventType) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.EVENT.value, _eventType[0], ObjectReaderStrings.eventTypeLabel, _eventType[1])
   }

   _getEventAgent(_self, _eventAgent) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.EVENT.value, _eventAgent[0], ObjectReaderStrings.eventAgentLabel, _eventAgent[1])
   }

   _getEventTime(_self, _eventTime) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.EVENT.value, _eventTime[0], ObjectReaderStrings.eventTimeLabel, _eventTime[1])
   }

   _getEventData () {
     const params = [this.ipfsHash]
     this.web3Handler.callParamHandler(this, this.premisObject.getNumEvents, params, this._getNumEvents, false)
   }

   _getNumAgents (_self, _numAgents) {
     _self.readNoKeyCb(_self.reader, _self.reader.OUTPUTDATA.AGENT.value, ObjectReaderStrings.numAgentsLabel, _numAgents)
     for (let i = 0; i < _numAgents; i++) {
       const params = [_self.ipfsHash, i]
       _self.web3Handler.callParamHandler(_self, _self.premisObject.getAgent, params, _self._getAgent, false)
     }
   }

   _getAgent(_self, _agentId) {
     const key = _agentId[0]
     const id =  _agentId[1]
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.AGENT.value, key, ObjectReaderStrings.agentLabel, id )
     const params = [id]
     _self.web3Handler.callParamHandler(_self, _self.premisAgent.getName, params, _self._getAgentName, false)
     _self.web3Handler.callParamHandler(_self, _self.premisAgent.getType, params, _self._getAgentType, false)
   }

   _getAgentName (_self, _agentName) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.AGENT.value, _agentName[0], ObjectReaderStrings.agentNameLabel, _agentName[1])
   }

   _getAgentType (_self, _agentType) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.AGENT.value, _agentType[0], ObjectReaderStrings.agentTypeLabel, _agentType[1])
   }

   _getAgentData () {
     const params = [this.ipfsHash]
     this.web3Handler.callParamHandler(this, this.premisObject.getNumAgents, params, this._getNumAgents, false)
   }

   _getNumRights(_self, _numRights) {
     _self.readNoKeyCb(_self.reader, _self.reader.OUTPUTDATA.RIGHTS.value, ObjectReaderStrings.numRightsLabel, _numRights)
     for (let i = 0; i < _numRights; i++) {
       const params = [_self.ipfsHash, i]
       _self.web3Handler.callParamHandler(_self, _self.premisObject.getRights, params, _self._getRights, false)
     }
   }

   _getRights (_self, _rightsId) {
     const key = _rightsId[0]
     const id =  _rightsId[1]
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.RIGHTS.value, key, ObjectReaderStrings.rightsLabel, id)
     const params = [id]
     _self.web3Handler.callParamHandler(_self, _self.premisRights.getObject, params, _self._getRightsObject, false)
     _self.web3Handler.callParamHandler(_self, _self.premisRights.getRightsBasis, params, _self._getRightsBasis, false)
     _self.web3Handler.callParamHandler(_self, _self.premisRights.getCopyrightStatus, params, _self._getCopyrightStatus, false)
     _self.web3Handler.callParamHandler(_self, _self.premisRights.getCopyrightJurisdiction, params, _self._getCopyrightJurisdiction, false)
     _self.web3Handler.callParamHandler(_self, _self.premisRights.getCopyrightDeterminationDate, params, _self._getCopyrightDeterminationDate, false)
     _self.web3Handler.callParamHandler(_self, _self.premisRights.getCopyrightNote, params, _self._getCopyrightNote, false)
     _self.web3Handler.callParamHandler(_self, _self.premisRights.getGrantedAct, params, _self._getGrantedAct, false)
     _self.web3Handler.callParamHandler(_self, _self.premisRights.getGrantedRestriction, params, _self._getGrantedRestriction, false)
     _self.web3Handler.callParamHandler(_self, _self.premisRights.getAgent, params, _self._getRightsAgent, false)
   }

   _getRightsObject (_self, _rightsObject) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.RIGHTS.value, _rightsObject[0], ObjectReaderStrings.rightsObjectLabel, _rightsObject[1])
   }

   _getRightsBasis (_self, _rightsBasis) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.RIGHTS.value, _rightsBasis[0], ObjectReaderStrings.rightsBasisLabel, _rightsBasis[1])
   }

   _getCopyrightStatus (_self, _rightsCopyrightStatus) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.RIGHTS.value, _rightsCopyrightStatus[0], ObjectReaderStrings.rightsCopyrightStatusLabel, _rightsCopyrightStatus[1])
   }

   _getCopyrightJurisdiction (_self, _rightsCopyrightJurisdiction) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.RIGHTS.value, _rightsCopyrightJurisdiction[0], ObjectReaderStrings.rightsCopyrightJurisdictionLabel, _rightsCopyrightJurisdiction[1])
   }

   _getCopyrightDeterminationDate (_self, _rightsCopyrightDeterminationDate) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.RIGHTS.value, _rightsCopyrightDeterminationDate[0], ObjectReaderStrings.rightsCopyrightDeterminationDateLabel, _rightsCopyrightDeterminationDate[1])
   }

   _getCopyrightNote(_self, _rightsCopyrightNote) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.RIGHTS.value, _rightsCopyrightNote[0], ObjectReaderStrings.rightsCopyrightNoteLabel, _rightsCopyrightNote[1])
   }

   _getGrantedAct(_self, _rightsGrantedAct) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.RIGHTS.value, _rightsGrantedAct[0], ObjectReaderStrings.rightsGrantedActLabel, _rightsGrantedAct[1])
   }

   _getGrantedRestriction (_self, _rightsGrantedRestriction) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.RIGHTS.value, _rightsGrantedRestriction[0], ObjectReaderStrings.rightsGrantedRestrictionLabel, _rightsGrantedRestriction[1])
   }

   _getRightsAgent (_self, _rightsAgent) {
     _self.readKeyCb(_self.reader, _self.reader.OUTPUTDATA.RIGHTS.value, _rightsAgent[0], ObjectReaderStrings.rightsAgentLabel, _rightsAgent[1])
   }

   _getRightsData () {
     const params = [this.ipfsHash]
     this.web3Handler.callParamHandler(this, this.premisObject.getNumRights, params, this._getNumRights, false)
   }

   getReaderData(_callee, _keyCb, _noKeyCb, ipfsHash) {

     if ((typeof _keyCb === 'function') && (typeof _noKeyCb === 'function')) {

       this.reader = _callee
       this.ipfsHash = ipfsHash
       this.readKeyCb = _keyCb
       this.readNoKeyCb = _noKeyCb

       this._getObjectData()
       this._getEventData()
       this._getAgentData()
       this._getRightsData()
     }
   }
}

export default ContractReader
