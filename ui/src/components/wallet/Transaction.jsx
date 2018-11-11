import React, {Component} from 'react';
import styled from "styled-components"

const Container = styled.div`
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
`

const Amount = styled.span`
  color: green;
`

class Transaction extends Component {
  render() {
    const {transaction, ownAddress} = this.props

    const total = transaction.Outputs.reduce((total, output) => output.ToAddress === ownAddress ? total + output.Amount : total, 0)

    return (
      <Container>
        <Amount>+ {total} coins</Amount>
        <span>12/11/2018 01:13</span>
        <span>{transaction.Id}</span>
      </Container>
    )
  }
}

export default Transaction
