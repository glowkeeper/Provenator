import React from 'react'
import PropTypes from 'prop-types'
import PremisRightsHandler from '../../utils/premisRightsHandler'
import {PremisRightsStrings} from '../../helpers/outputStrings'
import {IOHeading, IOTextInput, IOTextAreaInput, IOSelect} from '../../components/io'

class SetPremisRights extends React.Component {

  constructor (props) {
    super(props)
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
        <IOHeading heading={PremisRightsStrings.heading} />
        <IOSelect
          parentFunc={this._handleBasis.bind(this)}
          selections={PremisRightsHandler.basisOptions}
          label={PremisRightsStrings.basisLabel}
          tip={PremisRightsStrings.basisTip}
        />
        <IOTextInput parentFunc={this._handleStatus.bind(this)} placeHolder={PremisRightsStrings.statusPlaceHolder} label={PremisRightsStrings.statusLabel} tip={PremisRightsStrings.statusTip} />
        <IOSelect
          parentFunc={this._handleCountryCode.bind(this)}
          selections={PremisRightsStrings.jurisdictionCountryCodes}
          label={PremisRightsStrings.jurisdictionLabel}
          tip={PremisRightsStrings.countryCodeTip}
        />
        <IOTextInput parentFunc={this._handleDeterminationDate.bind(this)} placeHolder={PremisRightsStrings.determinationDatePlaceHolder} label={PremisRightsStrings.determinationDateLabel} tip={PremisRightsStrings.determinationDateTip} />
        <IOTextAreaInput parentFunc={this._handleNote.bind(this)} placeHolder={PremisRightsStrings.notePlaceHolder} label={PremisRightsStrings.noteLabel} tip={PremisRightsStrings.noteTip} />
        <IOSelect
          parentFunc={this._handleRightsGrantedAct.bind(this)}
          selections={PremisRightsHandler.actOptions}
          label={PremisRightsStrings.actLabel}
          tip={PremisRightsStrings.actTip}
        />
        <IOSelect
          parentFunc={this._handleRightsGrantedRestriction.bind(this)}
          selections={PremisRightsHandler.restrictionsOptions}
          label={PremisRightsStrings.restrictionsLabel}
          tip={PremisRightsStrings.restrictionTip}
        />
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
