import React from 'react'
import PropTypes from 'prop-types'

class Logger extends React.Component {

  /* constructor(props) {
    super(props)
  } */

  render () {
    let logs = this.props.log.map((text, index) =>
      <span key={index}>
        {text}
        <br />
      </span>
    )
    return (
      <div className="section">
        <h2 className="section-heading">{this.props.heading}</h2>
        <p>{logs}</p>
      </div>
    )
  }
}

Logger.propTypes = {
  log: PropTypes.array,
  heading: PropTypes.string
}

export default Logger
