import React from 'react'
import PropTypes from 'prop-types'
import PremisObjectHandler from '../utils/premisObjectHandler'
import {PremisHeading, PremisTextInput, PremisTextAreaInput, PremisSelect, PremisSelectPlus} from '../components/premis'
import {PremisObjectStrings} from '../utils/outputStrings'

class SetPremisObject extends React.Component {

  constructor (props) {
    super(props)

    const numCategories = PremisObjectHandler.categoryOptions.length
    let categorySelections = []
    for (let i = 0; i < numCategories; i++) {
      categorySelections[i] = { value: i, label: PremisObjectHandler.categoryOptions[i] }
    }

    this.state = {
      categoryId: undefined,
      categorySelections: categorySelections,
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
        <PremisHeading heading={PremisObjectStrings.heading} />
        <PremisSelect parentFunc={this._handleCategory.bind(this)} placeHolder={PremisObjectStrings.categoryPlaceHolder} label={PremisObjectStrings.categoryLabel} selections={this.state.categorySelections} selection={this.state.categoryId} />
        <PremisTextInput parentFunc={this._handleFormat.bind(this)} placeHolder={PremisObjectStrings.formatPlaceHolder} label={PremisObjectStrings.formatLabel} />
        <PremisSelectPlus parentFunc={this._handlePropertyType.bind(this)} placeHolder={PremisObjectStrings.propertyTypePlaceHolder} label={PremisObjectStrings.propertyTypeLabel} selections={this.state.propertyTypeSelections} selection={this.state.propertyTypeId} />
        <PremisTextAreaInput parentFunc={this._handlePropertyValue.bind(this)} placeHolder={PremisObjectStrings.propertyValuePlaceHolder} label={PremisObjectStrings.propertyValueLabel} />
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
