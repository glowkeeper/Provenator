import React from 'react'
import {OverviewStrings} from '../helpers/outputStrings'
import {IOHeading, IOPlainTextOutput} from '../components/io'

class Overview extends React.Component {

  constructor (props) {
    super(props)
  }

  render () {
    return (
      <div>
        <IOHeading heading={OverviewStrings.heading} />
        <IOPlainTextOutput text={OverviewStrings.info} />
        <hr />
      </div>
    )
  }
}

export default Overview
