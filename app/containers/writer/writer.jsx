import React from 'react'
import PropTypes from 'prop-types'
import HashFile from '../helpers/hashFile'
import ObjectWriter from './objectWriter'
import {WriterStrings} from '../../helpers/outputStrings'
import {IOPlainTextOutput} from '../../components/io'
import { Icon } from 'antd'

class Writer extends React.Component {

  constructor (props) {
    super(props)

    // console.log(this.props.premisObject)

    const info = WriterStrings.info
    //const info = 'Use this form to record information about a Fake News digital media object.'

    this.state = {
      info: info,
      spinning: false,
      hash: undefined
    }
  }

  handleInfo (_info, _isLoading) {
    this.setState({ info: _info })
    this.setState({ spinning: _isLoading })
  }

  handleHashFile (_hash) {
    this.setState({hash: _hash})
  }

  render () {
    return (
      <div>
        <Icon type={WriterStrings.icon} spin={this.state.spinning} /> {this.state.info}
        <hr />
        <HashFile parentFunc={this.handleHashFile.bind(this)} />
        <hr />
        <ObjectWriter messageFunc={this.handleInfo.bind(this)} hash={this.state.hash} contracts={this.props.contracts} web3={this.props.web3} />
        <hr />
        <Icon type={WriterStrings.icon} spin={this.state.spinning} /> {this.state.info}
      </div>
    )
  }
}

Writer.propTypes = {
  contracts: PropTypes.object,
  web3: PropTypes.object
}

export default Writer
