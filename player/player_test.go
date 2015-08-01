package player

import "testing"

func mockGameState() *GameState {
    state := new(GameState)
    // TODO: populate state
    return state
}

func TestBetRequest(t *testing.T) {
    state := mockGameState()
    result := BetRequest(state)
    if result < 0 {
        t.Errorf("Expected BetRequest to return positive integer, instead saw %d", result)
    }
}

func BenchmarkBetRequest(b *testing.B) {
    state := mockGameState()
    for i := 0; i < b.N; i++ {
        BetRequest(state)
    }
}