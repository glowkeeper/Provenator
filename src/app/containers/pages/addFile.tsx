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
import FileReaderInput from 'react-file-reader-input'

import Grid from '@material-ui/core/Grid'
import RightCircleOutlined from '@ant-design/icons/lib/icons/RightCircleOutlined'
import { Okay, OptionsStyles } from '../../styles'

import { initialise, getData } from '../../store/app/get/actions'
import { addFile } from '../../store/app/blockchain'
import { initialise as txInitialise } from '../../store/app/tx/actions'

import { history, getDictEntries, addressToBytes32, bytes32ToAddress } from '../../utils'

import { NumberOptionType, FormHelpers, GeneralError, Transaction, Remote, Local, Misc, File as FileConfig } from '../../config'

import {
    ApplicationState,
    AppDispatch,
    Author,
    CreativeWorks,
    PayloadProps,
    GetProps,
    TxData } from '../../store/types'

import { themeStyles } from '../../styles'

/*
copyrightHolderName: Yup.string()
  .ensure(),
copyrightHolderEMail: Yup.string()
  .email(),
copyrightHolderURL: Yup.string()
  .url(`${FileConfig.validURL}`),
publisherName: Yup.string()
  .ensure(),
publisherEMail: Yup.string()
  .email(),
publisherURL: Yup.string()
  .url(`${FileConfig.validURL}`),

  {
      id: "",
      name: values.authorName,
      email: values.authorEMail,
      url:  values.authorURL
  },
  copyrightHolder: {
      id: "",
      name: values.copyrightHolderName,
      email: values.copyrightHolderEMail,
      url:  values.copyrightHolderURL
  },
  publisher: {
      id: "",
      name: values.publisherName,
      email: values.publisherEMail,
      url:  values.publisherURL
  }


  <Field
    name='authorName'
    label={FileConfig.authorName}
    component={TextField}
  />
  <Field
    name='authorEMail'
    label={FileConfig.authorEMail}
    component={TextField}
  />
  <Field
    name='authorURL'
    label={FileConfig.authorURL}
    component={TextField}
  />


  <Field
    name='copyrightHolderName'
    label={FileConfig.copyrightHolderName}
    component={TextField}
  />
  <Field
    name='copyrightHolderEMail'
    label={FileConfig.copyrightHolderEMail}
    component={TextField}
  />
  <Field
    name='copyrightHolderURL'
    label={FileConfig.copyrightHolderURL}
    component={TextField}
  />
  <Field
    name='publisherName'
    label={FileConfig.publisherName}
    component={TextField}
  />
  <Field
    name='publisherEMail'
    label={FileConfig.publisherEMail}
    component={TextField}
  />
  <Field
    name='publisherURL'
    label={FileConfig.publisherURL}
    component={TextField}
  />

  */

const addFileSchema = Yup.object().shape({
  hash: Yup.string()
    .required(`${GeneralError.required}`),
  workType: Yup.object({
    value: Yup.number()
      .positive(`${FileConfig.validType}`)
  }),
  url: Yup.string()
    .url(`${FileConfig.validURL}`),
  description: Yup.string()
    .required(`${GeneralError.required}`),
  license: Yup.object({
    value: Yup.number()
      .min(0, `${FileConfig.validLicense}`),
  }),
})

interface FileStateProps {
  info: PayloadProps
  user: GetProps
  address: string
}

interface FileDispatchProps {
  initialise: () => void
  getData: (url: string) => void
  handleSubmit: (values: CreativeWorks) => void
}

type Props =  FileDispatchProps & FileStateProps

