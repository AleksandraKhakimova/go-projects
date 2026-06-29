// ws-gateway.js
const WebSocket = require('ws');
const net = require('net');

const wss = new WebSocket.Server({ port: 8081 });
const TCP_HOST = 'localhost';
const TCP_PORT = 8080;

console.log('WebSocket шлюз запущен на порту 8081');

wss.on('connection', (ws) => {
    const tcpClient = new net.Socket();
    
    tcpClient.connect(TCP_PORT, TCP_HOST, () => {
        console.log('TCP подключен к бекенду');
    });

    tcpClient.on('data', (data) => {
        ws.send(data.toString());
    });

    ws.on('message', (message) => {
        tcpClient.write(message);
    });

    ws.on('close', () => {
        tcpClient.destroy();
        console.log('Клиент отключён');
    });

    tcpClient.on('error', (err) => {
        console.error('TCP ошибка:', err.message);
        ws.close();
    });

    ws.on('error', (err) => {
        console.error('WS ошибка:', err.message);
        tcpClient.destroy();
    });
});