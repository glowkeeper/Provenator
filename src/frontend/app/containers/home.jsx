import React from 'react'
import {HomeStrings} from '../utils/outputStrings'
import {PremisHeading} from '../components/premis'
import {PremisPlainTextOutput} from '../components/premis'

class Home extends React.Component {

  constructor (props) {
    super(props)
  }

  render () {
    return (
      <div>
        <PremisHeading heading={HomeStrings.heading} />
        <PremisPlainTextOutput text={HomeStrings.info} />
        <hr />
      </div>
    )
  }
}

export default Home
