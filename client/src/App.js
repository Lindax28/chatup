import './App.css';

import React, {
  useEffect,
  useState,
} from 'react';

import {
  connect,
  sendMessage,
} from './api';
import ChatHistory from './components/ChatHistory';
import Header from './components/Header/Header';

export default function App(props) {
  const [chatHistory, setChatHistory] = useState([]);
  useEffect(() => {
    connect((msg) => {
      console.log("New Message");
      setChatHistory([...chatHistory, msg]);
    });
  });

  const send = () => {
    sendMessage("hello");
  };

  return(
      <div className="App">
          <Header />
          <ChatHistory chatHistory={chatHistory} />
          <button onClick={send}>Send</button>
      </div>
  );
};