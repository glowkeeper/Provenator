import React from 'react'
import {HomeStrings} from '../../helpers/outputStrings'
import {IOHeading, IOPlainTextOutput} from '../../components/io'

class Home extends React.Component {

  constructor (props) {
    super(props)
  }

  render () {
    return (
      <div>
        <IOHeading heading={HomeStrings.heading} />
        <IOPlainTextOutput text={HomeStrings.info} />
        <hr />
      </div>
    )
  }
}

export default Home
