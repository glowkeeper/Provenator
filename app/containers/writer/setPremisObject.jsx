import React from 'react'
import PropTypes from 'prop-types'
import PremisObjectHandler from '../../utils/premisObjectHandler'
import {PremisObjectStrings} from '../../helpers/outputStrings'
import {IOHeading, IOTextInput, IOTextAreaInput, IOSelect} from '../../components/io'

class SetPremisObject extends React.Component {

  constructor (props) {
    super(props)

    this.state = {
      propertyTypeId: undefined,
      propertyTypeSelections: []
    }
  }

  setPropTypeName (_self, _name) {
    // console.log(_name)
    let propertyTypeSelections = _self.state.propertyTypeSelections
    propertyTypeSelections.push({ value: propertyTypeSelections.length, label: _name })
    _self.setState({propertyTypeSelections: propertyTypeSelections})
  }

  _handleCategory (_selection) {
    // console.log('Category is ' + category)
    this.props.objectHandler.setCategory(_selection.label)
    this.setState({categoryId: _selection.value})
  }

  _handleFormat (_format) {
    // console.log('format is ' + format)
    this.props.objectHandler.setFormat(_format)
  }

  _handlePropertyType (_selection) {
    // console.log('propertyType is ' + _selection.value)
    this.props.objectHandler.setPropertyType(_selection.label)
    this.setState({propertyTypeId: _selection.value})
  }

  _handlePropertyValue (_propertyValue) {
    this.props.objectHandler.setPropertyValue(_propertyValue)
  }

  render () {

    return (
      <div>
        <IOHeading heading={PremisObjectStrings.heading} />
        <IOSelect
          parentFunc={this._handleCategory.bind(this)}
          selections={PremisObjectHandler.categoryOptions}
          label={PremisObjectStrings.categoryLabel}
          tip={PremisObjectStrings.categoryTip}
        />
        <IOTextInput parentFunc={this._handleFormat.bind(this)} placeHolder={PremisObjectStrings.formatPlaceHolder} label={PremisObjectStrings.formatLabel} tip={PremisObjectStrings.formatTip} />
        <IOTextAreaInput parentFunc={this._handlePropertyValue.bind(this)} label={PremisObjectStrings.propertyDescriptionLabel} tip={PremisObjectStrings.propertyDescriptionTip} />
      </div>
    )
  }

  /* componentWillUpdate() {

  } */

  componentDidMount () {
    // console.log('In object ' + this.props.premisObject)
    this.props.contractWriter.getPropTypes(this, this.setPropTypeName)
  }

}

SetPremisObject.propTypes = {
  contractWriter: PropTypes.object,
  objectHandler: PropTypes.object
}

export default SetPremisObject
