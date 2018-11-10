import React from 'react';
import {Link} from "react-router-dom"

const Header = () => {
  return (
    <div className="Header">
      <Link to="/" className="Header-Link" >Home</Link>
      <Link to="/wallet" className="Header-Link" >Wallet</Link>
      <Link to="/explorer" className="Header-Link" >Explorer</Link>
    </div>
  )
}

export default Header
