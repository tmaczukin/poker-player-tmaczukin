package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"player"
)

const PORT = 4711

func main() {
	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatal(err)
	}
}

func handleRequest(w http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		log.Printf("Error parsing form data: %s", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	action := request.FormValue("action")
	log.Printf("Request method=%s url=%s action=%s from client=%s\n", request.Method, request.URL, action, request.RemoteAddr)
	switch action {
	case "bet_request":
		gameState := parseGameState(request.FormValue("game_state"))
		if gameState == nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
		result := player.BetRequest(gameState)
		fmt.Fprintf(w, "%d", result)
		return
	case "showdown":
		gameState := parseGameState(request.FormValue("game_state"))
		if gameState == nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
		player.Showdown(gameState)
		fmt.Fprint(w, "")
		return
	case "version":
		fmt.Fprint(w, player.Version())
		return
	default:
		http.Error(w, "Invalid action", 400)
	}
}

func parseGameState(stateStr string) *player.GameState {
	stateBytes := []byte(stateStr)
	gameState := new(player.GameState)
	if err := json.Unmarshal(stateBytes, &gameState); err != nil {
		log.Printf("Error parsing game state: %s", err)
		return nil
	}
	return gameState
}
