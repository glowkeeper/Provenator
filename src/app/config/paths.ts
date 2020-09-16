class Local {

  static readonly home = '/'
  static readonly addUser = '/setUser'
  static readonly addFile = '/addFile'
  static readonly blockchain = '/blockchain'
  static readonly about = '/about'
  static readonly help = '/help'
  static readonly faq = '/faq'
  static readonly blog = '/blog'
  static readonly contact = '/contact'

  static readonly file = '/file'

  static readonly listMyArtefacts = '/myArtefacts'
  static readonly addToArtefacts = '/addToArtefacts'
  
  static readonly addAuthorToArtefact = '/addAuthor'
  static readonly addAuthorToArtefactId = `${Local.addAuthorToArtefact}/:Id`
  static readonly addCopyrightHolderToArtefact = '/addCopyrightHolder'
  static readonly addCopyrightHolderToArtefactId = `${Local.addCopyrightHolderToArtefact}/:Id`
  static readonly addPublisherToArtefact = '/addPublisher'
  static readonly addCPublisherToArtefactId = `${Local.addPublisherToArtefact}/:Id`


}

class Remote {

    static readonly secure = 'https://'
    static readonly insecure = 'http://'

    static readonly server = 'localhost'
    static readonly port = '10000'

    static readonly entities = "/entities"
    static readonly artefacts = "/artefacts"
    static readonly artefactsEntity = "/artefacts-entity"
}

export { Local, Remote }
