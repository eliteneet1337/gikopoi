/**
 * Represents a room and stores the state of its contents in the world.
 *
 * This class is utilized directly (and only) by the hub. In this code, users
 * are referred to as "entities" (bots may later become entities).
 */
package main


/**
 * Room Struct definition
 */
type Room struct {
  hub *Hub                               // pointer to hub
  uuid string                            // UUID
  name string                            // Visible name of the room
  entities []string                      // array of UUIDs, web clients or bots
  entityCoords map[string](*EntityCoordinates)
}

type EntityCoordinates struct {
  x int
  y int
}


/**
 * Initialize a default room.
 *
 * Returns a struct with the default fields initialized.
 */
func newDefaultRoom(hub *Hub) *Room {
  room := &Room{
    hub: hub,
    uuid: makeUuid(),
    name: "Developer's Lounge",
    entities: make([]string, 0),
    entityCoords: make(map[string](*EntityCoordinates)),
  }

  return room
}


/**
 * Add a new entity to the room by UUID.
 */
func (room *Room) addEntity(uuid string) {
  room.entities = append(room.entities, uuid);
  room.entityCoords[uuid] = &EntityCoordinates{x: 0, y: 0,}
}


/**
 * Get struct summarizing in the room; used for sending as update to the
 * client.
 */
func (room *Room) getEntityDescriptions() []*ClientMessageEntityDescription {
  descriptions := make([]*ClientMessageEntityDescription, 0)

  for _, entityUuid := range room.entities {
    // Try looking for web client
    webClient := room.hub.webClients[entityUuid]

    if webClient != nil {
      description := &ClientMessageEntityDescription{
        Uuid: webClient.uuid,
        Name: webClient.name,
        Character: webClient.character,
        Coords: &ClientMessageRoomCoordinates{
          X: room.entityCoords[webClient.uuid].x,
          Y: room.entityCoords[webClient.uuid].y,
        },
      }

      descriptions = append(descriptions, description)
    }
  }

  return descriptions
}
