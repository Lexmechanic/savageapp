package encounter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Upgrader is used to upgrade HTTP connections to WebSocket connections.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan []byte)            // Broadcast channel
var mutex = &sync.Mutex{}                    // Protect clients map

func HandleMessage(message []byte) {
	// Parse JSON
	var input map[string]any
	encounter := GetEncounter()

	err := json.Unmarshal(message, &input)
	if err != nil {
		fmt.Println("Error Unmarshalling WS message:", err)
	}

	//If action is nil then don't do anything. Likely a badly structured send
	if input["action"] == nil {
		return
	}

	//Perform action
	switch input["action"].(string) {
	case "add-bestiary":
		if input["id"] != nil {
			encounter.AddBestiary(int(input["id"].(float64)))
		}
	case "delete-bestiary":
		if input["id"] != nil {
			encounter.DeleteBestiary(int(input["id"].(float64)))
		}
	case "update-bestiary":
		if input["encounter_bestiary"] != nil {
			var entry EncounterBestiary
			entryData, err := json.Marshal(input["encounter_bestiary"])
			if err != nil {
				fmt.Println("Failed to marshal encounter_bestiary:", err)
				return
			}
			err = json.Unmarshal(entryData, &entry)
			if err != nil {
				fmt.Println("Failed to unmarshal update-bestiary entry:", err)
				return
			} else {
				encounter.UpdateBestiary(entry)
			}
		}
	case "end-encounter":
		encounter.EndEncounter()
	case "start-encounter":
		encounter.StartEncounter()
	case "next-round":
		encounter.NextRound()
	case "draw-card":
		if input["id"] != nil {
			id := int(input["id"].(float64))
			if id >= 0 && id < len(encounter.BestiaryList) {
				encounter.performDrawEncounterId(id)
			} else {
				fmt.Println("Invalid ID for drawing card:", id)
			}
		}
	case "select-card":
		if input["id"] != nil && input["card_id"] != nil {
			id := int(input["id"].(float64))
			cardId := int(input["card_id"].(float64))
			encounter.SelectCard(id, cardId)
		}
	case "get-encounter":

	}

	data, err := json.Marshal(encounter)
	if err != nil {
		fmt.Println("Error Marshaling encounter message:", err)
	}

	broadcast <- data
}

func HandleNewConnection(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	// Listen for incoming messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}
		HandleMessage(message)
	}
}

func handleBroadCast() {
	for {
		// Grab the next message from the broadcast channel
		message := <-broadcast

		// Send the message to all connected clients
		mutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
