// server.js
const WebSocket = require('ws');

// Specify the IP address and port to bind the server to
const ip = '192.168.68.116'; // You can use 'localhost' or the actual IP address of your machine
const port = 8080;

// Create a WebSocket server instance and bind it to the specified IP address and port
const wss = new WebSocket.Server({ host: ip, port: port });

// Event listener for connection establishment
wss.on('connection', function connection(ws) {
  console.log('Client connected');

  // Send a welcome message to the client


  // Event listener for receiving messages from the client
  ws.on('message', function incoming(message) {
    console.log('Received message:', message);

    // Broadcast the received message to all clients except the sender
    wss.clients.forEach(function each(client) {
      if (client !== ws && client.readyState === WebSocket.OPEN) {
        client.send(message);
      }
    });
  });

  // Event listener for client disconnection
  ws.on('close', function close() {
    console.log('Client disconnected');
  });
});

console.log(`WebSocket server is running at ws://${ip}:${port}`);
