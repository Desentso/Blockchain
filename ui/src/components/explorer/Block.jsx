import React, {Component} from 'react';
import styled from "styled-components"

import {formatTimestamp} from "../../utils/time"
import Card from "../shared/Card"
import BlockTransaction from './BlockTransaction';

const Container = styled(Card)`
  border-left: 7px solid #4353fe;
`

const Flex = styled.div`
  display: flex;
  justify-content: space-between;
  padding: 0 4%;
  align-items: center;
  flex-wrap: wrap;
  cursor: pointer;
`

const Title = styled.h4`
  margin: 0;
  word-break: break-word;
`

const Expanded = styled.div`
  width: 90%;
  margin-left: 10%;
`


class Block extends Component {
  constructor(props) {
    super(props)
    this.state = {
      expanded: false
    }
  }

  toggleExpand = () => this.setState({expanded: !this.state.expanded})
  
  render() {

    const {block} = this.props
    const {expanded} = this.state

    return (
      <Container>
        <Flex onClick={this.toggleExpand}>
          <Title>{block.index}</Title>
          <Title>{block.hash}</Title>
          <p>{formatTimestamp(block.timestamp)}</p>
          <p>{block.transactions ? block.transactions.length : 0} Transactions</p>
        </Flex>
        {expanded
          ? <Expanded>
              {block.transactions
                ? block.transactions.map(tx => 
                    <BlockTransaction transaction={tx} key={tx.id} />
                  )
                : "No transactions"
              }
            </Expanded>
          : null
        }
      </Container>
    )
  }
}

export default Block
