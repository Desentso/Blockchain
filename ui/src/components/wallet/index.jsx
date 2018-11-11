import React, {Component} from 'react';
import styled from "styled-components"
import {connect} from "react-redux"

import NewTransaction from './NewTransaction'
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

  componentWillReceiveProps(nextProps) {
    if (nextProps.data.address && !this.state.transactionsInited) {
      this.props.getTransactions()
      this.setState({transactionsInited: true})
    }
  } 

  render() {
    console.log(this.props)
    const {address, balance, transactions} = this.props.data

    return (
      <div>
        <div>
          <h1>Wallet</h1>
        </div>

        <FlexContainer>
          <FlexElementContainer>
            <Balance balance={balance} />
            <NewTransaction ownAddress={address} />
          </FlexElementContainer>
          <Transactions transactions={transactions} ownAddress={address} />
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
