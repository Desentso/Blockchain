import React, {Component} from 'react';
import styled from "styled-components"
import {connect} from "react-redux"

import Input from "../shared/Input"
import WalletCard from "./WalletCard"
import NewTransaction from './NewTransaction';
import Balance from './Balance';
import {loadData} from "../../stores/reducers/basicData"

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

  componentWillMount() {
    this.props.loadData()
  }

  render() {
    console.log(this.props)
    return (
      <div>
        <div>
          <h1>Wallet</h1>
        </div>

        <FlexContainer>
          <FlexElementContainer>
            <Balance />
            <NewTransaction ownAddress={this.props.data.address} />
          </FlexElementContainer>
          <WalletCard>
            <h3>Latest Transactions</h3>
          </WalletCard>
        </FlexContainer>
      </div>
    )
  }
}

const mapStateToProps = state => ({
  data: state.data
})

const mapDispatchToProps = {
  loadData
}

export default connect(mapStateToProps, mapDispatchToProps)(Wallet)
