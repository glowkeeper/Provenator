import React from 'react'
import PropTypes from 'prop-types'
import ContractReader from '../utils/contractReader'
import Logger from '../components/logger'
import {ObjectReaderStrings} from '../utils/outputStrings'

class ObjectReader extends React.Component {

  constructor (props) {
    super(props)

    this.OUTPUTDATA = {
      OBJECT : {value: 0, name: 'Object', code: 'O'},
      EVENT: {value: 1, name: 'Event', code: 'E'},
      AGENT : {value: 2, name: 'Agent', code: 'A'},
      RIGHTS : {value: 3, name: 'Rights', code: 'R'}
    }
    Object.freeze(this.OUTPUTDATA)

    this.contractReader = new ContractReader(this.props.web3, this.props.contracts)

    this.objectData = []
    this.eventData = []
    this.agentData = []
    this.rightsData = []

    this.hash = undefined
    this.lastRead = Date.now()

    this.state = {
      outputData: []
    }
  }

  getEventString () {
    return this.eventString
  }

  getAgentString () {
    return this.agentString
  }

  getRightsString () {
    return this.rightsString
  }

  _outputData () {
    let outputData = this.state.outputData
    const space = ' '
    outputData.push(...this.objectData)
    outputData.push(space)
    outputData.push(ObjectReaderStrings.eventHeading)
    outputData.push(space)
    outputData.push(...this.eventData)
    outputData.push(space)
    outputData.push(ObjectReaderStrings.agentHeading)
    outputData.push(space)
    outputData.push(...this.agentData)
    outputData.push(space)
    outputData.push(ObjectReaderStrings.rightsHeading)
    outputData.push(space)
    outputData.push(...this.rightsData)
    this.setState({outputData: outputData})
  }

  _sleep (ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  async _checkWrite () {
    const wait = 3000
    while ((Date.now() - this.lastRead) < wait) {
      await this._sleep(wait)
    }
    this.props.messageFunc(ObjectReaderStrings.finishedFetching)
    this._outputData()
  }

  writeData (_self, _outputIndex, _label, _value) {
    const data = _label + _value
    switch(_outputIndex) {
      case 0:
          _self.objectData.push(data)
          break;
      case 1:
          _self.eventData.push(data)
          break;
      case 2:
          _self.agentData.push(data)
          break;
      case 3:
          _self.rightsData.push(data)
          break;
      default:
          console.error('Error - no such output data index: ' + _outputIndex)
    }
    _self.lastRead = Date.now()
  }

  writeKeyData (_self, _outputIndex, _key, _label, _value) {
    const data = _key + ' - ' + _label + _value
    switch(_outputIndex) {
      case 0:
          _self.objectData.push(data)
          break;
      case 1:
          _self.eventData.push(data)
          break;
      case 2:
          _self.agentData.push(data)
          break;
      case 3:
          _self.rightsData.push(data)
          break;
      default:
          console.error('Error - no such output data index: ' + _outputIndex)
    }
    _self.lastRead = Date.now()
  }

  _getData () {
    if (this.props.hash !== this.hash) {

      this.hash = this.props.hash

      this.props.messageFunc(ObjectReaderStrings.fetching)
      this.objectData = []
      this.eventData = []
      this.agentData = []
      this.rightsData = []
      this.setState({ outputData: [] })

      this.contractReader.getReaderData(this, this.writeKeyData, this.writeData, this.props.hash)

      this.lastRead = Date.now()
      this._checkWrite()
    }
  }

  render () {
    return (
      <div>
        <Logger heading={ObjectReaderStrings.heading} log={this.state.outputData} />
      </div>
    )
  }

  componentDidMount () {
    setInterval(this._getData.bind(this), 1500)
  }
}

ObjectReader.propTypes = {
  messageFunc: PropTypes.func,
  info: PropTypes.string,
  hash: PropTypes.string,
  contracts: PropTypes.object
}

export default ObjectReader
