const WebSocket = require('ws');

const socket = new WebSocket('ws://192.168.0.148:8084');

socket.on('open', () => {
    console.log('Connected to WebSocket server');
    const identifier = "node-client";
    socket.identifier = identifier;
    console.log(`WebSocket connection opened with identifier: ${socket.identifier}`);
    socket.send(JSON.stringify({type: "register", identifier: socket.identifier}));
});

socket.on('message', (message) => {
    console.log('Message from server:', message);
});

socket.on('close', () => {
    console.log('Disconnected from WebSocket server');
});

socket.on('error', (error) => {
    console.log('sdljkfsdjf')
    console.error('WebSocket error:', error?.message);
});
