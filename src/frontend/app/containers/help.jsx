import React from 'react'
import {HelpStrings} from '../utils/outputStrings'
import {PremisHeading} from '../components/premis'
import {PremisPlainTextOutput} from '../components/premis'

class Help extends React.Component {

  constructor (props) {
    super(props)
  }

  render () {
    return (
      <div>
        <PremisHeading heading={HelpStrings.heading} />
        <PremisPlainTextOutput text={HelpStrings.info} />
        <hr />
      </div>
    )
  }
}

export default Help
