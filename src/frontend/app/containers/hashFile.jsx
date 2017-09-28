import React from 'react'
import PropTypes from 'prop-types'
import { keccak256 } from 'js-sha3';
import File from '../components/file'
import {PremisTextOutput} from '../components/premis'
import {HashStrings} from '../utils/outputStrings'

class HashFile extends React.Component {

  constructor (props) {
    super(props)

    this.state = {
      fileName: undefined,
      hash: undefined
    }
  }

  _handleSubmitFile (e, results) {
    // console.log(e)
    results.forEach(result => {
      const [thisProgressEvent, file] = result
      const buffer = Buffer.from(thisProgressEvent.target.result)
      const hash = keccak256(buffer)
      const fileName = file.name
      this.setState({fileName: fileName})
      this.setState({hash: hash})
      this.props.parentFunc(hash)
    })
  }

  render () {
    return (
      <div>
        <File heading={HashStrings.heading} parentFunc={this._handleSubmitFile.bind(this)} label={HashStrings.browseFileLabel} buttonLabel={HashStrings.browseFileButtonLabel} />
        <PremisTextOutput label={HashStrings.fileLabel} text={this.state.fileName} />
        <PremisTextOutput label={HashStrings.hashLabel} text={this.state.hash} />
      </div>
    )
  }
}

HashFile.propTypes = {
  parentFunc: PropTypes.func
}

export default HashFile
