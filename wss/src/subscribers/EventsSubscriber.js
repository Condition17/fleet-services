const {wssSubscriptionName} = require("../../constants");
const {EventEmitter} = require("events");
const {PubSub} = require('@google-cloud/pubsub');

const NOOP_FUNCTION = () => {};
const events = {
  EVENT_RECEIVED: "event_received",
};

class EventsSubscriber {
  constructor() {
    this.pubSubClient = new PubSub();
    this.subscription = null;
    this.emitter = new EventEmitter();
  }

  async init() {
    this.subscription = await this.pubSubClient.subscription(wssSubscriptionName)
    console.log(`Subscription was setup - subscription '${wssSubscriptionName}'`);
    this.subscription.on("message", (message) => {
      console.log("Received message on subscription")
      this.emitter.emit(events.EVENT_RECEIVED, message.data.toString());
      message.ack();
    })
  }

  onEventReceived(callback = NOOP_FUNCTION) {
    this.emitter.on(events.EVENT_RECEIVED, callback)
  }
}

const eventsSubscriber = new EventsSubscriber();

module.exports = eventsSubscriber;