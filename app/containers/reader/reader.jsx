import React from 'react'
import PropTypes from 'prop-types'
import HashFile from '../helpers/hashFile'
import ObjectReader from './objectReader'
import {ReaderStrings} from '../../helpers/outputStrings'
import {IOPlainTextOutput} from '../../components/io'

class Reader extends React.Component {

  constructor (props) {
    super(props)

    const info = ReaderStrings.info

    this.state = {
      info: info,
      hash: undefined
    }
  }

  handleInfo (_info) {
    this.setState({ info: _info })
  }

  handleHashFile (_hash) {
    this.setState({hash: _hash})
  }

  render () {
    return (
      <div>
        <div>
          <IOPlainTextOutput text={this.state.info} />
          <hr />
        </div>
        <div>
          <HashFile parentFunc={this.handleHashFile.bind(this)} />
          <hr />
          <ObjectReader messageFunc={this.handleInfo.bind(this)} hash={this.state.hash} contracts={this.props.contracts}  web3={this.props.web3} />
          <hr />
        </div>
        <div>
          <IOPlainTextOutput text={this.state.info} />
        </div>
      </div>
    )
  }
}

Reader.propTypes = {
  contracts: PropTypes.object,
  web3: PropTypes.object
}

export default Reader
