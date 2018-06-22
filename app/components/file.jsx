import React from 'react'
import ReactDOM from 'react-dom'
import PropTypes from 'prop-types'
import FileReaderInput from 'react-file-reader-input'
import {Tooltip, Button} from 'antd'
import {IOButton} from './io'
import {ObjectWriterStrings} from '../helpers/outputStrings'

class File extends React.Component {

  constructor (props) {
    super(props)

    this.state = {
      loading: false
    }
  }

  _handleSetFilename (e, results) {
    //console.log(e, results)
    this.props.parentFunc(e, results)
    this.setState({loading: false})
  }

  _setLoading () {
    const loading = !(this.state.loading)
    this.setState({loading: loading})
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
              <Button
                loading={this.state.loading}
                onClick={this._setLoading.bind(this)}
                type='normal'
                icon={this.props.icon}
              >
                {ObjectWriterStrings.fileLoad}
              </Button>
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
  icon: PropTypes.string,
  tip: PropTypes.string
}

export default File
