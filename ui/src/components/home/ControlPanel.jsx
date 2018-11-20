import React, {Component} from 'react'

import Card from "../shared/Card"
import MineBlock from "./MineBlock"
import AddPeer from "./AddPeer"

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
