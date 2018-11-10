import React, {Component} from 'react';
import styled from "styled-components"

const WalletCard = styled.div`

`

class Wallet extends Component {
  render() {
    return (
      <div>
        <div>
          <h1>Wallet</h1>
        </div>

        <div className="Flex">
          <WalletCard>
            <div>
              <h1>123 coins</h1>
            </div>
            <div>
              <h3>New Transaction</h3>

            </div>
          </WalletCard>
          <WalletCard>
            <h3>Latest Transactions</h3>
          </WalletCard>
        </div>
      </div>
    )
  }
}

export default Wallet
