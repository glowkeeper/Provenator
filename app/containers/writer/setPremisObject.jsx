import React from 'react'
import PropTypes from 'prop-types'
import PremisObjectHandler from '../../utils/writer/premisObjectHandler'
import {ObjectStrings} from '../../helpers/outputStrings'
import {PremisObjectStrings} from '../../helpers/strings'
import {IOHeading, IOTextInput, IOTextAreaInput, IOSelect} from '../../components/io'

class SetPremisObject extends React.Component {

  constructor (props) {
    super(props)

    this.props.objectHandler.setPropertyType(PremisObjectStrings.propertyType)
  }

  _handleCategory (_selection) {
    // console.log('Category is ' + category)
    this.props.objectHandler.setCategory(_selection)
  }

  _handleFormat (_format) {
    //console.log('format is ' + _format)
    this.props.objectHandler.setFormat(_format)
  }

  _handleSetDescription (_description) {
    //console.log(_description)
    this.props.objectHandler.setPropertyValue(_description)
  }

  render () {

    return (
      <div>
        <IOHeading heading={ObjectStrings.heading} />
        <IOSelect
          parentFunc={this._handleFormat.bind(this)}
          selections={ObjectStrings.formats}
          label={ObjectStrings.formatLabel}
          tip={ObjectStrings.formatTip}
        />
        <IOSelect
          parentFunc={this._handleCategory.bind(this)}
          selections={ObjectStrings.categoryOptions}
          label={ObjectStrings.categoryLabel}
          tip={ObjectStrings.categoryTip}
        />
        <IOTextAreaInput
          parentFunc={this._handleSetDescription.bind(this)}
          label={ObjectStrings.propertyDescriptionLabel}
          tip={ObjectStrings.propertyDescriptionTip}
        />
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
