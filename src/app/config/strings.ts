class App {

  static readonly appName = 'Provenator'
  static readonly title = 'Provenator'
  static readonly tagline = 'Prove the provenance of your digital assets'
  static readonly author = '[Dr Steve Huckle](https://glowkeeper.github.io/)'
  static readonly copyright = `© Copyright 2020 _${App.author}_`
}

class Paths {

  static readonly home = 'Home'
  static readonly addUser = 'Add User'
  static readonly addAuthor = 'Add Author'
  static readonly addCopyrightHolder = 'Add Copyright Holder'
  static readonly addPublisher = 'Add Publisher'
  static readonly addFile = 'Hash File'
  static readonly blockchain = 'Blockchain'
  static readonly about = 'About'
  static readonly help = 'Help'
  static readonly faq = 'FAQ'
  static readonly contact = 'Contact'

  static readonly file = 'File'
  static readonly listMyArtefacts = 'My Files'
  static readonly addToArtefacts = 'Add Info to Files'

  static readonly addAuthorToArtefact = 'Add an Author to a File'
  static readonly addCopyrightHolderToArtefact = 'Add a Copyright Holder to a File'
  static readonly addPublisherToArtefact = 'Add a Publisherto a File'
}


class Blockchain {

  static heading = 'Blockchain Data'
}

class GeneralError {

    static readonly required = "Required"
}

class Transaction {

    static readonly pending = "Please wait - transaction is pending"
    static readonly success = "Submitted successfully"
    static readonly failure = 'Submission failure'

    static readonly errorGettingData = "Error getting data"
}

class Home {

  static readonly heading = 'Home'

  static readonly info = `**${App.appName}**`
}

class About {

  static readonly heading = 'About'

  static readonly info = `**${App.appName}** version 0.1.0.<br /><br />Created by _${App.author}_`

}

class Help {

  static readonly heading = 'Help'

  static readonly info = `Use **${App.appName}** to help prove the provenance of your digital assets. Select _${Paths.addFile}_ from the menu to add an asset **now**.`
}

class Faq {

  static readonly heading = 'FAQ'

  static readonly info = `Coming soon`
}

class Contact {

  static readonly heading = 'Contact'

  static readonly info = `steve at glowkeeper dot uk`
}

class User {

    static readonly headingUser = "Set Your Author Details"

    static readonly userName = "User Name"
    static readonly userEMail = "User eMail"
    static readonly userURL = "User URL"

    static readonly validURL = "Please enter a valid URL"

    static readonly submitTip = "Submit your user details to the Ethereum blockchain"

    static readonly addUserButton = "Submit Info"
}

class Author {

    static readonly headingAuthor = "Set Additional Author Details"

    static readonly authorName = "Author Name"
    static readonly authorEMail = "Author eMail"
    static readonly authorURL = "Author URL"

    static readonly validURL = "Please enter a valid URL"

    static readonly submitTip = "Submit your author details to the Ethereum blockchain"

    static readonly addAuthorButton = "Submit Info"
}

class CopyrightHolder {

    static readonly heading = "Set Copyright Holder Details"

    static readonly copyrightHolderName = "Copyright Holder Name"
    static readonly copyrightHolderEMail = "Copyright Holder eMail"
    static readonly copyrightHolderURL = "Copyright Holder URL"

    static readonly holderId = "Copyright Holder ID"

    static readonly validURL = "Please enter a valid URL"

    static readonly submitTip = "Submit copyright holder details to the Ethereum blockchain"

    static readonly addCopyrightHolderButton = "Submit Info"
}

class Publisher {

    static readonly heading = "Set Publisher Details"

    static readonly publisherName = "Publisher Name"
    static readonly publisherEMail = "Publisher eMail"
    static readonly publisherURL = "Publisher URL"

    static readonly validURL = "Please enter a valid URL"

    static readonly submitTip = "Submit publisher details to the Ethereum blockchain"

    static readonly addPublisherButton = "Submit Info"
}

class File {

    static readonly headingFile = "Set the Provenance of a Digital Media File"
    static readonly getFile = "Get the Digital Media File"
    static readonly fileName = "File Name"
    static readonly hash = "File Hash"
    static readonly fileUrl = "File URL"
    static readonly workType = "Works Type"
    static readonly description = 'Description'
    static readonly license = "License"

    static readonly authorName = "Author Name"
    static readonly authorEMail = "Author eMail"
    static readonly authorURL = "Author URL"

    static readonly copyrightHolderName = "Copyright Holder Name"
    static readonly copyrightHolderEMail = "Copyright Holder eMail"
    static readonly copyrightHolderURL = "Copyright Holder  URL"

    static readonly publisherName = "Publisher Name"
    static readonly publisherEMail = "Publisher eMail"
    static readonly publisherURL = "Publisher URL"

    static readonly fileTip = "Select a file to be hashed"
    static readonly submitTip = "Submit the hash of the file to the Ethereum blockchain"

    static readonly loadingError = "File did not load"
    static readonly validURL = "Please enter a valid URL"
    static readonly validType = "Please select a valid creative works type"
    static readonly validLicense = "Please select a valid license (or specify 'none')"

    static readonly addFileButton = "Submit Info"
}

class Artefacts {

    static readonly headingFileList = "My Digital Media"
    static readonly headingAddToArtefact = "Add Info to Files"

    static readonly workType = "Work Type"
    static readonly license = "License"
    static readonly dateCreated = "Date Created"
    static readonly dateModified = "Date Modified"
    static readonly url = "URL"
    static readonly artefactName = "Name"
    static readonly description = "Description"

    static readonly id = "Works ID"

    static readonly authorId = "Author ID"
    static readonly copyrightHolderId = "Copyright Holder ID"
    static readonly publisherId = "Publisher ID"

    static readonly addAuthorButton = "Add Author"
    static readonly addCopyrightHolderButton = "Add Copyright Holder"
    static readonly addPublisherButton = "Add Publisher"
}


export { App,
         Paths,
         Blockchain,
         GeneralError,
         Transaction,
         Home,
         About,
         Help,
         Faq,
         Contact,
         Author,
         User,
         CopyrightHolder,
         Publisher,
         File,
         Artefacts
       }
