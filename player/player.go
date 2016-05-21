package player

import (
	"github.com/lean-poker/poker-player-go/leanpoker"

	"fmt"
)

const VERSION = "20160520202941"

func BetRequest(state *leanpoker.Game) int {
	fmt.Printf("BetRequest: %s\n", state)
	return 0
}

func Showdown(state *leanpoker.Game) {
	fmt.Printf("Showdown: %s\n", state)
}

func Version() string {
	fmt.Printf("VERSION: %s\n", VERSION)
	return VERSION
}
