// server.js
const WebSocket = require('ws');
const port = 8080;

const wss = new WebSocket.Server({ port: port });

wss.on('listening', () => {
  console.log(`WebSocket server is running at ws://localhost:${port}`);
});

wss.on('connection', function connection(ws) {
  console.log('Client connected');

  ws.send('Welcome to the WebSocket server!');

  ws.on('message', function incoming(message) {
    console.log('Received message:', message);

    wss.clients.forEach(function each(client) {
      if (client !== ws && client.readyState === WebSocket.OPEN) {
        client.send(message);
      }
    });
  });

  ws.on('close', function close() {
    console.log('Client disconnected');
  });

  ws.on('error', function error(err) {
    console.error('WebSocket error:', err);
  });
});

wss.on('error', function error(err) {
  console.error('WebSocket server error:', err);
});