export const getFile = (props: Props) => {

    let isFirstRun = useRef(true)
    const [isLoading, setIsLoading] = useState(false)
    const [fileName, setFileName] = useState("")
    const [hash, setHash] = useState("")
    const [workType, setWorkType] = useState({ label: FileConfig.workType, value: 0} as NumberOptionType)
    const [license, setLicense] = useState({ label: FileConfig.license, value: -1} as NumberOptionType)
    const [summary, setSummary] = useState("")
    const [isSubmitting, setSubmit] = useState(false)
    const [user, setUser] = useState({id: "", name: "", email: "", url: ""})
    const [info, setInfo] = useState("")

    const themeClasses = themeStyles()

    useEffect(() => {

        if ( isFirstRun.current ) {

            isFirstRun.current = false
            props.initialise()

            const id = addressToBytes32(props.address)
            props.getData(`${Remote.insecure}${Remote.server}:${Remote.port}${Remote.entities}/${id}`)

        } else if ( ( props.user.data.length > 0 ) && ( user.name == "" ) ) {

            //console.log(props.user.data)
            setUser(props.user.data[0])

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

    }, [props.info, props.user])

    const getFile = (e: any, results: any) => {

        const [load, file] = results[0]
        setFileName(file.name)

        const blobSlice = File.prototype.slice
        const chunkSize = Misc.chunkSize                             // Read in chunks of 2MB
        const chunks = Math.ceil(file.size / chunkSize)
        let currentChunk = 0
        let spark = new SparkMD5.ArrayBuffer()
        const fileReader = new FileReader()

        fileReader.onload = function (e) {
            //console.log('read chunk nr', currentChunk + 1, 'of', chunks);
            spark.append(e.target!.result as ArrayBuffer)                   // Append array buffer
            currentChunk++

            if (currentChunk < chunks) {
                loadNext()
            } else {
                // Compute hash
                setHash(spark.end())
                setIsLoading(false)
            }
        }

        fileReader.onerror = () => {
            setIsLoading(false)
            console.warn(`${FileConfig.loadingError}`)
        }

        const loadNext = () => {

            const start = currentChunk * chunkSize
            const end = ((start + chunkSize) >= file.size) ? file.size : start + chunkSize
            fileReader.readAsArrayBuffer(blobSlice.call(file, start, end))
        }

        loadNext()
    }

    const setLoading = () => {
        setFileName("")
        setHash("")
        setIsLoading(!isLoading)
    }

    return (
      <>
        <h2>{FileConfig.headingFile}</h2>
        <hr />
        <FileReaderInput
          as="binary"
          id="fileInput"
          onChange={getFile}
        >
            <Tooltip title={FileConfig.fileTip}>
                <Okay onClick={setLoading} type='submit' variant="contained" color="primary" disabled={isLoading} endIcon={<RightCircleOutlined spin={isLoading}/>}>
                  {FileConfig.getFile}
                </Okay>
            </Tooltip>
        </FileReaderInput>
        <p>
            {FileConfig.fileName}: {fileName}
        </p>
        <Formik
          initialValues={ {hash: hash} }
          enableReinitialize={true}
          validationSchema={addFileSchema}
          onSubmit={(values: any) => {

            setSubmit(true)
            props.initialise()

            const thisWorkType = (workType as NumberOptionType).value
            const thisLicense = (license as NumberOptionType).value

            let d = new Date(Date.now())

            //ethers.utils.hexlify(hash)

            const fileInfo: CreativeWorks = {
                workType: thisWorkType,
                license: thisLicense,
                id: ethers.utils.formatBytes32String(d.toISOString()),
                authorId: user.id,
                dateCreated: ethers.utils.formatBytes32String(d.toISOString()),
                dateModified: ethers.utils.formatBytes32String(d.toISOString()),
                url: values.url,
                name: fileName,
                description: values.description
            }
            props.handleSubmit(fileInfo)
          }}
        >
          {(formProps: FormikProps<any>) => (
            <Form>
              <FormControl fullWidth={true}>
                  <Field
                    name='hash'
                    label={FileConfig.hash}
                    component={TextField}
                  />
                  <Select
                    name='workType'
                    value={workType}
                    onChange={(ev: ValueType<NumberOptionType>) => {
                        const option = ev as NumberOptionType
                        formProps.setFieldValue('workType', option)
                        setWorkType(option)
                    }}
                    styles={OptionsStyles}
                    options={FormHelpers.worksSelect}
                  />
                  <ErrorMessage name="workType.value" component="div" className={themeClasses.error}/>
                  <Field
                    name='url'
                    label={FileConfig.fileUrl}
                    component={TextField}
                  />
                  <Field
                    name='description'
                    label={FileConfig.description}
                    component={TextField}
                  />
                  <Select
                    name='license'
                    value={license}
                    onChange={(ev: ValueType<NumberOptionType>) => {
                        const option = ev as NumberOptionType
                        formProps.setFieldValue('license', option)
                        setLicense(option)
                    }}
                    styles={OptionsStyles}
                    options={FormHelpers.licensesSelect}
                  />
                  <ErrorMessage name="license.value" component="div" className={themeClasses.error}/>
                  <Grid container>
                      <Grid item xs={12} sm={3}>
                        <Tooltip title={FileConfig.submitTip}>
                          <Okay type='submit' variant="contained" color="primary" disabled={isSubmitting} endIcon={<RightCircleOutlined spin={isSubmitting}/>}>
                            {FileConfig.addFileButton}
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


const mapStateToProps = (state: ApplicationState): FileStateProps => {
  //console.log(state.orgReader)
  return {
    info: state.tx as PayloadProps,
    user: state.data as GetProps,
    address: state.chainInfo.data.Account
  }
}

const mapDispatchToProps = (dispatch: AppDispatch): FileDispatchProps => {
  return {
    initialise: () => dispatch(txInitialise()),
    getData: (url: string) => dispatch(getData(url)),
    handleSubmit: (values: CreativeWorks) => dispatch(addFile(values))
  }
}

export const AddFile = connect<FileStateProps, FileDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(getFile)
