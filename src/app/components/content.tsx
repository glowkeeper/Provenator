import React from 'react'
import { Switch, Route } from 'react-router-dom'

import { InfoTypes } from '../store/types'

import { BlockchainInfo, Info, Home, ListMyArtefacts, AddToArtefacts } from '../components/pages'
import { AddFile, AddUser, AddAuthor, AddCopyrightHolder, AddPublisher } from '../containers/pages/'

import { Paths, Local } from '../config'

export const Content = () => {

    return (

      <Switch>

        <Route name={Paths.help} exact path={Local.help} render={() => <Info type={InfoTypes.HELP}/>} />
        <Route name={Paths.contact} exact path={Local.contact} render={() => <Info type={InfoTypes.CONTACT}/>} />
        <Route name={Paths.about} exact path={Local.about} render={() => <Info type={InfoTypes.ABOUT}/>} />
        <Route name={Paths.blockchain} exact path={Local.blockchain} render= {() => <BlockchainInfo />} />

        <Route name={Paths.listMyArtefacts} path={Local.listMyArtefacts} render={() => <ListMyArtefacts />} />

        <Route name={Paths.addUser} exact path={Local.addUser} render= {() => <AddUser />} />
        <Route name={Paths.addAuthor} exact path={Local.addAuthor} render= {() => <AddAuthor />} />
        <Route name={Paths.addCopyrightHolder} exact path={Local.addCopyrightHolder} render= {() => <AddCopyrightHolder />} />
        <Route name={Paths.addPublisher} exact path={Local.addPublisher} render= {() => <AddPublisher />} />

        <Route name={Paths.addFile} exact path={Local.addFile} render= {() => <AddFile />} />

        <Route name={Paths.addToArtefacts} path={Local.addToArtefacts} render={() => <AddToArtefacts />} />

        <Route name={Paths.home} path={Local.home} render= {() => <Home />} />

      </Switch>
    )
}
