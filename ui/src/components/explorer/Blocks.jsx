import React, {Component} from 'react';
import styled from "styled-components"

import Card from "../shared/Card"
import Block from "./Block"

const Container = styled.div`
  padding-bottom: 150px;
`

const Flex = styled.div`
  display: flex;
  justify-content: space-between;
  padding: 0 5%;
  align-items: center;
  flex-wrap: wrap;
`

const Label = styled.p`
  font-weight: 600;
`

const BlockHashLabel = styled(Label)`
  width: 575px;
`

class Blocks extends Component {
  constructor(props) {
    super(props)
    this.state = {
      copied: null
    }
  }
  
  render() {

    const {blockchain} = this.props

    return (
      <Container>
        <Flex>
          <Label>Index</Label>
          <BlockHashLabel>Block Hash</BlockHashLabel>
          <Label>Timestamp</Label>
          <Label>No. of Transactions</Label>
        </Flex>
        {blockchain
          ? blockchain.sort((a,b) => b.timestamp - a.timestamp)
            .map(block => 
              <Block block={block} />
            )
          : null
        }
      </Container>
    )
  }
}

export default Blocks
