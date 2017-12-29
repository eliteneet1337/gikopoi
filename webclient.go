/**
 * The web client is an object for each websocket connection. It is a user.
 *
 * This layer primarily manages imtermediary logic between the hub and browser
 * (via websockets). The hub may communicate with the web client by sending it
 * a `WebClientIncomingEvent` through the `webClient.incomingEvents` channel.
 */
package main

import (
  "log"
  "encoding/json"
  "github.com/gorilla/websocket"
)

/**
 * WebClient Struct definition
 */
type WebClient struct {
  name string                            // Visible username in the room
  uuid string                            // Unique identifier for this user
  character string                       // Name of character (mona, giko, etc)
  hub *Hub                               // Pointer to master hub
  conn *websocket.Conn                   // Pointer to websocket connection

  // Joint communication channel with hub and browser
  incomingEvents chan *WebClientIncomingEvent
}


/**
 * The web client incoming event wraps events from the hub or the websocket.
 */
type WebClientIncomingEvent struct {
  eventType string                      // event identifier
  source string                         // either "hub" or "client"
  payload map[string]interface{}        // when from client, decoded from JSON
}


/**
 * Initialize a new webclient.
 *
 * Returns a struct with the default fields initialized.
 */
func newWebClient(hub *Hub, conn *websocket.Conn) *WebClient {
  return &WebClient{
    name: "",
    uuid: "",
    character: "",
    hub: hub,
    conn: conn,
    incomingEvents: make(chan *WebClientIncomingEvent),
  }
}


/**
 * Process the incoming events for this client.
 *
 * These events may be written by the hub or locally through the method
 * pumpWebsocketMessages().
 */
func (webClient *WebClient) handleIncomingEvents() {
  for {
    incomingEvent := <- webClient.incomingEvents
    log.Println("client got event", incomingEvent.eventType,
      incomingEvent.payload);

    if incomingEvent.eventType == "clientInit" {
      webClient.onClientInit(incomingEvent.payload);
    } else if incomingEvent.eventType == "clientResp" {
      webClient.onClientResp(incomingEvent);
    } else {
      log.Printf("Do not know how to handle event ", incomingEvent.eventType)
    }
  }
}

/**
 * Hnadler for the clientResp event from the hub.
 *
 * The hub sends the payload for the response in the incoming event, here we
 * just echo it over to the websocket.
 */
func (webClient *WebClient) onClientResp(incomingEvent *WebClientIncomingEvent) {
  message := &ClientMessage{
    EventType: "clientInitResp",
    Payload: incomingEvent.payload,
  }

  jsonOut, err := json.Marshal(message)

  log.Println("error was ", err)
  log.Println("sending message", string(jsonOut))

  if err == nil {
    webClient.conn.WriteMessage(websocket.TextMessage, jsonOut)
  }
}


/**
 * Handler for clientInit event arriving from the browser.
 *
 * Set some initial fields on the web client struct and then send it over to
 * the hub for registration.
 */
func (webClient *WebClient) onClientInit(payload interface{}) {
  // Set attributes on the instance
  payloadAsMap := payload.(map[string]interface{})

  webClient.uuid = makeUuid()
  webClient.character = payloadAsMap["character"].(string)
  webClient.name = payloadAsMap["name"].(string)

  // Register the client with the hub
  webClient.hub.register <- webClient
}


/**
 * Loops over websocket messages and pumps them into the incomingEvents
 * channel.
 *
 * Websocket messages are expected to be valid JSON with a key "eventType" that
 * maps to a string. Otherwise, they are dropped.
 */
func (webClient *WebClient) pumpWebsocketMessages() {
  for {
    // get next message
    _, message, err := webClient.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
    }

    log.Printf("Got websocket message ", string(message))

    // Parse the message
    var jsonMap map[string]interface{}

    err = json.Unmarshal(message, &jsonMap)

    if err != nil {
      log.Printf("error umarshalling", err)
      continue
    }

    eventType, ok := jsonMap["EventType"].(string)

    if ! ok {
      log.Printf("error: bad JSON (event Type not string)");
      continue
    }

    // create event for message
    webClient.incomingEvents <- &WebClientIncomingEvent{
      eventType: eventType,
      source: "client",
      payload: jsonMap,
    }
  }
}
