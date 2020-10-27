const {REDIS_HOST, REDIS_PORT} = require("./environment");
const app = require('express')();
const http = require('http').createServer(app);
const io = require('socket.io')(http)
const redisAdapter = require('socket.io-redis');

io.adapter(redisAdapter({host: REDIS_HOST, port: REDIS_PORT}))

app.get('/', (req, res) => {
  res.send({testKey: "value"});
});

io.on('connection', (socket) => {
  console.log('a user is connected')
  io.emit('msg', "this is a test message that should be displayed in front-end app")
  socket.on("initialized", (msg) => {
    console.log("[Socket][Client]:", msg)
  })
});

http.listen(3001, () => {
  console.log('listening on *:3001');
});