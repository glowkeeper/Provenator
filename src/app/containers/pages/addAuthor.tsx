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
import AuthorReaderInput from 'react-file-reader-input'

import Grid from '@material-ui/core/Grid'
import RightCircleOutlined from '@ant-design/icons/lib/icons/RightCircleOutlined'
import { Okay, OptionsStyles } from '../../styles'

import { initialise, getData } from '../../store/app/get/actions'
import { addAuthor } from '../../store/app/blockchain'
import { initialise as txInitialise } from '../../store/app/tx/actions'

import { history, getDictEntries, md5ToBytes32, addressToBytes32, bytes32ToAddress } from '../../utils'

import { NumberOptionType, FormHelpers, GeneralError, Transaction, Remote, Local, Misc, Author as AuthorConfig } from '../../config'

import {
    ApplicationState,
    AppDispatch,
    Author,
    PayloadProps,
    GetProps,
    EntityProps,
    TxData } from '../../store/types'

import { themeStyles } from '../../styles'

const addAuthorSchema = Yup.object().shape({
  authorName: Yup.string()
    .required(`${GeneralError.required}`),
  authorEMail: Yup.string()
    .email()
    .required(`${GeneralError.required}`),
  authorURL: Yup.string()
    .url(`${AuthorConfig.validURL}`),
})

interface AuthorStateProps {
  info: PayloadProps
}

interface AuthorDispatchProps {
  initialise: () => void
  handleSubmit: (values: Author) => void
}

type Props =  AuthorDispatchProps & AuthorStateProps

export const getAuthor = (props: Props) => {

    let isFirstRun = useRef(true)
    const [isLoading, setIsLoading] = useState(false)
    const [isSubmitting, setSubmit] = useState(false)
    const [author, setAuthor] = useState({name: "", email: "", url: ""})
    const [info, setInfo] = useState("")

    const themeClasses = themeStyles()

    useEffect(() => {

        if ( isFirstRun.current ) {

            isFirstRun.current = false
            props.initialise()

        } else {

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

    }, [props.info])

    return (
      <>
        <h2>{AuthorConfig.headingAuthor}</h2>
        <hr />
        <Formik
          initialValues={ {authorName: author.name, authorEMail: author.email, authorURL: author.url} }
          enableReinitialize={true}
          validationSchema={addAuthorSchema}
          onSubmit={(values: any) => {

            setSubmit(true)
            props.initialise()

            const hash = SparkMD5.hash(values.authorName + values.authorEMail);
            const id = md5ToBytes32(hash)
            const authorInfo: Author = {
                id: id,
                name: values.authorName,
                email: values.authorEMail,
                url:  values.authorURL
            }
            props.handleSubmit(authorInfo)
          }}
        >
          {(formProps: FormikProps<any>) => (
            <Form>
              <FormControl fullWidth={true}>
                  <Field
                    name='authorName'
                    label={AuthorConfig.authorName}
                    component={TextField}
                  />
                  <Field
                    name='authorEMail'
                    label={AuthorConfig.authorEMail}
                    component={TextField}
                  />
                  <Field
                    name='authorURL'
                    label={AuthorConfig.authorURL}
                    component={TextField}
                  />
                  <Grid container>
                      <Grid item xs={12} sm={3}>
                        <Tooltip title={AuthorConfig.submitTip}>
                          <Okay type='submit' variant="contained" color="primary" disabled={isSubmitting} endIcon={<RightCircleOutlined spin={isSubmitting}/>}>
                            {AuthorConfig.addAuthorButton}
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

const mapStateToProps = (state: ApplicationState): AuthorStateProps => {
  return {
    info: state.tx as PayloadProps
  }
}

const mapDispatchToProps = (dispatch: AppDispatch): AuthorDispatchProps => {
  return {
    initialise: () => dispatch(txInitialise()),
    handleSubmit: (values: Author) => dispatch(addAuthor(values))
  }
}

export const AddAuthor = connect<AuthorStateProps, AuthorDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(getAuthor)
