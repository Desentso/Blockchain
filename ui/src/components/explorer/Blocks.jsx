import React, {Component} from 'react';
import styled from "styled-components"

import Block from "./Block"
import Err from "../shared/Error"

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

    const {blockchain, error} = this.props

    return (
      <Container>
        <Flex>
          <Label>Index</Label>
          <BlockHashLabel>Block Hash</BlockHashLabel>
          <Label>Timestamp</Label>
          <Label>No. of Transactions</Label>
        </Flex>
        {error
          ? <Err>{error}</Err>
          : blockchain
            ? blockchain.sort((a,b) => b.timestamp - a.timestamp)
              .map(block => 
                <Block block={block} key={block.hash} />
              )
            : null
        }
      </Container>
    )
  }
}

export default Blocks
