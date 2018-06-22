import React from 'react'
import PropTypes from 'prop-types'
import { keccak256 } from 'js-sha3'
import File from '../../components/file'
import {IOTextOutput} from '../../components/io'
import {HashStrings} from '../../helpers/outputStrings'

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
        <File heading={HashStrings.heading} parentFunc={this._handleSubmitFile.bind(this)} icon={HashStrings.icon} tip={HashStrings.browseFileTip} />
        <br/>
        <IOTextOutput label={HashStrings.fileLabel} text={this.state.fileName} />
        <IOTextOutput label={HashStrings.hashLabel} text={this.state.hash} />
      </div>
    )
  }
}

HashFile.propTypes = {
  parentFunc: PropTypes.func
}

export default HashFile
