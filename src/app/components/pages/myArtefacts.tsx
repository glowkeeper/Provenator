import React, { useRef, useState, useEffect } from 'react'
import { connect } from 'react-redux'

import Spinner from 'react-spinner-material'
import { SimpleArrayRenderer } from '../simpleRenderer'

import { getArtefactsHTML } from '../../utils/artefacts'

import { addressToBytes32 } from '../../utils'

import { initialise, getData } from '../../store/app/get/actions'

import {
    ApplicationState,
    AppDispatch,
    GetProps,
    CreativeWorksProps
} from '../../store/types'

import { FormHelpers, Artefacts, GeneralError, Remote } from '../../config'

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

const artefactsReader = (props: Props) => {

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

        } else if (typeof props.artefacts.data !== 'undefined') {
            if ( props.artefacts.data.length > 0 ) {

                const artefactsInfo = getArtefactsHTML(props.artefacts.data[0] as CreativeWorksProps)
                setArtefactsInfo(artefactsInfo)
                setLoading(false)
            }
        }

    }, [props.artefacts])

    return (
      <div>
        <h2>{Artefacts.headingFileList}</h2>
        <hr />
        <div>
            {isLoading ? <div className={themeClasses.spinner}><Spinner radius={40} color={"#38348B"} stroke={5} visible={isLoading} /></div> : <SimpleArrayRenderer data={artefactsInfo} /> }
        </div>
      </div>
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

export const ListMyArtefacts = connect<ArtefactsProps, ArtefactsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(artefactsReader)
