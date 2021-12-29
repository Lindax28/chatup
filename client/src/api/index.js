var socket = new WebSocket("ws://localhost:8080/ws");

let connect = msgCB => {
    console.log("Connecting...");

    socket.onopen = () => {
        console.log("Successfully Connected");
      };
    
      socket.onmessage = msg => {
        console.log(msg);
        msgCB(msg)
      };
    
      socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
      };
    
      socket.onerror = error => {
        console.log("Socket Error: ", error);
      };
};

let sendMessage = (msg) => {
    console.log("Sending message: ", msg);
    socket.send(msg);
};

export { connect, sendMessage };