import React from 'react'
import PropTypes from 'prop-types'
import PremisRightsHandler from '../utils/premisRightsHandler'
import {PremisHeading, PremisTextInput, PremisTextAreaInput, PremisSelect} from '../components/premis'
import {PremisRightsStrings} from '../utils/outputStrings'

class SetPremisRights extends React.Component {

  constructor (props) {
    super(props)

    const numBasis = PremisRightsHandler.basisOptions.length
    let basisSelections = []
    for (let i = 0; i < numBasis; i++) {
      basisSelections[i] = { value: i, label: PremisRightsHandler.basisOptions[i] }
    }

    const numActs = PremisRightsHandler.actOptions.length
    let actSelections = []
    for (let i = 0; i < numActs; i++) {
      actSelections[i] = { value: i, label: PremisRightsHandler.actOptions[i] }
    }

    const numRestrictions = PremisRightsHandler.restrictionsOptions.length
    let restrictionSelections = []
    for (let i = 0; i < numRestrictions; i++) {
      restrictionSelections[i] = { value: i, label: PremisRightsHandler.restrictionsOptions[i] }
    }

    const numCountryCodes = PremisRightsHandler.jurisdictionCountryCodes.length
    let countryCodeSelections = []
    for (let i = 0; i < numCountryCodes; i++) {
      countryCodeSelections[i] = { value: i, label: PremisRightsHandler.jurisdictionCountryCodes[i] }
    }

    this.state = {
      basisId: undefined,
      basisSelections: basisSelections,
      countryCodeId: undefined,
      countryCodeSelections: countryCodeSelections,
      actId: undefined,
      actSelections: actSelections,
      restrictionId: undefined,
      restrictionSelections: restrictionSelections
    }
  }

  // Set new values
  _handleBasis (_selection) {
    this.props.rightsHandler.setBasis(_selection.label)
    this.setState({basisId: _selection.value})
  }

  _handleStatus (_status) {
    this.props.rightsHandler.setStatus(_status)
  }

  _handleCountryCode (_selection) {
    this.props.rightsHandler.setCountryCode(_selection.label)
    this.setState({countryCodeId: _selection.value})
  }

  _handleDeterminationDate (_date) {
    this.props.rightsHandler.setDeterminationDate(_date)
  }

  _handleNote (_note) {
    this.props.rightsHandler.setNote(_note)
  }

  _handleRightsGrantedAct (_selection) {
    this.props.rightsHandler.setAct(_selection.label)
    this.setState({actId: _selection.value})
  }

  _handleRightsGrantedRestriction (_selection) {
    // console.log('Restriction Id is ' + selection.value + '. Restriction Name is ' + selection.label)
    // this.rights.restrictionName = selection.label
    this.props.rightsHandler.setRestriction(_selection.label)
    this.setState({restrictionId: _selection.value})
    // this.props.parentFunc(this.rights)
  }

  render () {
    return (
      <div>
        <PremisHeading heading={PremisRightsStrings.heading} />
        <PremisSelect parentFunc={this._handleBasis.bind(this)} placeHolder={PremisRightsStrings.basisPlaceHolder} tip={PremisRightsStrings.basisTip} selections={this.state.basisSelections} selection={this.state.basisId} />
        <PremisTextInput parentFunc={this._handleStatus.bind(this)} placeHolder={PremisRightsStrings.statusPlaceHolder} label={PremisRightsStrings.statusLabel} tip={PremisRightsStrings.statusTip} />
        <PremisSelect parentFunc={this._handleCountryCode.bind(this)} placeHolder={PremisRightsStrings.countryCodePlaceHolder} tip={PremisRightsStrings.countryCodeTip} selections={this.state.countryCodeSelections} selection={this.state.countryCodeId} />
        <PremisTextInput parentFunc={this._handleDeterminationDate.bind(this)} placeHolder={PremisRightsStrings.determinationDatePlaceHolder} label={PremisRightsStrings.determinationDateLabel} tip={PremisRightsStrings.determinationDateTip} />
        <PremisTextAreaInput parentFunc={this._handleNote.bind(this)} placeHolder={PremisRightsStrings.notePlaceHolder} label={PremisRightsStrings.noteLabel} tip={PremisRightsStrings.noteTip} />
        <PremisSelect parentFunc={this._handleRightsGrantedAct.bind(this)} placeHolder={PremisRightsStrings.actPlaceHolder} tip={PremisRightsStrings.actTip} selections={this.state.actSelections} selection={this.state.actId} />
        <PremisSelect parentFunc={this._handleRightsGrantedRestriction.bind(this)} placeHolder={PremisRightsStrings.restrictionPlaceHolder} tip={PremisRightsStrings.restrictionTip} selections={this.state.restrictionSelections} selection={this.state.restrictionId} />
      </div>
    )
  }

  /* componentWillUpdate() {

  } */
}

SetPremisRights.propTypes = {
  web3: PropTypes.object,
  rightsHandler: PropTypes.object
}

export default SetPremisRights
