import React, {Component} from 'react';
import styled from "styled-components"

import WalletCard from "./WalletCard"

const Address = styled.h4`
  word-break: break-all;
`

class ReceiveTransaction extends Component {
  constructor(props) {
    super(props)
    this.state = {
      copied: null
    }
  }

  copy = () => {
    const selection = window.getSelection();
    const range = document.createRange();
    range.selectNodeContents(this.addressElem);
    selection.removeAllRanges();
    selection.addRange(range);

    console.log(selection, range)

    document.execCommand('copy');

    this.setState({
      copied: "Copied!"
    })

    window.setTimeout(() => this.setState({copied: null}), 2000)
  }
  
  render() {
    const {ownAddress} = this.props
    const {copied} = this.state

    return (
      <WalletCard>
        <h3>Receive payment:</h3>
        <Address onClick={this.copy} ref={(addressElem) => {this.addressElem = addressElem}} >{btoa(ownAddress)}</Address>
        <span>{copied}</span>
      </WalletCard>
    )
  }
}

export default ReceiveTransaction
