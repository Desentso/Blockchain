import React, {Component} from 'react';
import styled from "styled-components"

import {formatTimestamp} from "../../utils/time"

const Container = styled.div`
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
`

const Amount = styled.span`
  color: ${props => props.receivedTransaction
    ? "green"
    : "red"
  };
`

class Transaction extends Component {
  render() {
    const {transaction, ownAddress, pending} = this.props

    const receivedTransaction = transaction.to === ownAddress

    const total = receivedTransaction
      ? transaction.outputs.reduce((total, output) => output.toAddress === ownAddress ? total + output.amount : total, 0)
      : transaction.outputs.reduce((total, output) => output.toAddress !== ownAddress ? total + output.amount : total, 0)

    return (
      <Container>
        <Amount receivedTransaction={receivedTransaction}>{receivedTransaction ? "+" : "-"} {total} coins</Amount>
        <span>{formatTimestamp(transaction.timestamp)}</span>
        <span>{transaction.id}</span>
        <span>{pending ? "Pending" : "Finished"}</span>
      </Container>
    )
  }
}

export default Transaction
