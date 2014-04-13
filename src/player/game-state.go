package player

type GameState struct {
    small_blind     int
    current_buy_in  int
    pot             int
    minimum_raise   int
    dealer          int
    orbits          int
    in_action       int
    players         []Player
    community_cards []Card
}
           
type Player struct {
    id         int
    name       string
    status     string
    version    string
    stack      int
    bet        int
    hole_cards []Card
}

type Card struct {
    rank string
    suit string
}
