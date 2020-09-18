import React, { useRef, useState, useEffect } from 'react'
import { connect } from 'react-redux'

import { ethers } from 'ethers'

import Markdown from 'react-markdown'

import SparkMD5 from 'spark-md5'

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

import { history, getDictEntries, addressToBytes32, bytes32ToAddress } from '../../utils'

import { NumberOptionType, FormHelpers, GeneralError, Transaction, Remote, Local, Misc, CopyrightHolder as CopyrightHolderConfig } from '../../config'

import {
    ApplicationState,
    AppDispatch,
    CopyrightHolder,
    PayloadProps,
    GetProps,
    EntityProps,
    TxData } from '../../store/types'

import { themeStyles } from '../../styles'

const addCopyrightHolderSchema = Yup.object().shape({
  copyrightHolderName: Yup.string()
    .required(`${GeneralError.required}`),
  copyrightHolderEMail: Yup.string()
    .email()
    .required(`${GeneralError.required}`),
  copyrightHolderURL: Yup.string()
    .url(`${CopyrightHolderConfig.validURL}`),
})

interface CopyrightHolderStateProps {
  info: PayloadProps
  data: GetProps
  address: string
}

interface CopyrightHolderDispatchProps {
  initialise: () => void
  getData: (url: string) => void
  handleSubmit: (values: CopyrightHolder) => void
}

type Props =  CopyrightHolderDispatchProps & CopyrightHolderStateProps

export const getCopyrightHolder = (props: Props) => {

    let isFirstRun = useRef(true)
    const [isLoading, setIsLoading] = useState(false)
    const [isSubmitting, setSubmit] = useState(false)
    const [copyrightHolder, setCopyrightHolder] = useState({name: "", email: "", url: ""})
    const [info, setInfo] = useState("")

    const themeClasses = themeStyles()

    useEffect(() => {

        if ( isFirstRun.current ) {

            isFirstRun.current = false
            props.initialise()

            const id = addressToBytes32(props.address)
            props.getData(`${Remote.insecure}${Remote.server}:${Remote.port}${Remote.entities}/${id}`)

        } else {

            if ( props.data.data.length > 0 ) {

                const copyrightHolderData = props.data.data[0] as EntityProps
                //console.log(copyrightHolderData)
                if ( copyrightHolderData.entities ) {
                    if ( copyrightHolderData.entities.length > 0 ) {
                        const copyrightHolder: CopyrightHolder = copyrightHolderData.entities[0] as CopyrightHolder
                        if ( ( copyrightHolder.name != copyrightHolder.name ) || (copyrightHolder.email != copyrightHolder.email) || ( copyrightHolder.url != copyrightHolder.url ) ) {
                            setCopyrightHolder(copyrightHolder)
                        }
                    }
                }
            }

            const txData: TxData = props.info.data as TxData
            const infoData = getDictEntries(props.info)
            setInfo( infoData )

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
        <Formik
          initialValues={ {copyrightHolderName: copyrightHolder.name, copyrightHolderEMail: copyrightHolder.email, copyrightHolderURL: copyrightHolder.url} }
          enableReinitialize={true}
          validationSchema={addCopyrightHolderSchema}
          onSubmit={(values: any) => {

            setSubmit(true)
            props.initialise()

            const id = addressToBytes32(props.address)
            //console.log(id, bytes32ToAddress(id))

            const copyrightHolderInfo: CopyrightHolder = {
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
                  <Field
                    name='copyrightHolderName'
                    label={CopyrightHolderConfig.copyrightHolderName}
                    component={TextField}
                  />
                  <Field
                    name='copyrightHolderEMail'
                    label={CopyrightHolderConfig.copyrightHolderEMail}
                    component={TextField}
                  />
                  <Field
                    name='copyrightHolderURL'
                    label={CopyrightHolderConfig.copyrightHolderURL}
                    component={TextField}
                  />
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
        <Markdown escapeHtml={false} source={info} />
      </>
    )
}

const mapStateToProps = (state: ApplicationState): CopyrightHolderStateProps => {
  return {
    info: state.tx as PayloadProps,
    data: state.data as GetProps,
    address: state.chainInfo.data.Account
  }
}

const mapDispatchToProps = (dispatch: AppDispatch): CopyrightHolderDispatchProps => {
  return {
    initialise: () => dispatch(txInitialise()),
    getData: (url: string) => dispatch(getData(url)),
    handleSubmit: (values: CopyrightHolder) => dispatch(addCopyrightHolder(values))
  }
}

export const AddCopyrightHolder = connect<CopyrightHolderStateProps, CopyrightHolderDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(getCopyrightHolder)
