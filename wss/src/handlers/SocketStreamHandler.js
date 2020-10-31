const {REDIS_HOST, REDIS_PORT} = require("../../environment");
const socketIO = require("socket.io");

class SocketStreamHandler {

  init(server) {
    const io = socketIO.listen(server);
    const redisAdapter = require('socket.io-redis');
    io.adapter(redisAdapter({host: REDIS_HOST, port: REDIS_PORT}))
    
    io.on('connection', (socket) => {
      console.log("Enstablished new connection");
      this.handleSocketConnection(socket);
    });
  }

  handleSocketConnection(socket) {
    this.connected = socket;
    socket.on("disconnect", () => {
      console.log("Disconnected");
    });
  }

  sendMessage(message) {
    this.connected.emit("notification", message);
    // check the message type and behave accordingly
  }
}
const socketStreamHandler = new SocketStreamHandler();

module.exports = socketStreamHandler;