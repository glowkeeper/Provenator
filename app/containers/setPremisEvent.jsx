import React from 'react'
import PropTypes from 'prop-types'
import PremisEventHandler from '../utils/premisEventHandler'
//import {PremisHeading, PremisTextInput, PremisTextOutput, PremisSelectPlus} from '../components/premis'
import {PremisEventStrings} from '../helpers/outputStrings'
import {IOHeading, IOTextInput, IOSelect} from '../components/io'

class SetPremisEvent extends React.Component {

  constructor (props) {
    super(props)

    this.state = {
      eventId: undefined,
      eventSelections: []
    }
  }

  setEventTypeName (_self, _name) {
    // console.log(_name)
    let eventSelections = _self.state.eventSelections
    eventSelections.push({ value: eventSelections.length, label: _name })
    _self.setState({eventSelections: eventSelections})
  }

  _handleEvent (_selection) {
    // console.log('Event Id is ' + selection.value + '. Event Name is ' + selection.label)
    this.props.eventHandler.setName(_selection.label)
    this.setState({eventId: _selection.value})
  }

  _handleEventDate (_date) {
    // console.log('Event date is ' + date)
    this.props.eventHandler.setDate(_date)
  }

  render () {
    return (
      <div>
        <IOHeading heading={PremisEventStrings.heading} />
        <IOSelect parentFunc={this._handleEvent.bind(this)} placeHolder={PremisEventStrings.eventPlaceHolder} tip={PremisEventStrings.eventTip} selections={this.state.eventSelections} selection={this.state.eventId} />
        <IOTextInput parentFunc={this._handleEventDate.bind(this)} placeHolder={PremisEventStrings.datePlaceHolder} label={PremisEventStrings.dateLabel} tip={PremisEventStrings.dateTip} />
      </div>
    )
  }

  /* componentWillUpdate() {

  } */

  componentDidMount () {
    this.props.contractWriter.getEventTypes(this, this.setEventTypeName)
  }

}

SetPremisEvent.propTypes = {
  contractWriter: PropTypes.object,
  eventHandler: PropTypes.object
}

export default SetPremisEvent
