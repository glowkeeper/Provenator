import React, { useRef, useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'
import { connect } from 'react-redux'

import { ethers } from 'ethers'

import Markdown from 'react-markdown'

import * as Yup from 'yup'

import Select, { ValueType, components } from 'react-select'
import { Formik, Form, Field, FormikProps, ErrorMessage } from 'formik'
import { FormControl } from '@material-ui/core'
import { TextField } from "material-ui-formik-components"

import Tooltip from '@material-ui/core/Tooltip'
import CopyrightHolderReaderInput from 'react-file-reader-input'

import Grid from '@material-ui/core/Grid'
import RightCircleOutlined from '@ant-design/icons/lib/icons/RightCircleOutlined'
import { Okay, OptionsStyles } from '../../styles'

import { initialise, getData } from '../../store/app/get/actions'
import { addCopyrightHolder } from '../../store/app/blockchain'
import { initialise as txInitialise } from '../../store/app/tx/actions'

import { history, getArtefactHTML, addressToBytes32, getDictEntries } from '../../utils'

import { StringOptionType, FormHelpers, GeneralError, Transaction, Remote, Local, Misc, CopyrightHolder as CopyrightHolderConfig } from '../../config'

import {
    ApplicationState,
    AppDispatch,
    Author,
    PayloadProps,
    GetProps,
    CreativeWorksProps,
    TxData
} from '../../store/types'

import { themeStyles } from '../../styles'

const addCopyrightHolderSchema = Yup.object().shape({
    copyrightHolder: Yup.object({
        value: Yup.string()
          .required(`${GeneralError.required}`)
    })
})

interface CopyrightHolderStateProps {
  data: GetProps
  info: PayloadProps
  address: string
}

interface CopyrightHolderDispatchProps {
  initialise: () => void
  getData: (url: string) => void
  handleSubmit: (values: Author) => void
}

type Props =  CopyrightHolderDispatchProps & CopyrightHolderStateProps

export const getCopyrightHolder = (props: Props) => {

    let isFirstRun = useRef(true)
    const {id} = useParams()
    const [isLoading, setIsLoading] = useState(false)
    const [isSubmitting, setSubmit] = useState(false)
    const [userId, setUserId] = useState("")
    const [copyrightHolderIds, setCopyrightHolderIds] = useState([] as StringOptionType[])
    const [copyrightHolder, setCopyrightHolder] = useState({ label: CopyrightHolderConfig.holderId, value: ""} as StringOptionType)
    const [fileInfo, seFileInfo] = useState([] as any[])
    const [txInfo, setTxInfo] = useState("")

    const themeClasses = themeStyles()

    useEffect(() => {

        if ( isFirstRun.current ) {

            isFirstRun.current = false
            props.initialise()

            const artefact = getArtefactHTML(props.data.data[0] as CreativeWorksProps, id)
            seFileInfo(artefact)

            const userId = addressToBytes32(props.address)
            setUserId(userId)

            props.getData(`${Remote.insecure}${Remote.server}:${Remote.port}${Remote.entityRelations}/${userId}`)

        } else {

            const txData: TxData = props.info.data as TxData
            const txInfoData = getDictEntries(props.info)
            setTxInfo( txInfoData )

            if( txData.summary == Transaction.success || txData.summary == Transaction.failure ) {
                setSubmit(false)
                setTimeout(() => {
                    history.push(`${Local.home}`)
                }, Misc.delay)
            }
        }

    }, [props.info, props.data])

    return (
      <>
        <h2>{CopyrightHolderConfig.heading}</h2>
        <hr />
            {fileInfo}
        <hr />
        <Formik
          initialValues={ {} }
          enableReinitialize={true}
          validationSchema={addCopyrightHolderSchema}
          onSubmit={(values: any) => {

            setSubmit(true)
            props.initialise()

            const copyrightHolderInfo: Author = {
                id: id,
                name: values.copyrightHolderName,
                email: values.copyrightHolderEMail,
                url:  values.copyrightHolderURL
            }
            props.handleSubmit(copyrightHolderInfo)
          }}
        >
          {(formProps: FormikProps<any>) => (
            <Form>
              <FormControl fullWidth={true}>
                  <Select
                    name='copyrightHolder'
                    value={copyrightHolder}
                    onChange={(ev: ValueType<StringOptionType>) => {
                        const option = ev as StringOptionType
                        formProps.setFieldValue('copyrightHolder', option)
                        setCopyrightHolder(option)
                    }}
                    styles={OptionsStyles}
                    options={copyrightHolderIds}
                  />
                  <ErrorMessage name="subbieRef.value" component="div" className={themeClasses.error}/>
                  <Grid container>
                      <Grid item xs={12} sm={3}>
                        <Tooltip title={CopyrightHolderConfig.submitTip}>
                          <Okay type='submit' variant="contained" color="primary" disabled={isSubmitting} endIcon={<RightCircleOutlined spin={isSubmitting}/>}>
                            {CopyrightHolderConfig.addCopyrightHolderButton}
                          </Okay>
                        </Tooltip>
                      </Grid>
                      <Grid item xs={12} sm={9}>
                          &nbsp;
                      </Grid>
                  </Grid>
              </FormControl>
            </Form>
        )}
        </Formik>
        <hr />
        <Markdown escapeHtml={false} source={txInfo} />
      </>
    )
}

const mapStateToProps = (state: ApplicationState): CopyrightHolderStateProps => {
  return {
    data: state.data as GetProps,
    info: state.tx as PayloadProps,
    address: state.chainInfo.data.Account
  }
}

const mapDispatchToProps = (dispatch: AppDispatch): CopyrightHolderDispatchProps => {
  return {
    initialise: () => dispatch(txInitialise()),
    getData: (url: string) => dispatch(getData(url)),
    handleSubmit: (values: Author) => dispatch(addAuthor(values))
  }
}

export const AddCopyrightHolder = connect<CopyrightHolderStateProps, CopyrightHolderDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(getCopyrightHolder)
