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

import { initialise, getData } from '../../store/app/get/actions'
import { addUser } from '../../store/app/blockchain'
import { initialise as txInitialise } from '../../store/app/tx/actions'

import { history, getDictEntries, addressToBytes32, bytes32ToAddress } from '../../utils'

import { NumberOptionType, FormHelpers, GeneralError, Transaction, Remote, Local, Misc, User as UserConfig } from '../../config'

import {
    ApplicationState,
    AppDispatch,
    User,
    PayloadProps,
    GetProps,
    EntityProps,
    TxData } from '../../store/types'

import { themeStyles } from '../../styles'

const addUserSchema = Yup.object().shape({
  userName: Yup.string()
    .required(`${GeneralError.required}`),
  userEMail: Yup.string()
    .email()
    .required(`${GeneralError.required}`),
  userURL: Yup.string()
    .url(`${UserConfig.validURL}`),
})

interface UserStateProps {
  info: PayloadProps
  user: GetProps
  address: string
}

interface UserDispatchProps {
  initialise: () => void
  getData: (url: string) => void
  handleSubmit: (values: User) => void
}

type Props =  UserDispatchProps & UserStateProps

export const getUser = (props: Props) => {

    let isFirstRun = useRef(true)
    const [isLoading, setIsLoading] = useState(false)
    const [isSubmitting, setSubmit] = useState(false)
    const [user, setUser] = useState({name: "", email: "", url: ""})
    const [info, setInfo] = useState("")

    const themeClasses = themeStyles()

    useEffect(() => {

        if ( isFirstRun.current ) {

            isFirstRun.current = false
            props.initialise()

            const id = addressToBytes32(props.address)
            props.getData(`${Remote.insecure}${Remote.server}:${Remote.port}${Remote.entities}/${id}`)

        } else {

            if ( props.user.data.length > 0 ) {

                const userData = props.user.data[0] as EntityProps
                if ( userData.entities ) {
                    if ( userData.entities.length > 0 ) {
                        const user: User = userData.entities[0] as User
                        if ( ( user.name != user.name ) || (user.email != user.email) || ( user.url != user.url ) ) {
                            setUser(user)
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

    }, [props.info, props.user])

    return (
      <>
        <h2>{UserConfig.headingUser}</h2>
        <hr />
        <Formik
          initialValues={ {userName: user.name, userEMail: user.email, userURL: user.url} }
          enableReinitialize={true}
          validationSchema={addUserSchema}
          onSubmit={(values: any) => {

            setSubmit(true)
            props.initialise()

            const id = addressToBytes32(props.address)

            const userInfo: User = {
                id: id,
                name: values.userName,
                email: values.userEMail,
                url:  values.userURL
            }
            props.handleSubmit(userInfo)
          }}
        >
          {(formProps: FormikProps<any>) => (
            <Form>
              <FormControl fullWidth={true}>
                  <Field
                    name='userName'
                    label={UserConfig.userName}
                    component={TextField}
                  />
                  <Field
                    name='userEMail'
                    label={UserConfig.userEMail}
                    component={TextField}
                  />
                  <Field
                    name='userURL'
                    label={UserConfig.userURL}
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
    user: state.data as GetProps,
    address: state.chainInfo.data.Account
  }
}

const mapDispatchToProps = (dispatch: AppDispatch): UserDispatchProps => {
  return {
    initialise: () => dispatch(txInitialise()),
    getData: (url: string) => dispatch(getData(url)),
    handleSubmit: (values: User) => dispatch(addUser(values))
  }
}

export const AddUser = connect<UserStateProps, UserDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(getUser)
