import React from 'react'
import ReactDOM from 'react-dom'
import PropTypes from 'prop-types'
import classNames from 'classnames/bind'
import Select from 'react-select'

import '!style-loader!css-loader!react-select/dist/react-select.css'

import { rTComponents } from './theme'

import { Button } from 'react-toolbox/lib/button'
import MdDone from 'react-icons/lib/md/done'
import { Tooltip } from 'react-tippy'
import ReactMarkdown from 'react-markdown'

class AppHeading extends React.Component {

  /* constructor(props) {
    super(props)
  } */

  render () {
    return (
      <div className="heading">
        <h1 className="heading">{this.props.heading}</h1>
      </div>
    )
  }
}

AppHeading.propTypes = {
  heading: PropTypes.string
}

class PremisHeading extends React.Component {

  /* constructor(props) {
    super(props)
  } */

  render () {
    return (
      <div className="heading">
        <h2 className="section-heading">{this.props.heading}</h2>
      </div>
    )
  }
}

PremisHeading.propTypes = {
  heading: PropTypes.string
}

class PremisPlainTextOutput extends React.Component {

  constructor(props) {
    super(props)
  }

  render () {
    return (
      <ReactMarkdown escapeHtml={false} source={this.props.text} />
    )
  }
}

PremisPlainTextOutput.propTypes = {
  text: PropTypes.string
}


class PremisTextOutput extends React.Component {

  /* constructor(props) {
    super(props)
  } */

  render () {
    return (
      <p>{this.props.label} {this.props.text}</p>
    )
  }
}

PremisTextOutput.propTypes = {
  label: PropTypes.string,
  text: PropTypes.string
}

class PremisTextInput extends React.Component {

  /* constructor(props) {
    super(props)
  } */

  _handleChange (e) {
    this.props.parentFunc(e.target.value)
  }

  render () {
    return (
      <div className="textInput">
        <Tooltip title={this.props.tip} position="top">
          <p>{this.props.label}
            <input
              type="text"
              placeholder={this.props.placeHolder}
              onChange={this._handleChange.bind(this)}
            />
          </p>
        </Tooltip>
      </div>
    )
  }
}

PremisTextInput.propTypes = {
  parentFunc: PropTypes.func,
  placeHolder: PropTypes.string,
  label: PropTypes.string,
  tip: PropTypes.string
}

class PremisTextAreaInput extends React.Component {

  /* constructor(props) {
    super(props)
  } */

  _handleChange (e) {
    this.props.parentFunc(e.target.value)
  }

  render () {
    return (
      <div className="textAreaInput">
        <Tooltip title={this.props.tip} position="top">
          <p>{this.props.label}
            <textarea
              placeholder={this.props.placeHolder}
              onChange={this._handleChange.bind(this)}
            />
          </p>
        </Tooltip>
      </div>
    )
  }
}

PremisTextAreaInput.propTypes = {
  parentFunc: PropTypes.func,
  placeHolder: PropTypes.string,
  label: PropTypes.string,
  tip: PropTypes.string
}

class PremisSelect extends React.Component {

  constructor(props) {
    super(props)
  }

  _handleChange (value) {
    this.props.parentFunc(value)
  }

  render () {
    return (
      <Tooltip title={this.props.tip} position="top">
        <Select
          placeholder={this.props.placeHolder}
          searchable={this.props.searchable}
          disabled={this.props.disabled}
          clearable={this.props.clearable}
          options={this.props.selections}
          value={this.props.selection}
          onChange={this._handleChange.bind(this)}
        />
      </Tooltip>
    )
  }
}

PremisSelect.propTypes = {
  parentFunc: PropTypes.func,
  selections: PropTypes.array,
  selection: PropTypes.number,
  searchable: PropTypes.bool,
  placeHolder: PropTypes.string,
  tip: PropTypes.string
}

PremisSelect.defaultProps = {
  disabled: false,
  clearable: true,
  searchable: true
}

class PremisSelectPlus extends React.Component {

  constructor(props) {
    super(props)
  }

  _handleChange (value) {
    this.props.parentFunc(value)
  }

  render () {
    return (
      <Tooltip title={this.props.tip} position="top">
        <Select.Creatable
          placeholder={this.props.placeHolder}
          searchable={this.props.searchable}
          disabled={this.props.disabled}
          clearable={this.props.clearable}
          options={this.props.selections}
          value={this.props.selection}
          onChange={this._handleChange.bind(this)}
        />
      </Tooltip>
    )
  }
}

PremisSelectPlus.propTypes = {
  parentFunc: PropTypes.func,
  selections: PropTypes.array,
  // selection: PropTypes.string,
  searchable: PropTypes.bool,
  placeHolder: PropTypes.string,
  tip: PropTypes.string
}

PremisSelectPlus.defaultProps = {
  disabled: false,
  clearable: true,
  searchable: true
}

class FormSubmit extends React.Component {

  constructor(props) {
    super(props);
  }

  _handleSubmit() {
    this.props.parentFunc()
  }

  render() {
    return (
      <div className="section">
        <Tooltip title={this.props.tip} position="right">
          <Button className={rTComponents.buttonPrimary} onClick={this._handleSubmit.bind(this)} raised ripple><MdDone/></Button>
        </Tooltip>
      </div>
    );
  }
}

FormSubmit.propTypes = {
  parentFunc:PropTypes.func,
  label:PropTypes.string,
  tip:PropTypes.string
}

export {AppHeading, PremisHeading, PremisPlainTextOutput, PremisTextOutput, PremisTextInput,
        PremisTextAreaInput, PremisSelect, PremisSelectPlus, FormSubmit}
