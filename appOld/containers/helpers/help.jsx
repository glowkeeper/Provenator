import React from 'react'
import {HelpStrings} from '../../helpers/outputStrings'
import {IOHeading, IOPlainTextOutput} from '../../components/io'

class Help extends React.Component {

  constructor (props) {
    super(props)
  }

  render () {
    return (
      <div>
        <IOHeading heading={HelpStrings.heading} />
        <IOPlainTextOutput text={HelpStrings.info} />
      </div>
    )
  }
}

export default Help
