import React from 'react'
import {AboutStrings} from '../../helpers/outputStrings'
import {IOHeading, IOPlainTextOutput} from '../../components/io'

class About extends React.Component {

  constructor (props) {
    super(props)
  }

  render () {
    return (
      <div>
        <IOHeading heading={AboutStrings.heading} />
        <IOPlainTextOutput text={AboutStrings.info} />
      </div>
    )
  }
}

export default About
