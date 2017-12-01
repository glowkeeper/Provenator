import React from 'react'
import {OverviewStrings} from '../utils/outputStrings'
import {PremisHeading} from '../components/premis'
import {PremisPlainTextOutput} from '../components/premis'

class Overview extends React.Component {

  constructor (props) {
    super(props)
  }

  render () {
    return (
      <div>
        <PremisHeading heading={OverviewStrings.heading} />
        <PremisPlainTextOutput text={OverviewStrings.info} />
        <hr />
      </div>
    )
  }
}

export default Overview
