import React, {Component} from 'react';
import styled from "styled-components"

import WalletCard from "./WalletCard"

const Address = styled.h4`
  white-space: pre-wrap;
`

class ReceiveTransaction extends Component {
  render() {
    const {ownAddress} = this.props

    return (
      <WalletCard>
        <h3>Receive payment:</h3>
        <Address>{ownAddress}</Address>
      </WalletCard>
    )
  }
}

export default ReceiveTransaction
