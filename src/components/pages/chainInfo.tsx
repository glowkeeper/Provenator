import * as React from 'react'
import { connect } from 'react-redux'
import Markdown from 'react-markdown'

import { ApplicationState, ChainDataProps } from '../../store'
import { get } from '../../utils/list'

import { Blockchain } from '../../config/strings'

interface InfoProps {
  chainData: object
}

const info = (props: InfoProps) => {

  const chainInfo = get(props.chainData)

  return (
      <div>
        <h2>{Blockchain.heading}</h2>
        <Markdown escapeHtml={false} source={chainInfo} />
      </div>
  )
}

const mapStateToProps = (state: ApplicationState): InfoProps => {

    return {
      chainData: state.chainInfo as ChainDataProps
    }
}

export const BlockchainInfo = connect<InfoProps, {}, {}, ApplicationState>(
  mapStateToProps
)(info)
