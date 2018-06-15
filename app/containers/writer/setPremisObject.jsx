import React from 'react'
import PropTypes from 'prop-types'
import PremisObjectHandler from '../../utils/premisObjectHandler'
import {PremisObjectStrings, ObjectWriterStrings} from '../../helpers/outputStrings'
import {IOHeading, IOTextInput, IOTextAreaInput, IOSelect} from '../../components/io'

class SetPremisObject extends React.Component {

  constructor (props) {
    super(props)

    this.state = {
      categoryId: undefined,
      categorySelections: PremisObjectHandler.categoryOptions,
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
          selections={this.state.categorySelections}
          label={ObjectWriterStrings.categoryLabel}
          tip={PremisObjectStrings.categoryTip}
        />
        <IOTextInput parentFunc={this._handleFormat.bind(this)} placeHolder={PremisObjectStrings.formatPlaceHolder} label={PremisObjectStrings.formatLabel} tip={PremisObjectStrings.formatTip} />
        <IOSelect parentFunc={this._handlePropertyType.bind(this)} placeHolder={PremisObjectStrings.propertyTypePlaceHolder} tip={PremisObjectStrings.propertyTypeTip} selections={this.state.propertyTypeSelections} selection={this.state.propertyTypeId} />
        <IOTextAreaInput parentFunc={this._handlePropertyValue.bind(this)} placeHolder={PremisObjectStrings.propertyValuePlaceHolder} label={PremisObjectStrings.propertyValueLabel} tip={PremisObjectStrings.propertyValueTip} />
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
