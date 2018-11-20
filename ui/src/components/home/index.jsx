import React, {Component} from 'react'
import styled from "styled-components"

import Card from "../shared/Card"
import ControlPanel from "./ControlPanel"

const ContainerCard = styled(Card)`
  margin-top: 20px;
`

const Title = styled.h1`
  margin-top: 0;
`

const List = styled.ul`
  margin: auto;
  max-width: 400px;
  text-align: left;
`

class Home extends Component {
  render() {
    return (
      <div>
        <ContainerCard>
          <Title>Home</Title>
          <h3>A MVB (Minimum viable blockchain) implemented with Go. </h3>
          <h3>Features include:</h3>
          <List>
            <li>Peer-to-peer nodes</li>
            <li>Proof-of-work mining function</li>
            <li>Transactions</li>
            <li>Wallets with RSA used for addresses and signing</li>
            <li>Wallet UI</li>
            <li>Block explorer UI</li>
          </List>
        </ContainerCard>

        <ControlPanel />

      </div>
    )
  }
}

export default Home
