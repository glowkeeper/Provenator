import React from 'react'
import PropTypes from 'prop-types'
import PremisAgentHandler from '../../utils/premisAgentHandler'
import {AgentStrings} from '../../helpers/outputStrings'
import {IOHeading, IOTextInput, IOSelect} from '../../components/io'

class SetPremisAgent extends React.Component {

  constructor (props) {
    super(props)
  }

  _handleName (_name) {
    this.props.agentHandler.setName(_name)
  }

  _handleType (_selection) {
    //console.log(_selection)
    this.props.agentHandler.setType(_selection)
  }

  render () {
    return (
      <div>
        <IOHeading heading={AgentStrings.heading} />
        <IOTextInput
          parentFunc={this._handleName.bind(this)}
          placeHolder={AgentStrings.namePlaceHolder}
          label={AgentStrings.nameLabel}
          tip={AgentStrings.nameTip}
        />
        <IOSelect
          parentFunc={this._handleType.bind(this)}
          selections={AgentStrings.typeOptions}
          label={AgentStrings.agentTypeLabel}
          tip={AgentStrings.typeTip}
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
