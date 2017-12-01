import React from 'react'
import PropTypes from 'prop-types'

import FileReaderInput from 'react-file-reader-input'
import { rTComponents } from './theme'
import { Button } from 'react-toolbox/lib/button'

class File extends React.Component {

  /* constructor (props) {
    super(props)
  } */

  _handleSetFilename (e, results) {
    this.props.parentFunc(e, results)
  }

  render () {
    return (
      <div className="section">
        <h2>{this.props.heading}</h2>
        <p>{this.props.label}</p>
        <div className="RTButton.buttonPrimary">
          <FileReaderInput
            as="binary"
            id="my-file-input"
            onChange={this._handleSetFilename.bind(this)}>
            <Button  className={rTComponents.buttonPrimary} raised ripple>{this.props.buttonLabel}</Button>
          </FileReaderInput>
        </div>
      </div>
    )
  }
}

File.propTypes = {
  parentFunc: PropTypes.func,
  heading: PropTypes.string,
  label: PropTypes.string,
  buttonLabel: PropTypes.string
}

export default File
