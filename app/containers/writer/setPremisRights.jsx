import React from 'react'
import PropTypes from 'prop-types'
import PremisRightsHandler from '../../utils/premisRightsHandler'
import {RightsStrings} from '../../helpers/outputStrings'
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
        <IOHeading heading={RightsStrings.heading} />
        <IOSelect
          parentFunc={this._handleBasis.bind(this)}
          selections={RightsStrings.basisOptions}
          label={RightsStrings.basisLabel}
          tip={RightsStrings.basisTip}
        />
        <IOTextInput
          parentFunc={this._handleStatus.bind(this)}
          placeHolder={RightsStrings.statusPlaceHolder}
          label={RightsStrings.statusLabel}
          tip={RightsStrings.statusTip}
        />
        <IOSelect
          parentFunc={this._handleCountryCode.bind(this)}
          selections={RightsStrings.jurisdictionCountryCodes}
          label={RightsStrings.jurisdictionLabel}
          tip={RightsStrings.countryCodeTip}
        />
        <IOTextInput
          parentFunc={this._handleDeterminationDate.bind(this)}
          placeHolder={RightsStrings.determinationDatePlaceHolder}
          label={RightsStrings.determinationDateLabel}
          tip={RightsStrings.determinationDateTip}
        />
        <IOTextAreaInput
          parentFunc={this._handleNote.bind(this)}
          placeHolder={RightsStrings.notePlaceHolder}
          label={RightsStrings.noteLabel}
          tip={RightsStrings.noteTip}
        />
        <IOSelect
          parentFunc={this._handleRightsGrantedAct.bind(this)}
          selections={RightsStrings.actOptions}
          label={RightsStrings.actLabel}
          tip={RightsStrings.actTip}
        />
        <IOSelect
          parentFunc={this._handleRightsGrantedRestriction.bind(this)}
          selections={RightsStrings.restrictionsOptions}
          label={RightsStrings.restrictionsLabel}
          tip={RightsStrings.restrictionTip}
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
