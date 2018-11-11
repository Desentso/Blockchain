import React, {Component} from 'react';
import styled from "styled-components"

import Input from "../shared/Input"
import WalletCard from "./WalletCard"
import Button from "../shared/Button"
import {postRequest} from "../../utils/requests"

const Label = styled.label`
  margin-right: 5px;
  width: 200px;
  display: inline-block;
  text-align: left;
`

const InlineLabel = styled.label`
  margin-right: 5px;
  display: inline-block;
  text-align: left;
`

const InputContainer = styled.div`
  display: flex;
  margin: 10px 0;
  flex-wrap: wrap;
`

const AddressInput = styled(Input)`
  flex-grow: 1;
  min-width: 250px;
`

const InlineLabels = styled.div`
  display: flex;
  flex-grow: 1;
  justify-content: space-around;
`

class NewTransaction extends Component {
  constructor(props) {
    super(props)
    this.initialState = {
      address: "",
      amount: 0,
      fees: 0,
      total: 0
    }

    this.state = this.initialState
  }

  addressOnChange = e => this.setState({address: e.target.value})
  amountOnChange = e => {
    const amount = e.target.value
    const fees = 0
    const total = amount

    this.setState({amount, fees, total})
  }

  sendNewPayment = () => {
    const {address, amount} = this.state
    const {ownAddress} = this.props

    console.log(address, amount)

    if (address && amount && amount > 0) {
      console.log("post")
      postRequest("/newTransaction", {from: ownAddress, to: address, amount: parseInt(amount)})
      .then(resp => {
        if (resp.status === 200) {
          this.setState(this.initialState)
        }
      })
      
    } else {
      this.setState({error: "Invalid address"})
    }
  }

  render() {

    const {address, amount, fees, total} = this.state

    return (
      <WalletCard>
        <h3>New Transaction</h3>
        <InputContainer>
          <Label>Receiver Address:</Label>
          <AddressInput type="text" value={address} onChange={this.addressOnChange} />
        </InputContainer>
        <InputContainer>
          <Label>Amount:</Label>
          <Input type="number" min="0" value={amount} onChange={this.amountOnChange} />
          <InlineLabels>
            <div>
              <InlineLabel>Fees:</InlineLabel>
              <span>{fees} coins</span>
            </div>
            <div>
              <InlineLabel>Total:</InlineLabel>
              <span>{total} coins</span>
            </div>
          </InlineLabels>
        </InputContainer>
        <Button onClick={this.sendNewPayment}>Send</Button>
      </WalletCard>
    )
  }
}

export default NewTransaction
