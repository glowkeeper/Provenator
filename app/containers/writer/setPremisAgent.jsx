import React from 'react'
import PropTypes from 'prop-types'
import PremisAgentHandler from '../../utils/premisAgentHandler'
import {PremisAgentStrings} from '../../helpers/outputStrings'
import {IOHeading, IOTextInput, IOSelect} from '../../components/io'

class SetPremisAgent extends React.Component {

  constructor (props) {
    super(props)
  }

  _handleName (_name) {
    this.props.agentHandler.setName(_name)
  }

  _handleType (_selection) {
    // console.log(_selection.value)
    this.props.agentHandler.setType(_selection.label)
    this.setState({typeId: _selection.value})
  }

  render () {
    return (
      <div>
        <IOHeading heading={PremisAgentStrings.heading} />
        <IOTextInput parentFunc={this._handleName.bind(this)} placeHolder={PremisAgentStrings.namePlaceHolder} label={PremisAgentStrings.nameLabel} tip={PremisAgentStrings.nameTip} />
        <IOSelect
          parentFunc={this._handleType.bind(this)}
          selections={PremisAgentHandler.typeOptions}
          label={PremisAgentStrings.agentTypeLabel}
          tip={PremisAgentStrings.typeTip}
        />
      </div>
    )
  }

  /* componentWillUpdate() {

  } */
}

SetPremisAgent.propTypes = {
  web3: PropTypes.object,
  agentHandler: PropTypes.object
}

export default SetPremisAgent
