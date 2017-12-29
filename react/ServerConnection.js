/**
 * Manages a web socket connection to the game backend.
 */

import settings from "./settings";


export default class ServerConnection {
  constructor(opts) {
    this._onMessage = this._onMessage.bind(this);
    this.callbacks = {};
    this.ws = new WebSocket(settings.WEBSOCKETS_URL);
    this.ws.onopen = () => opts.onOpen();
    this.ws.onmessage = this._onMessage;
  }

  _onMessage(evt) {
    // Try to parse contents
    var contents;

    try {
      contents = JSON.parse(evt.data);
    } catch(e) {
      console.log("Got message from backend that couldn't parse");
      return;
    }

    // Call handler if one is defined
    const handler = this.callbacks[contents.EventType];

    if (handler) {
      handler(contents.Payload);
    }
  }

  on(eventType, callback) {
    this.callbacks[eventType] = callback;
  }

  clientInit(initOpts) {
    this.ws.send(JSON.stringify({
      EventType: "clientInit",
      name: initOpts.name,
      character: initOpts.character,
    }));
  }
}
