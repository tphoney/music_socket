# Music_socket

This is a simple javascript web application that allows users to affect music together in real time. Users can add visual effect and audio effects to the music. The music is played on the server and the audio effects are applied on the server.

## Getting Started

### Server

To get started, clone the repository and run the following commands:

```bash
docker run -it --rm --name music_socket -v "$PWD":/usr/src/app -w /usr/src/app node:current-alpine npm install
docker run -it --rm --name music_socket -v "$PWD":/usr/src/app -w /usr/src/app -p 9090:9090 node:current-alpine node server.js
```

### Client

```bash
docker run -it --rm --name client -v "$PWD":/usr/src/app -w /usr/src/app node:current-alpine npm install
docker run -it --rm --name client -v "$PWD":/usr/src/app -w /usr/src/app -p 8080:8080 node:current-alpine node client.js
```


### Output

```bash
Server listening on port 9090
Client connected
Light level: 18.378846958753048 delay: 2ms Sent @ 1718029683012 recieved @ 1718029683014
Sound level: 12.918869839974501 delay: 1ms Sent @ 1718029684013 recieved @ 1718029684014
Light level: 23.500039531560212 delay: 4ms Sent @ 1718029684014 recieved @ 1718029684018
Light level: 93.40424473177195 delay: 2ms Sent @ 1718029685015 recieved @ 1718029685017
Sound level: 79.66534014640847 delay: 0ms Sent @ 1718029686015 recieved @ 1718029686015
Light level: 69.11103222529727 delay: 1ms Sent @ 1718029686015 recieved @ 1718029686016
Light level: 66.9189311713361 delay: 1ms Sent @ 1718029687016 recieved @ 1718029687017
Sound level: 81.76191798049138 delay: 2ms Sent @ 1718029688015 recieved @ 1718029688017
Light level: 94.9991528501585 delay: 2ms Sent @ 1718029688016 recieved @ 1718029688018
Light level: 47.49824236600602 delay: 1ms Sent @ 1718029689017 recieved @ 1718029689018
Sound level: 58.67693557449258 delay: 1ms Sent @ 1718029690016 recieved @ 1718029690017
Light level: 40.2727309948449 delay: 1ms Sent @ 1718029690018 recieved @ 1718029690019
Light level: 50.654070848299604 delay: 1ms Sent @ 1718029691018 recieved @ 1718029691019
^C✅ The server has been stopped Shutdown information This shutdown was initiated by CTRL+C.
Sound level: 3.1106696121108834 delay: 0ms Sent @ 1718029692017 recieved @ 1718029692017
Light level: 82.03426992322748 delay: 0ms Sent @ 1718029692018 recieved @ 1718029692
```
