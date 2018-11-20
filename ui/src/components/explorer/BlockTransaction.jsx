import React, {Component} from 'react';
import styled from "styled-components"


const Flex = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-around;
  border: 1px solid #ddd;
  border-left: 5px solid #ddd;
  border-radius: 3px;
  margin: 20px 0;
  flex-wrap: wrap;
`

const TxContainer = styled.div`
  width: 47%;
  min-width: 200px;
  flex-grow: 5;
`

const Address = styled.h4`
  white-space: pre-wrap;
  word-break: break-word;
`

const ToIcon = styled.span`
  font-size: 30px;
`

class BlockTransaction extends Component {

  render() {

    const {transaction} = this.props

    const amount = transaction.outputs.reduce((sum, output) => output.toAddress === transaction.to ? sum + output.amount : sum, 0)

    return (
      <Flex>
        <TxContainer>
          <Address>{transaction.from ? transaction.from : "Mining reward"}</Address>
        </TxContainer>
        <div>
          <ToIcon>&#10140;</ToIcon>
          <p>{amount} coins</p>
        </div>
        <TxContainer>
          <Address>{transaction.to}</Address>
        </TxContainer>
      </Flex>
    )
  }
}

export default BlockTransaction
