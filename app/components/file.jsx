import React from 'react'
import ReactDOM from 'react-dom'
import PropTypes from 'prop-types'
import FileReaderInput from 'react-file-reader-input'

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
        <div className="RTButton.buttonPrimary">
          <FileReaderInput
            as="binary"
            id="my-file-input"
            onChange={this._handleSetFilename.bind(this)}>
            <Tooltip title={this.props.tip} position="right">
              <Button className={rTComponents.buttonPrimary} raised ripple><MdFileUpload/></Button>
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
