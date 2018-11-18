import React, {Component} from 'react'
import styled from "styled-components"

import Card from "../shared/Card"
import Button from "../shared/Button"
import Input from "../shared/Input"
import MineBlock from "./MineBlock"
import AddPeer from "./AddPeer"

import {postRequest} from "../../utils/requests"

const Label = styled.label`
  
`

class ControlPanel extends Component {

  render() {
    return (
      <Card>
        <h2>Control Panel</h2>
        <MineBlock />
        <AddPeer />
      </Card>
    )
  }
}

export default ControlPanel
