import React, {Component} from 'react'
import styled from "styled-components"

import Button from "../shared/Button"
import Input from "../shared/Input"

import {postRequest} from "../../utils/requests"

const Container = styled.div`
  display: flex;
  margin: auto;
  max-width: 625px;
  justify-content: space-around;
  align-items: center;
  flex-wrap: wrap;
  margin-top: 20px;
`


const Label = styled.label`
  margin-right: 5px;
`



class AddPeer extends Component {

  constructor(props) {
    super(props)
    
    this.state = {
      peerIP: null,
      peerPort: null
    }
  }

  ipOnChange = e => this.setState({peerIP: e.target.value})
  portOnChange = e => this.setState({peerPort: e.target.value})

  addPeer = () => {
    const {peerPort: port, peerIP: address} = this.state
    postRequest("/peer", {port, address})
  }

  render() {
    return (
      <Container>

        <div>
          <Label>Url/IP:</Label>
          <Input onChange={this.ipOnChange} type="text" />
        </div>

        <div>
          <Label>Port:</Label>
          <Input onChange={this.portOnChange} type="text" />
        </div>

        <Button onClick={this.addPeer}>Add Peer</Button>
      </Container>
    )
  }
}

export default AddPeer
