/**
 * Holds definitions for the WebClient struct which represents a client that
 * has connected through the web. Manages layer between the hub and the user
 * via websockets.
 */
package main

import (
  "log"
  "encoding/json"
  "github.com/gorilla/websocket"
)

type WebClient struct {
  username string                            // initializes to ""
  hub *Hub
  conn *websocket.Conn
  incomingEvents chan *WebClientIncomingEvent
}

type WebClientIncomingEvent struct {
  eventType string
  payload map[string]interface{}
}


/**
 * Initialize a new webclient with its master hub and connection socket.
 */
func newWebClient(hub *Hub, conn *websocket.Conn) *WebClient {
  return &WebClient{
    username: "",
    hub: hub,
    conn: conn,
    incomingEvents: make(chan *WebClientIncomingEvent),
  }
}

/**
 * Process the incoming events for this client. These may be written by the hub
 * or through pumpWebsocketMessages().
 */
func (webClient *WebClient) handleIncomingEvents() {
  for {
    incomingEvent := <- webClient.incomingEvents
    log.Println("client got event", incomingEvent.eventType,
      incomingEvent.payload);
  }
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

    // Parse the message
    var jsonMap map[string]interface{}

    err = json.Unmarshal(message, &jsonMap)

    if err != nil {
      log.Fatal("error umarshalling", err)
      continue
    }

    eventType, ok := jsonMap["eventType"].(string)

    if ! ok {
      log.Printf("error: bad JSON (event Type not string)");
      continue
    }

    // create event for message
    webClient.incomingEvents <- &WebClientIncomingEvent{
      eventType: eventType,
      payload: jsonMap,
    }
  }
}
