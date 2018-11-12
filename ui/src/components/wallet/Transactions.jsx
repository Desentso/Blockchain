import React, {Component} from 'react';

import WalletCard from "./WalletCard"
import Transaction from "./Transaction"

class Transactions extends Component {
  render() {
    const {transactions, ownAddress} = this.props

    return (
      <WalletCard>
        <h3>Latest Transactions</h3>
        {transactions 
          ? transactions.map(transaction => (
            <Transaction transaction={transaction} ownAddress={ownAddress} key={transaction.Id} />
          ))
          : null
        }
      </WalletCard>
    )
  }
}

export default Transactions
