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
import UserReaderInput from 'react-file-reader-input'

import Grid from '@material-ui/core/Grid'
import RightCircleOutlined from '@ant-design/icons/lib/icons/RightCircleOutlined'
import { Okay, OptionsStyles } from '../../styles'

import { addAuthor } from '../../store/app/blockchain'
import { initialise as txInitialise } from '../../store/app/tx/actions'

import { history, getDictEntries } from '../../utils'

import { NumberOptionType, FormHelpers, GeneralError, Transaction, Local, Misc, User as UserConfig } from '../../config'

import {
    ApplicationState,
    AppDispatch,
    Author,
    PayloadProps,
    TxData } from '../../store/types'

import { themeStyles } from '../../styles'

const addUserSchema = Yup.object().shape({
  authorName: Yup.string()
    .required(`${GeneralError.required}`),
  authorEMail: Yup.string()
    .email()
    .required(`${GeneralError.required}`),
  authorURL: Yup.string()
    .url(`${UserConfig.validURL}`),
})

interface UserStateProps {
  info: PayloadProps
  address: string
}

interface UserDispatchProps {
  initialise: () => void
  handleSubmit: (values: Author) => void
}

type Props =  UserDispatchProps & UserStateProps

export const getUser = (props: Props) => {

    let isFirstRun = useRef(true)
    const [isLoading, setIsLoading] = useState(false)
    const [isSubmitting, setSubmit] = useState(false)
    const [info, setInfo] = useState("")

    const themeClasses = themeStyles()

    useEffect(() => {

        if ( isFirstRun.current ) {

            isFirstRun.current = false

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
        <h2>{UserConfig.headingUser}</h2>
        <hr />
        <Formik
          initialValues={ {} }
          enableReinitialize={true}
          validationSchema={addUserSchema}
          onSubmit={(values: any) => {

            setSubmit(true)
            props.initialise()

            // 0x79b0e7De13a17a74AB23Fd2e6c69AA3Cf93F4E1c
            const id = props.address + "000000000000000000000000"

            const userInfo: Author = {
                id: id,
                name: values.authorName,
                email: values.authorEMail,
                url:  values.authorURL
            }
            props.handleSubmit(userInfo)
          }}
        >
          {(formProps: FormikProps<any>) => (
            <Form>
              <FormControl fullWidth={true}>
                  <Field
                    name='authorName'
                    label={UserConfig.authorName}
                    component={TextField}
                  />
                  <Field
                    name='authorEMail'
                    label={UserConfig.authorEMail}
                    component={TextField}
                  />
                  <Field
                    name='authorURL'
                    label={UserConfig.authorURL}
                    component={TextField}
                  />
                  <Grid container>
                      <Grid item xs={12} sm={3}>
                        <Tooltip title={UserConfig.submitTip}>
                          <Okay type='submit' variant="contained" color="primary" disabled={isSubmitting} endIcon={<RightCircleOutlined spin={isSubmitting}/>}>
                            {UserConfig.addUserButton}
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

const mapStateToProps = (state: ApplicationState): UserStateProps => {
  return {
    info: state.tx as PayloadProps,
    address: state.chainInfo.data.Account
  }
}

const mapDispatchToProps = (dispatch: AppDispatch): UserDispatchProps => {
  return {
    initialise: () => dispatch(txInitialise()),
    handleSubmit: (values: Author) => dispatch(addAuthor(values))
  }
}

export const AddUser = connect<UserStateProps, UserDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(getUser)
