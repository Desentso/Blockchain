import React, {Component} from 'react';

import WalletCard from "./WalletCard"

class Balance extends Component {
  render() {

    const {balance} = this.props

    return (
      <WalletCard>
        <h1>{balance} coins</h1>
      </WalletCard>
    )
  }
}

export default Balance
