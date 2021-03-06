import React, {Component} from 'react';
import styled from "styled-components"
import {connect} from "react-redux"

import SendNewTransaction from './NewTransaction'
import ReceiveTransaction from './ReceiveTransaction'
import Balance from './Balance'
import Transactions from './Transactions'
import {loadData, getBalance, getTransactions} from "../../stores/reducers/basicData"

const FlexContainer = styled.div`
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
  margin: 0 15px;
`

const FlexElementContainer = styled.div`
  flex-grow: 1;
  max-width: 950px;
`


class Wallet extends Component {

  constructor(props) {
    super(props)

    this.state = {
      transactionsInited: false
    }
  }

  componentWillMount() {
    this.props.loadData()
    this.props.getBalance()
  }

  componentWillUnmount() {
    window.clearInterval(this.state.removeInterval)
  }

  componentWillReceiveProps(nextProps) {
    const {getTransactions, getBalance} = this.props
    if (nextProps.data.address && !this.state.transactionsInited) {
      getTransactions()
      const removeInterval = window.setInterval(() => {getTransactions(); getBalance()}, 2000)
      this.setState({transactionsInited: true, removeInterval})
    }
  } 

  render() {
    const {
      address, 
      balance, 
      finishedTransactions, 
      pendingTransactions,
      addressError,
      balanceError,
      txsError
    } = this.props.data

    return (
      <div>
        <div>
          <h1>Wallet</h1>
        </div>

        <FlexContainer>
          <FlexElementContainer>
            <Balance balance={balance} error={balanceError} />
            <SendNewTransaction ownAddress={address} />
            <ReceiveTransaction ownAddress={address} error={addressError} />
          </FlexElementContainer>
          <Transactions 
            transactions={finishedTransactions} 
            pendingTransactions={pendingTransactions} 
            ownAddress={address} 
            error={txsError}
          />
        </FlexContainer>
      </div>
    )
  }
}

const mapStateToProps = state => ({
  data: state.data
})

const mapDispatchToProps = {
  loadData,
  getBalance,
  getTransactions
}

export default connect(mapStateToProps, mapDispatchToProps)(Wallet)
