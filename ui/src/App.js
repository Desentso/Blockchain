import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import { Provider } from 'react-redux'
import { createStore } from 'redux'

import './App.css';
import Header from "./components/shared/Header"
import Home from "./components/home"
import Wallet from "./components/wallet"
import Explorer from "./components/explorer"
import reduxStore from './stores'

//const store = createStore(reduxStore)

class App extends Component {
  render() {
    return (
      <Provider store={reduxStore}>
        <BrowserRouter>
          <div className="App">
            <Header />
            <Route path="/" exact component={Home} />
            <Route path="/wallet" component={Wallet} />
            <Route path="/explorer" component={Explorer} />
          </div>
        </BrowserRouter>
      </Provider>
    );
  }
}

export default App;
