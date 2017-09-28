import React from 'react'
import {AboutStrings} from '../utils/outputStrings'
import {PremisHeading} from '../components/premis'
import {PremisPlainTextOutput} from '../components/premis'

class About extends React.Component {

  constructor (props) {
    super(props)
  }

  render () {
    return (
      <div>
        <PremisHeading heading={AboutStrings.heading} />
        <PremisPlainTextOutput text={AboutStrings.info} />
        <hr />
      </div>
    )
  }
}

export default About
