/**
 * The hub is the master object controlling the server as a world and of users
 * and rooms.
 *
 * The hub accepts new events through clients thorugh its multiple channels,
 * each channel describing a different type of event. For example, see
 * `hub.register`.
 */
package main

import (
  "github.com/fatih/structs"
)


/**
 * Hub struct definition
 */
type Hub struct {
  webClients map[string]*WebClient          // maps UUID to web client
  rooms map[string]*Room                    // maps UUID to room
  defaultRoomUuid string                    // UUID of default room
  register chan *WebClient                  // Handles new registrations
}


/**
 * Initialize a new hub.
 *
 * Returns a struct with the default fields initialized. Some fields are not
 * initialized until `hub.start().
 */
func newHub() *Hub {
  return &Hub{
    webClients: make(map[string]*WebClient),
    rooms: make(map[string]*Room),
    defaultRoomUuid: "",
    register: make(chan *WebClient),
  }
}


/**
 * Runs the hub in a seperate 'go' thread.
 *
 * Loops over selecting new events from hub channels and responds accordingly.
 * This method is a stub. More hub event channels will be added later.
 */
func (hub *Hub) start() {
  // Initialize rooms, bots, etc
  hub.initializeRooms()

  // Loop over events from hub channels
  for {
    select {
    case webClient := <- hub.register:
      hub.handleRegister(webClient)
    }
  }
}


/**
 * Handles an event on the register channel.
 *
 * Adds the new web client to the default room and sends an outgoing message
 * to their browser describing the other members of the room.
 */
func (hub *Hub) handleRegister(webClient *WebClient) {
  // Add the user to the default room
  hub.webClients[webClient.uuid] = webClient
  room := hub.rooms[hub.defaultRoomUuid]
  room.addEntity(webClient.uuid)

  // Build and response
  payload := &ClientMessageInitResponse{
    RoomUuid: room.uuid,
    RoomName: room.name,
    Entities: room.getEntityDescriptions(),
  }

  webClient.incomingEvents <- &WebClientIncomingEvent{
    eventType: "clientResp",
    source: "hub",
    payload: structs.Map(payload),
  }
}


/**
 * Initialize the rooms in the system.
 *
 * This method is a stub. In the future it will load rooms from descriptor
 * files on the file system.
 */
func (hub *Hub) initializeRooms() {
  defaultRoom := newDefaultRoom(hub)
  hub.rooms[defaultRoom.uuid] = defaultRoom
  hub.defaultRoomUuid = defaultRoom.uuid
}
