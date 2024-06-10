const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 9090 });

wss.on('connection', (ws) => {
    console.log('Client connected');

    ws.on('message', (message) => {
        const data = JSON.parse(message);

        const timeDifference = new Date().getTime() - data.time;
        if (data.type === 'light') {            
            console.log(`Light level: ${data.value} delay: ${timeDifference}ms Sent @ ${data.time} recieved @ ${new Date().getMilliseconds}. Client: ${data.client}`);
        } else if (data.type === 'sound') {
            console.log(`Sound level: ${data.value} delay: ${timeDifference}ms Sent @ ${data.time} recieved @ ${new Date().getTime()}. Client: ${data.client}`);
        } else {
            console.warn(`Unknown event type: ${data.type}`);
        }
    });

    ws.on('close', () => {
        console.log('Client disconnected');
    });
});

console.log('Server listening on port 9090');


process.on('SIGINT', function () {
    console.log(`✅ The server has been stopped`, 'Shutdown information', 'This shutdown was initiated by CTRL+C.');
    setTimeout(() => process.exit(0), 1000);
});