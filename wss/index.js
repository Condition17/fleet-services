var app = require('express')();
var http = require('http').createServer(app);
var io = require('socket.io')(http)

app.get('/', (req, res) => {
  res.send({testKey: "value"});
});

io.on('connection', (socket) => {
  console.log('a user is connected')
  io.emit('msg', "this is a test message that should be displayed in front-end app")
  socket.on("initialized", (msg) => {
    console.log("[Socket][Client]:", msg)
  })
})

http.listen(3001, () => {
  console.log('listening on *:3001');
});