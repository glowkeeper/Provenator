class Local {

  static readonly home = '/'
  static readonly addUser = '/set-user'
  static readonly addFile = '/add-file'
  static readonly blockchain = '/blockchain'
  static readonly about = '/about'
  static readonly help = '/help'
  static readonly faq = '/faq'
  static readonly blog = '/blog'
  static readonly contact = '/contact'

  static readonly file = '/file'

  static readonly listMyArtefacts = '/my-artefacts'
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
