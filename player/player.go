package player

import "github.com/lean-poker/poker-player-go/leanpoker"

const VERSION = "20160520200234"

func BetRequest(state *leanpoker.Game) int {
	return 0
}

func Showdown(state *leanpoker.Game) {

}

func Version() string {
	return VERSION
}
