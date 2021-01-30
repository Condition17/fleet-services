const {isValidToken, getDecodedTokenValue} = require("../auth");
const {REDIS_HOST, REDIS_PORT} = require("../../environment");
const socketIO = require("socket.io");
const socketIdByUser = {};


class SocketStreamHandler {
  init(server) {
    this.socketIOServer = socketIO.listen(server);
    const redisAdapter = require('socket.io-redis');
    this.socketIOServer.adapter(redisAdapter({host: REDIS_HOST, port: REDIS_PORT}))

    this.socketIOServer.use(this.authorizeConnection)
    this.socketIOServer.use(this.storeConnectionDetails)
    this.socketIOServer.on('connection', (socket) => {
      const user = getDecodedTokenValue(socket.handshake.query.token)["User"];
      socketIdByUser[user.id] = socket.id;
      socket.on("disconnect", () => {
          delete socketIdByUser[user.id];
      });
    });
  }

  authorizeConnection(socket, next) {
    if(!socket.handshake.query.token ){
      socket.disconnect('unauthorized');
    }

    const token = socket.handshake.query.token
    if (!isValidToken(token)) {
      next(new Error("Invalid token"));
    }

    next();
  }

  handle(event) {
    if (!event.target) {
      console.log("Unhandled event - no target");
      return;
    }
    const targetSocketId = socketIdByUser[event.target.id];
    delete event.target;

    if (targetSocketId) {
      this.socketIOServer.sockets.connected[targetSocketId].emit("message", JSON.stringify(event))
    }
  }
}
const socketStreamHandler = new SocketStreamHandler();

module.exports = socketStreamHandler;