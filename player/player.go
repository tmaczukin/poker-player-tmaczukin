package player

import (
	"github.com/lean-poker/poker-player-go/leanpoker"

	"log"
)

const VERSION = "20160520203349"

func BetRequest(state *leanpoker.Game) int {
	log.Printf("BetRequest: %s\n", state)
	return 0
}

func Showdown(state *leanpoker.Game) {
	log.Printf("Showdown: %s\n", state)
}

func Version() string {
	log.Printf("VERSION: %s\n", VERSION)
	return VERSION
}
