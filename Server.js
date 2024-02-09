const express = require('express');
const http = require('http');
const WebSocket = require('ws');

const PORT = process.env.PORT || 3000;

const app = express();
const server = http.createServer(app);
const wss = new WebSocket.Server({ server });

// Serve static files (optional)
app.use(express.static('public'));

// WebSocket connection handler
wss.on('connection', function connection(ws) {
    console.log('A new client connected!');

    // Message event handler
    ws.on('message', function incoming(message) {
        console.log('Received message:', message);

        // Broadcast message to all clients except the sender
        wss.clients.forEach(function each(client) {
            if (client !== ws && client.readyState === WebSocket.OPEN) {
                client.send(message);
            }
        });
    });

    // Close event handler
    ws.on('close', function close() {
        console.log('Client disconnected');
    });
});

// Start the server
server.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
