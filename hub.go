/**
 * Master object controlling the server as a world and list of clients. Accepts
 * new events through multiple channels.
 */
package main

import (
  "log"
)

type Hub struct {
  webClients []*WebClient
  register chan *WebClient
}


func newHub() *Hub {
  return &Hub{
    webClients: make([]*WebClient, 0),
  }
}

func (hub *Hub) start() {
  for {
    select {
    case webClient := <- hub.register:
      hub.webClients = append(hub.webClients, webClient)
      log.Println("new client registered");
    }
  }
}
