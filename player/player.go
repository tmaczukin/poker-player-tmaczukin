package player

import (
	"github.com/tmaczukin/poker-player-tmaczukin/leanpoker"
)

const VERSION = "20160520203933"

func BetRequest(state *leanpoker.Game) int {
	return 0
}

func Showdown(state *leanpoker.Game) {
}

func Version() string {
	return VERSION
}
