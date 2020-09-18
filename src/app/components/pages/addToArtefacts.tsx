import React, { useRef, useState, useEffect } from 'react'
import { NavLink } from 'react-router-dom'
import { connect } from 'react-redux'

import Spinner from 'react-spinner-material'
import { SimpleArrayRenderer } from '../simpleRenderer'

import { addressToBytes32 } from '../../utils'

import { history } from '../../utils'

import { Okay, OptionsStyles } from '../../styles'

import { initialise, getData } from '../../store/app/get/actions'

import {
    ApplicationState,
    AppDispatch,
    GetProps,
    CreativeWorksProps
} from '../../store/types'

import { FormHelpers, Artefacts, File, GeneralError, Remote, Paths, Local } from '../../config'

import { themeStyles } from '../../styles'

interface ArtefactsProps {
  artefacts: GetProps
  address: string
}

interface ArtefactsDispatchProps {
  initialise: () => void
  getData: (url: string) => void
}

type Props =  ArtefactsProps & ArtefactsDispatchProps

const artefacts = (props: Props) => {

    const [isLoading, setLoading] = useState(true)
    let isFirstRun = useRef(true)
    let [artefactsInfo, setArtefactsInfo] = useState([] as any[])

    const themeClasses = themeStyles()

    useEffect(() => {

        if ( isFirstRun.current ) {

            isFirstRun.current = false
            props.initialise()

            const id = addressToBytes32(props.address)
            props.getData(`${Remote.insecure}${Remote.server}:${Remote.port}${Remote.artefactsEntity}/${id}`)

        } else if ( props.artefacts.data ) {
            if (props.artefacts.data.length > 0 ) {

                let artefactsInfo: any[] = []

                const data: CreativeWorksProps = props.artefacts.data[0] as CreativeWorksProps
                if( data.creativeWorks ) {
                    for ( var i = 0; i < data.creativeWorks.length; i++) {

                        const id = data.creativeWorks[i].id
                        const pathAddAuthor = `${Local.addAuthorToArtefact}/${id}`
                        const pathAddCopyrightHolder = `${Local.addCopyrightHolderToArtefact}/${id}`
                        const pathAddPublisher = `${Local.addPublisherToArtefact}/${id}`

                        const renderArtefactsHTML = (
                            <React.Fragment key={data.creativeWorks[i].id}>
                            <p>
                                {Artefacts.id}: {id}<br/>
                                {Artefacts.artefactName}: {data.creativeWorks[i].name}<br/>
                                {Artefacts.url}: {data.creativeWorks[i].url}<br/>
                                {Artefacts.description}: {data.creativeWorks[i].description}<br/>
                                {Artefacts.dateCreated}: {data.creativeWorks[i].dateCreated}<br/>
                                {Artefacts.dateModified}: {data.creativeWorks[i].dateModified}<br/>
                                {Artefacts.workType}: {FormHelpers.works[data.creativeWorks[i].workType]}<br/>
                                {Artefacts.license}: {FormHelpers.licenses[data.creativeWorks[i].license]}
                            </p>
                            </React.Fragment>
                        )
                        artefactsInfo.push(renderArtefactsHTML)

                        const buttonsHTML = (
                            <React.Fragment key={id}>
                            <p>
                                <Okay variant="contained" color="primary" onClick={() => history.push(`${pathAddAuthor}`)}>
                                  {Artefacts.addAuthorButton}
                                </Okay>
                                &nbsp;
                                <Okay variant="contained" color="primary" onClick={() => history.push(`${pathAddCopyrightHolder}`)}>
                                  {Artefacts.addCopyrightHolderButton}
                                </Okay>
                                &nbsp;
                                <Okay variant="contained" color="primary" onClick={() => history.push(`${pathAddPublisher}`)}>
                                  {Artefacts.addPublisherButton}
                                </Okay>
                            </p>
                            </React.Fragment>
                        )
                        artefactsInfo.push(buttonsHTML)
                    }
                }

                setArtefactsInfo(artefactsInfo)
                setLoading(false)
            }
        }
    }, [props.artefacts])

    return (
      <>
        <h2>{Artefacts.headingAddToArtefact}</h2>
        <hr />
        <div>
            {isLoading ? <div className={themeClasses.spinner}><Spinner radius={40} color={"#38348B"} stroke={5} visible={isLoading} /></div> : <SimpleArrayRenderer data={artefactsInfo} /> }
        </div>
      </>
    )
}

const mapStateToProps = (state: ApplicationState): ArtefactsProps => {
  //console.log(state.orgReader)
  return {
    artefacts: state.data as GetProps,
    address: state.chainInfo.data.Account
  }
}

const mapDispatchToProps = (dispatch: AppDispatch): ArtefactsDispatchProps => {
  return {
    initialise: () => dispatch(initialise()),
    getData: (url: string) => dispatch(getData(url))
  }
}

export const AddToArtefacts = connect<ArtefactsProps, ArtefactsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(artefacts)
