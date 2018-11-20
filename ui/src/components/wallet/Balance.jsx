import React, {Component} from 'react';

import WalletCard from "./WalletCard"
import Err from "../shared/Error"

class Balance extends Component {
  render() {

    const {balance = "0", error} = this.props

    return (
      <WalletCard>
        <h1>{balance} coins</h1>
        <Err>{error}</Err>
      </WalletCard>
    )
  }
}

export default Balance
