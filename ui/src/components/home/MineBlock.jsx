import React, {Component} from 'react'
import styled from "styled-components"
import {connect} from "react-redux"

import Card from "../shared/Card"
import Button from "../shared/Button"
import Input from "../shared/Input"

import {toggleMining} from "../../stores/reducers/mining"
import {postRequest} from "../../utils/requests"

const Container = styled.div`
  display: flex;
  justify-content: space-around;
  max-width: 350px;
  margin: auto;
  flex-wrap: wrap;
  margin-top: 50px;
`

const Label = styled.label`
  
`

class MineBlock extends Component {

  mineNewBlock = () => {
    postRequest("/mineblock", {data: "new block"})
    .then(resp => {

    })
  }

  toggleMining = () => {
    const {mining, intervalID} = this.props.mining
    if (mining) {
      window.clearInterval(intervalID)
      this.props.toggleMining(null)
    } else {
      const newIntervalID = window.setInterval(this.mineNewBlock, 1000)
      this.props.toggleMining(newIntervalID)
    }
  }

  render() {
    const {mining} = this.props.mining

    return (
      <Container>
        <Button onClick={this.mineNewBlock}>Mine One Block</Button>
        <Button onClick={this.toggleMining}>{mining ? "Stop" : "Start"} Mining</Button>
      </Container>
    )
  }
}

const mapStateToProps = state => ({
  mining: state.mining
})

const mapDispatchToProps = {
  toggleMining
}

export default connect(mapStateToProps, mapDispatchToProps)(MineBlock)
