const ws = new WebSocket('ws://192.168.1.150:9090');

ws.onopen = () => {
    console.log('Connected to server');
    // Simulate light and sound events here
    setInterval(() => {
        const lightLevel = Math.random() * 100;
        ws.send(JSON.stringify({ type: 'light', value: lightLevel, time: new Date().getTime(), client: 'jsclient' }));
    }, 1000);

    setInterval(() => {
        const soundLevel = Math.random() * 100;
        ws.send(JSON.stringify({ type: 'sound', value: soundLevel, time: new Date().getTime(), client: 'jsclient' }));
    }, 2000);
};

ws.onmessage = (event) => {
    console.log('Server message:', event.data);
};

ws.onclose = () => {
    console.log('Disconnected from server');
};

process.on('SIGINT', function () {
    console.log(`✅ The client has been stopped`, 'Shutdown information', 'This shutdown was initiated by CTRL+C.');
    setTimeout(() => process.exit(0), 1000);
});