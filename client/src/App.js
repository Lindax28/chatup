import './App.css';

import { Component } from 'react';

import {
  connect,
  sendMessage,
} from './api';

class App extends Component {
    constructor(props) {
        super(props);
        connect();
    }

    send() {
        console.log("hello");
        sendMessage("hello");
      }

    render() {
        return(
            <div className="App">
                <button onClick={this.send}>Send</button>
            </div>
        )
    }
}

export default App;