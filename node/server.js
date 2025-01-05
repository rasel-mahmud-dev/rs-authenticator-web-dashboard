const WebSocket = require('ws');
const server = new WebSocket.Server({port: 8084});

let clients = {};

server.on('connection', (socket) => {

    console.log('Client connected');

    socket.on('message', (data) => {
        const message = JSON.parse(data);

        if (message.identifier === "client") {
            console.log(`Received identifier from client: ${message.identifier}`);
            clients[message.identifier] = socket;
        } else {
            broadcastToAllClients(message.content);
        }

        console.log(clients.length)

    });

    socket.on('close', () => {
        console.log('Client disconnected');
    });
});

function broadcastToAllClients(content) {
    const clientSocket = clients["client"];
    if (clientSocket.readyState === WebSocket.OPEN) {
        clientSocket.send("reload");
    }
}