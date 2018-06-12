import React from 'react'
import ReactDOM from 'react-dom'
import PropTypes from 'prop-types'
import { Button, Tooltip, Input, Card, Select, Radio } from 'antd';
import ReactMarkdown from 'react-markdown'

class IOAppHeading extends React.Component {

  constructor (props) {
    super(props)
  }

  render () {
    return (
      <div>
        <h1>{this.props.heading}</h1>
      </div>
    )
  }
}

IOAppHeading.propTypes = {
  heading: PropTypes.string
}

class IOTagline extends React.Component {

  constructor (props) {
    super(props)
  }

  render () {
    return (
      <div>
        <h1>{this.props.tagLine}</h1>
      </div>
    )
  }
}

IOTagline.propTypes = {
  tagLine: PropTypes.string
}

class IOHeading extends React.Component {

  /* constructor(props) {
    super(props)
  } */

  render () {
    return (
      <div>
        <h2>{this.props.heading}</h2>
      </div>
    )
  }
}

IOHeading.propTypes = {
  heading: PropTypes.string
}

class IORadioOptions extends React.Component {

  constructor(props) {
    super(props)

    this.options = this.props.options

    this.state = {
      option: this.options[0]
    }
  }

  _handleChange (e) {
    this.setState({
      option: e.target.value
    })
    this.props.parentFunc(e.target.value)
  }

  render () {

    const RadioGroup = Radio.Group
    return (
      <div>
        {this.props.label}
        <Tooltip title={this.props.tip}>
          <RadioGroup
            options={this.options}
            onChange={this._handleChange.bind(this)}
            value={this.state.option}
          />
        </Tooltip>
      </div>
    )
  }
}

IORadioOptions.propTypes = {
  parentFunc: PropTypes.func,
  label: PropTypes.string,
  options: PropTypes.array,
  tip: PropTypes.string
}

class IOPlainTextOutput extends React.Component {

  constructor(props) {
    super(props)
  }

  render () {
    return (
      <div>
        <ReactMarkdown escapeHtml={false} source={this.props.text} />
      </div>
    )
  }
}

IOPlainTextOutput.propTypes = {
  text: PropTypes.string
}


class IOTextOutput extends React.Component {

  /* constructor(props) {
    super(props)
  } */

  render () {
    return (
      <div>
        <p>{this.props.label} {this.props.text}</p>
      </div>
    )
  }
}

IOTextOutput.propTypes = {
  label: PropTypes.string,
  text: PropTypes.string
}

class IOTextInput extends React.Component {

  /* constructor(props) {
    super(props)
  } */

  _handleChange (e) {
    this.props.parentFunc(e.target.value)
  }

  render () {
    return (
      <div>
        <Tooltip title={this.props.tip}>
          <Input
            addonBefore={this.props.label}
            defaultValue={this.props.placeHolder}
            onChange={this._handleChange.bind(this)}
          />
        </Tooltip>
      </div>
    )
  }
}

IOTextInput.propTypes = {
  parentFunc: PropTypes.func,
  placeHolder: PropTypes.string,
  numRows: PropTypes.string,
  label: PropTypes.string,
  tip: PropTypes.string
}

class IOTextAreaInput extends React.Component {

  /* constructor(props) {
    super(props)
  } */

  _handleChange (e) {
    this.props.parentFunc(e.target.value)
  }

  render () {

    const { TextArea } = Input

    return (
      <div>
        <p>{this.props.label}
          <TextArea
            rows={this.props.numRows}
            placeholder={this.props.placeHolder}
            onChange={this._handleChange.bind(this)}
          />
        </p>
      </div>
    )
  }
}

IOTextAreaInput.propTypes = {
  parentFunc: PropTypes.func,
  placeHolder: PropTypes.string,
  numRows: PropTypes.number,
  label: PropTypes.string,
  tip: PropTypes.string
}

IOTextAreaInput.defaultProps = {
  numRows: 2
}

class IOSelect extends React.Component {

  constructor(props) {
    super(props)
  }

  _handleChange (value) {
    this.props.parentFunc(value)
  }

  render () {

    const Option = Select.Option
    let children = []
    let size = this.props.selections.length
    for (let i = 0; i < size; i++) {
      children.push(<Option key={this.props.selections[i]}>{this.props.selections[i]}</Option>)
    }

    return (
      <Tooltip title={this.props.tip} position="top">
        <Select placeholder={this.props.label}
                style={{ width: '100%' }}
                onChange={this._handleChange.bind(this)}
                tokenSeparators={[',']}
        >
          {children}
        </Select>
      </Tooltip>
    )
  }
}

IOSelect.propTypes = {
  parentFunc: PropTypes.func,
  selections: PropTypes.array,
  label: PropTypes.string,
  tip: PropTypes.string
}

class IOLogger extends React.Component {

  constructor(props) {
    super(props)
  }

  render () {
    let logs = this.props.log.map((text, index) =>
      <span key={index}>
        {text}
        <br />
      </span>
    )
    return (
      <div>
        <Card>
          <p>{logs}</p>
        </Card>
      </div>
    )
  }
}

IOLogger.propTypes = {
  log: PropTypes.array
}

class IOButtonLoad extends React.Component {

  constructor (props) {
    super(props);
  }

  _handleSubmit () {
    this.props.parentFunc()
  }

  render () {
    return (
      <div>
        <Tooltip title={this.props.tip}>
          <Button icon={this.props.icon} loading={this.props.loading} onClick={this._handleSubmit.bind(this)}>{this.props.label}</Button>
        </Tooltip>
      </div>
    );
  }
}

IOButtonLoad.propTypes = {
  parentFunc: PropTypes.func,
  icon: PropTypes.string,
  label: PropTypes.string,
  loading: PropTypes.bool,
  tip: PropTypes.string
}

IOButtonLoad.defaultProps = {
  loading: false
}

class IOButton extends React.Component {

  constructor (props) {
    super(props);
  }

  _handleSubmit () {
    this.props.parentFunc()
  }

  render () {
    return (
      <div>
        <Tooltip title={this.props.tip}>
          <Button icon={this.props.icon} onClick={this._handleSubmit.bind(this)} >{this.props.label}</Button>
        </Tooltip>
      </div>
    );
  }
}

IOButton.propTypes = {
  parentFunc: PropTypes.func,
  icon: PropTypes.string,
  label: PropTypes.string,
  tip: PropTypes.string
}

export {IOAppHeading, IOTagline, IOHeading, IORadioOptions, IOPlainTextOutput, IOTextOutput, IOTextInput,
        IOTextAreaInput, IOSelect, IOLogger, IOButtonLoad, IOButton}
