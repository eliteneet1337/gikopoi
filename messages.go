/**
 * This file holds struct definitions for messages that are
 */
package main

/**
 * Outgoing message to the client.
 * Used by the hub to pipe outgoing messages to the WebClient.
 */
type ClientMessage struct {
  EventType string
  Payload map[string]interface{}
}


/**
 * Outgoing message to respond to client initialization.
 *
 * This response
 */
type ClientMessageInitResponse struct {
  RoomUuid string
  RoomName string
  Entities []*ClientMessageEntityDescription
}


/**
 * Component of message responses: holds description of an entity
 */
type ClientMessageEntityDescription struct {
  Uuid string
  Name string
  Character string
  Coords *ClientMessageRoomCoordinates
}


/**
 * Component of message responses: hold coordinates of an object in a room
 */
 type ClientMessageRoomCoordinates struct {
   X int
   Y int
 }
