import React from 'react'
import PropTypes from 'prop-types'
import {AppStrings, AppEventStrings} from '../../helpers/outputStrings'
import {IOLogger, IOHeading, IOPlainTextOutput} from '../../components/io'

class Events extends React.Component {

  constructor (props) {
    super(props)

    this.premisObject = this.props.contracts.getPremisObject()
    this.premisEvent = this.props.contracts.getPremisEvent()
    this.premisAgent = this.props.contracts.getPremisAgent()
    this.premisRights = this.props.contracts.getPremisRights()

    this.state = {
      log: []
    }
  }

  _logResult (result) {
    // const time = Date.now().toString()
    const date = new Date().toString()
    const logResults = date + ': ' + result
    let logs = this.state.log
    logs.push(logResults)
    // console.log(logResults)
    this.setState({ log: logs })
  }

  _allObjectEvents () {
    const self = this
    let allObjectEvt = this.premisObject.allEvents({}, {fromBlock: 1, toBlock: 'latest'})
    allObjectEvt.watch(function (e, result) {
      self._logResult(result)
    })
  }

  _allAgentEvents () {
    const self = this
    let allAgentEvt = this.premisAgent.allEvents({}, {fromBlock: 1, toBlock: 'latest'})
    allAgentEvt.watch(function (e, result) {
      self._logResult(result)
    })
  }

  _allEventsEvents () {
    const self = this
    let allEventEvt = this.premisEvent.allEvents({}, {fromBlock: 1, toBlock: 'latest'})
    allEventEvt.watch(function (e, result) {
      self._logResult(result)
    })
  }

  _allRightsEvents () {
    const self = this
    let allRightsEvt = this.premisRights.allEvents({}, {fromBlock: 1, toBlock: 'latest'})
    allRightsEvt.watch(function (e, result) {
      self._logResult(result)
    })
  }

  render () {
    return (
      <div>
        <div className="info">
          <IOHeading heading={AppStrings.heading} />
          <IOPlainTextOutput text={AppEventStrings.info} />
          <hr />
        </div>
        <div>
          <IOLogger log={this.state.log} />
        </div>
      </div>
    )
  }

  componentDidMount () {
    this._allObjectEvents()
    this._allAgentEvents()
    this._allEventsEvents()
    this._allRightsEvents()
  }

  componentWillUnmount () {

    this.premisObject.allEvents().stopWatching()
    this.premisAgent.allEvents().stopWatching()
    this.premisEvent.allEvents().stopWatching()
    this.premisRights.allEvents().stopWatching()

    this.props.web3.getWeb3().reset()
  }

}

Events.propTypes = {
  contracts: PropTypes.object,
  web3: PropTypes.object
}

export default Events
