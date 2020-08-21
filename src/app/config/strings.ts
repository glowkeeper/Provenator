class App {

  static readonly appName = 'Provenator'
  static readonly title = 'Provenator'
  static readonly tagline = 'Prove the provenance of your digital assets'
  static readonly author = '[Dr Steve Huckle](https://glowkeeper.github.io/)'
  static readonly copyright = `© Copyright 2020 _${App.author}_`
}

class Paths {

  // AppBar
  static readonly home = 'Home'
  static readonly addFile = 'Hash File'
  static readonly blockchain = 'Blockchain'
  static readonly about = 'About'
  static readonly help = 'Help'
  static readonly faq = 'FAQ'
  static readonly contact = 'Contact'

  static readonly file = 'File'
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

class File {

    static readonly headingFile = "Set the Provenance of a Digital Media File"
    static readonly getFile = "Get the Digital Media File"
    static readonly fileName = "File Name"
    static readonly hash = "File Hash"
    static readonly fileUrl = "File URL"
    static readonly workType = "Works Type"
    static readonly description = 'Description'

    static readonly copyrightHolderName = "Copyright Holder Name"
    static readonly copyrightHolderEMail = "Copyright Holder EMail"
    static readonly copyrightHolderURL = "Copyright Holder  URL"
    static readonly authorName = "Author Name"
    static readonly authorEMail = "Author EMail"
    static readonly authorURL = "Author URL"
    static readonly license = "License"
    static readonly publisherName = "Publisher Name"
    static readonly publisherEMail = "Publisher EMail"
    static readonly publisherURL = "Publisher URL"

    static readonly fileTip = "Select a file to be hashed"
    static readonly submitTip = "Submit the hash of the file to the Ethereum blockchain"

    static readonly loadingError = "File did not load"
    static readonly validURL = "Please enter a valid URL"
    static readonly validType = "Please select a valid creative works type"
    static readonly validLicense = "Please select a valid license (or specify 'none')"

    static readonly addFileButton = "Submit Info"
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
         File
       }
