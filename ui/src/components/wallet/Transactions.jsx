import React, {Component} from 'react';

import WalletCard from "./WalletCard"
import Transaction from "./Transaction"

class Transactions extends Component {
  render() {
    const {transactions, pendingTransactions, ownAddress} = this.props

    return (
      <WalletCard>
        <h3>Latest Transactions</h3>
        {pendingTransactions 
          ? pendingTransactions.sort((a,b) => b.timestamp - a.timestamp).map(transaction => (
            <Transaction transaction={transaction} ownAddress={ownAddress} key={transaction.id} pending />
          ))
          : null
        }
        {transactions 
          ? transactions.sort((a,b) => b.timestamp - a.timestamp).map(transaction => (
            <Transaction transaction={transaction} ownAddress={ownAddress} key={transaction.id} />
          ))
          : null
        }
      </WalletCard>
    )
  }
}

export default Transactions
