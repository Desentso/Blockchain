import React, {Component} from 'react'
import {connect} from "react-redux"

import {loadBlockchain} from "../../stores/reducers/blockchain"
import Blocks from "./Blocks"

class Explorer extends Component {

  componentWillMount() {
    this.props.loadBlockchain()
  }

  render() {
    const {blockchain, error} = this.props.blockchain

    return (
      <div>
        <h1>Explorer</h1>
        <h3>Explore the blockchain</h3>
        <Blocks blockchain={blockchain} error={error} />
      </div>
    )
  }
}

const mapStateToProps = state => ({
  blockchain: state.blockchain
})

const mapDispatchToProps = {
  loadBlockchain
}

export default connect(mapStateToProps, mapDispatchToProps)(Explorer)
