const express = require('express');
const app = express();
const http = require('http');
const WebSocket = require('ws');
const path = require('path');

const PORT = process.env.PORT || 3000;

// Create HTTP server
const server = http.createServer(app);

// Create WebSocket server
const wss = new WebSocket.Server({ server });

// WebSocket connection handler
wss.on('connection', function connection(ws) {
  console.log('vanakam da mapla server la irunthu!');
  ws.send('Welcome, new client!');

  // WebSocket message handler
  ws.on('message', function incoming(message) {
    console.log('Received: %s', message);

    // Broadcast message to all clients except the sender
    wss.clients.forEach(function each(client) {
      if (client !== ws && client.readyState === WebSocket.OPEN) {
        client.send(message);
      }
    });
  });
});

// Serve static files in production
if (process.env.NODE_ENV === 'production') {
  app.use(express.static(path.join(__dirname, 'public')));
}

// Define routes
app.get('/', (req, res) => res.send('Hello, World!'));

// Start the server
server.listen(PORT, () => console.log(`Server listening on port ${PORT}`));
