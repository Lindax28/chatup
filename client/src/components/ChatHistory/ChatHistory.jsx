import './ChatHistory.scss';

import React from 'react';

export default function ChatHistory(props) {
    const messages = props.chatHistory.map((msg, idx) => (
        <p key={idx}>
            {msg.data}
        </p>
    ));

    return (
    <div className="chat-history">
        <h2>Chat</h2>
        {messages}
    </div>
    );
};