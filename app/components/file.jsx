import React from 'react'
import ReactDOM from 'react-dom'
import PropTypes from 'prop-types'
import FileReaderInput from 'react-file-reader-input'
import {Tooltip} from 'antd'
import {IOButton} from './io'

class File extends React.Component {

  /* constructor (props) {
    super(props)
  } */

  _handleSetFilename (e, results) {
    this.props.parentFunc(e, results)
  }

  render () {

    return (
      <div>
        <h2>{this.props.heading}</h2>
        <div>
          <FileReaderInput
            as="binary"
            id="my-file-input"
            onChange={this._handleSetFilename.bind(this)}>
            <Tooltip title={this.props.tip}>
              <IOButton icon={null} onClick={null}>
                Load
              </IOButton>
            </Tooltip>
          </FileReaderInput>
        </div>
      </div>
    )
  }
}

File.propTypes = {
  parentFunc: PropTypes.func,
  heading: PropTypes.string,
  tip: PropTypes.string
}

export default File
