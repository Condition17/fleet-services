const app = require('express')();
const server = require('http').createServer(app);
const eventsSubscriber = require("./src/subscribers/EventsSubscriber");
const socketStreamHandler = require("./src/handlers/SocketStreamHandler");

socketStreamHandler.init(server);
eventsSubscriber.init();

eventsSubscriber.onEventReceived((event) => {
  console.log("Precessing '", event,"' externally");
  socketStreamHandler.sendMessage(event);
});

server.listen(3001, () => {
  console.log('listening on *:3001');
});