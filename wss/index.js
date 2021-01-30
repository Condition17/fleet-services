const app = require('express')();
const server = require('http').createServer(app);
const eventsSubscriber = require("./src/subscribers/EventsSubscriber");
const socketStreamHandler = require("./src/handlers/SocketStreamHandler");

const formatRecevedEvent = (e) => {
  const event = JSON.parse(e);
  console.log("Received event:", e);
  if (!event.target) {
    console.log("Received event has no specified 'target' key.");
    return event
  }

  event.target = JSON.parse(Buffer.from(event.target, 'base64').toString());

  return event;
}

socketStreamHandler.init(server);
eventsSubscriber.init();

eventsSubscriber.onEventReceived((event) => socketStreamHandler.handle(formatRecevedEvent(event)));

server.listen(3001, () => {
  console.log('listening on *:3001');
});