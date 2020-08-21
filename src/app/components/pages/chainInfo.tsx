import * as React from 'react'
import { connect } from 'react-redux'
import Markdown from 'react-markdown'

import { ApplicationState, ChainInfoProps } from '../../store'
import { get } from '../../utils/list'

import { Blockchain } from '../../config/strings'

interface InfoProps {
  chainData: ChainInfoProps
}

const info = (props: InfoProps) => {

  const chainInfo = get(props.chainData.data)

  return (
      <div>
        <h2>{Blockchain.heading}</h2>
        <Markdown escapeHtml={false} source={chainInfo} />
      </div>
  )
}

const mapStateToProps = (state: ApplicationState): InfoProps => {

    const info = state.chainInfo as ChainInfoProps
    delete info.data.Provider
    return {
      chainData: info
    }
}

export const BlockchainInfo = connect<InfoProps, {}, {}, ApplicationState>(
  mapStateToProps
)(info)
