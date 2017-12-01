class AppStrings {

  static heading = 'Provenator (alpha)'
  static home = 'Home'
  static about = 'About'
  static overview = 'Overview'
  static help = 'Help'
  static create = 'Record Object'
  static read = 'Get Object'
  static events = 'Events'
}

class HomeStrings {

  static heading = 'Home'

  static info = 'Use this application to record digital media objects and to get digital media object information.'

}

class AboutStrings {

  static heading = 'About Provenator'

  static info = 'This app came about as a result of an academic paper titled: \"Fake News - a Technological Approach to Proving the Origins of Content, Using Blockchains\". The article discusses how the trust mechanisms of blockchain technology might be used to combat fake news. The general idea is that content creators must have the ability to store relevant authentication information about their creations so that the data can be retrieved and used to prove the ownership and licensing of those same resources.'

}

class OverviewStrings {

  static heading = 'Overview of Provenator'

  static info = 'The \'store\' behaviour of the app\' allows content generators to specify information about their digital creations using the PREMIS format, which describes a resource using four relations, 1) Objects, 2) Events, 3) Agents, and 4) Rights. Suppose Alice, a professional landscape photographer, takes a vivid black and white still of an angry storm gathering over the English Channel (it is about to rage over the seaside town of Brighton). Later, she intends to send the photo to prospective buyers, so she needs to establish ownership and copyright. To do so, she loads the photograph into Provenator, which immediately creates a cryptographic hash of the picture _object_ (due to the deterministic and collision resistance properties of cryptographic hashes, that hash will be unique to that image). Alice then adds a description (which should further help identify her picture). She also records an event (the date of the taking of the photo), and lists herself as the agent. Finally, Alice creates a right - the image\'s license. Having established ownership rights of the photograph, Alice posts the picture online. Sometime later, a nature magazine is running a story about climate change, and they intend to use a dramatic weather-based photograph on their front cover.  They find Alice\'s image and wish to identify the owner to discuss the fee for its acquisition.  Hence, they load the image into the _read_ section of Provenator, which matches the cryptographic hash of the _object_ stored by Alice earlier. Consequently, the magazine identifies Alice as the owner of the photograph and arranges with her a fee for its use.'

}

class HelpStrings {

  static heading = 'Help'

  static info = 'To store a record about a digital resource, click on the \'Create\' link and fill in all fields. To retrieve information, click on the \'Read\' link and upload a digital resource. If no information is retrieved, there are no records for that resource.'

}

class HashStrings {

  static heading = 'Select a File Object for Hashing'
  static fileLabel = 'Filename: '
  static hashLabel = 'Hash: '
  static browseFileLabel = 'Select file:'
  static browseFileButtonLabel = 'Browse'

}

class WriterStrings {

  static info = 'Use this form to record information about a media object.'

}

class ReaderStrings {

  static info = 'Use this form to get information about a media object.'

}

class ObjectReaderStrings {

  static heading = 'Object Information'
  static eventHeading = 'Object Event Information'
  static agentHeading = 'Object Agent Information'
  static rightsHeading = 'Object Rights Information'
  static categoryLabel = 'Category: '
  static formatLabel = 'Format: '
  static numPropertiesLabel = 'No. Properties: '
  static propertiesLabel = 'Properties: '

  static numEventsLabel = 'No. Events: '
  static eventLabel = 'Event ID: '
  static eventObjectLabel = 'Event Object: '
  static eventTypeLabel = 'Event Type: '
  static eventAgentLabel = 'Event Agent: '
  static eventTimeLabel = 'Event Time: '

  static numAgentsLabel = 'No. Agents: '
  static agentLabel = 'Agent ID: '
  static agentNameLabel = 'Agent Name: '
  static agentTypeLabel = 'Agent Type: '

  static numRightsLabel = 'No. Rights: '
  static rightsLabel = 'Rights ID: '
  static rightsObjectLabel = 'Rights Object: '
  static rightsBasisLabel = 'Rights Basis: '
  static rightsCopyrightStatusLabel = 'Rights Copyright Status: '
  static rightsCopyrightJurisdictionLabel = 'Rights Copyright Jurisdiction: '
  static rightsCopyrightDeterminationDateLabel = 'Rights Copyright Determination Date: '
  static rightsCopyrightNoteLabel = 'Rights Copyright Note: '
  static rightsGrantedActLabel = 'Rights Granted Act: '
  static rightsGrantedRestrictionLabel = 'Rights Granted Restriction: '
  static rightsAgentLabel = 'Rights Agent: '

  static finishedFetching = 'Finished fetching records from the blockchain.'
  static fetching = 'Fetching file records from the blockchain.'

}

class ObjectWriterStrings {

  static completeAll = 'Please ensure all fields have been completed.'
  static writing = 'Writing records to the blockchain.'
  static finishedWriting = 'Finished writing records to the blockchain.'
  static submitting = 'Submitting records to the blockchain.'

  static submitObjectLabel = 'Submit Object Information:'
  static submitObjectButtonLabel = 'Submit'
}

class PremisRightsStrings {

  static heading = 'Object Rights Information'
  static basisPlaceHolder = 'Basis'
  static basisLabel = 'Basis: '
  static statusPlaceHolder = 'Status'
  static statusLabel = 'Status: '
  static countryCodePlaceHolder = 'Jurisdiction Country Code'
  static countryCodeLabel = 'Jurisdiction Country Code: '
  static determinationDatePlaceHolder = 'Determination Date'
  static determinationDateLabel = 'Determination Date: '
  static notePlaceHolder = 'Note'
  static noteLabel = 'Note: '
  static actPlaceHolder = 'Rights Granted Act'
  static actLabel = 'Rights Granted Act: '
  static restrictionPlaceHolder = 'Rights Granted Restriction'
  static restrictionLabel = 'Rights Granted Restriction: '

}

class PremisObjectStrings {

  static heading = 'Object Information'
  static categoryPlaceHolder = 'Category'
  static categoryLabel = 'Category: '
  static formatPlaceHolder = 'Format'
  static formatLabel = 'Format: '
  static propertyTypePlaceHolder = 'Property Type'
  static propertyTypeLabel = 'Property Type: '
  static propertyValuePlaceHolder = 'Property Value'
  static propertyValueLabel = 'Property Value: '
}

class PremisAgentStrings {

  static heading = 'Object Agent Information'
  static namePlaceHolder = 'Name'
  static nameLabel = 'Name: '
  static typePlaceHolder = 'Type'
  static typeLabel = 'Type: '

}

class PremisEventStrings {

  static heading = 'Object Event Information'
  static info = 'Blockchain write events will appear here.'
  static eventPlaceHolder = 'Event Type'
  static nameLabel = 'Event Type: '
  static datePlaceHolder = 'Event Date'
  static dateLabel = 'Event Date: '

}

class FakeNewsEventStrings {

  static heading = 'Provenator News Events'
  static info = 'Blockchain write events will appear here.'

}

export {AppStrings, HomeStrings, AboutStrings, OverviewStrings, HelpStrings, HashStrings, WriterStrings, ReaderStrings, ObjectReaderStrings, ObjectWriterStrings, PremisRightsStrings, PremisObjectStrings, PremisAgentStrings, PremisEventStrings, FakeNewsEventStrings}
