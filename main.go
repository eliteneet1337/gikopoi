/**
 * Defines the webserver in the main() function and provides a wrapper to the
 * hub and client abstractions. Serves static files out of static/.
 */
package main

import (
  "flag"
  "log"
  "net/http"
  "github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
}


func main() {
  // Parse and set command line flags
  var addr = flag.String("addr", ":8080", "http service address")
  flag.Parse()

  // Define hub
  var hub = newHub();
  go hub.start();

  // Define routes
  var fs = http.FileServer(http.Dir("static"))

  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.HandleFunc("/ws", func(response http.ResponseWriter, request *http.Request) {
    handleWs(hub, response, request)
  })

  http.HandleFunc("/", handleIndex)

  // Listen and serve
  err := http.ListenAndServe(*addr, nil)

  if err != nil {
    log.Fatal("Listening error: ", err)
  }
}

func handleIndex(response http.ResponseWriter, request *http.Request) {
  log.Println(request.URL)
  http.ServeFile(response, request, "static/index.html")
}

func handleWs(hub *Hub, response http.ResponseWriter, request *http.Request) {
  // Upgrade connection to websocket
  conn, err := upgrader.Upgrade(response, request, nil)

  if err != nil {
		log.Println(err)
		return
  }

  // Create new webclient and begin handling events
  var webClient = newWebClient(hub, conn);

  go webClient.handleIncomingEvents()

  webClient.pumpWebsocketMessages()
}
