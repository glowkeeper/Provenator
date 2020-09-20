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
import PublisherReaderInput from 'react-file-reader-input'

import Grid from '@material-ui/core/Grid'
import RightCircleOutlined from '@ant-design/icons/lib/icons/RightCircleOutlined'
import { Okay, OptionsStyles } from '../../styles'

import { initialise, getData } from '../../store/app/get/actions'
import { addPublisher } from '../../store/app/blockchain'
import { initialise as txInitialise } from '../../store/app/tx/actions'

import { history, getDictEntries, addressToBytes32, bytes32ToAddress } from '../../utils'

import { NumberOptionType, FormHelpers, GeneralError, Transaction, Remote, Local, Misc, Publisher as PublisherConfig } from '../../config'

import {
    ApplicationState,
    AppDispatch,
    Publisher,
    PayloadProps,
    GetProps,
    EntityProps,
    TxData } from '../../store/types'

import { themeStyles } from '../../styles'

const addPublisherSchema = Yup.object().shape({
  publisherName: Yup.string()
    .required(`${GeneralError.required}`),
  publisherEMail: Yup.string()
    .email()
    .required(`${GeneralError.required}`),
  publisherURL: Yup.string()
    .url(`${PublisherConfig.validURL}`),
})

interface PublisherStateProps {
  info: PayloadProps
  data: GetProps
  address: string
}

interface PublisherDispatchProps {
  initialise: () => void
  getData: (url: string) => void
  handleSubmit: (values: Publisher) => void
}

type Props =  PublisherDispatchProps & PublisherStateProps

export const getPublisher = (props: Props) => {

    let isFirstRun = useRef(true)
    const [isLoading, setIsLoading] = useState(false)
    const [isSubmitting, setSubmit] = useState(false)
    const [publisher, setPublisher] = useState({name: "", email: "", url: ""})
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

                const publisherData = props.data.data[0] as EntityProps
                //console.log(publisherData)
                if ( publisherData.entities ) {
                    if ( publisherData.entities.length > 0 ) {
                        const publisher: Publisher = publisherData.entities[0] as Publisher
                        if ( ( publisher.name != publisher.name ) || (publisher.email != publisher.email) || ( publisher.url != publisher.url ) ) {
                            setPublisher(publisher)
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
        <h2>{PublisherConfig.heading}</h2>
        <hr />
        <Formik
          initialValues={ {publisherName: publisher.name, publisherEMail: publisher.email, publisherURL: publisher.url} }
          enableReinitialize={true}
          validationSchema={addPublisherSchema}
          onSubmit={(values: any) => {

            setSubmit(true)
            props.initialise()

            const id = addressToBytes32(props.address)
            //console.log(id, bytes32ToAddress(id))

            const publisherInfo: Publisher = {
                id: id,
                name: values.publisherName,
                email: values.publisherEMail,
                url:  values.publisherURL
            }
            props.handleSubmit(publisherInfo)
          }}
        >
          {(formProps: FormikProps<any>) => (
            <Form>
              <FormControl fullWidth={true}>
                  <Field
                    name='publisherName'
                    label={PublisherConfig.publisherName}
                    component={TextField}
                  />
                  <Field
                    name='publisherEMail'
                    label={PublisherConfig.publisherEMail}
                    component={TextField}
                  />
                  <Field
                    name='publisherURL'
                    label={PublisherConfig.publisherURL}
                    component={TextField}
                  />
                  <Grid container>
                      <Grid item xs={12} sm={3}>
                        <Tooltip title={PublisherConfig.submitTip}>
                          <Okay type='submit' variant="contained" color="primary" disabled={isSubmitting} endIcon={<RightCircleOutlined spin={isSubmitting}/>}>
                            {PublisherConfig.addPublisherButton}
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

const mapStateToProps = (state: ApplicationState): PublisherStateProps => {
  return {
    info: state.tx as PayloadProps,
    data: state.data as GetProps,
    address: state.chainInfo.data.Account
  }
}

const mapDispatchToProps = (dispatch: AppDispatch): PublisherDispatchProps => {
  return {
    initialise: () => dispatch(txInitialise()),
    getData: (url: string) => dispatch(getData(url)),
    handleSubmit: (values: Publisher) => dispatch(addPublisher(values))
  }
}

export const AddPublisher = connect<PublisherStateProps, PublisherDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(getPublisher)
