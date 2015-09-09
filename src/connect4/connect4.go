package connect4

import (
	"bytes"
	"db"
)

var (
	client = db.GetInstance()
)

type Game struct {
	GameId     string
	GamerOneId string
	GamerTwoId string
	LastPlayer bool
	State      [][]uint8
}

func StartGame(gameId, gamerOneId string) (game Game, err error) {
	game = Game{GameId: gameId, GamerOneId: gamerOneId}

	game.State = make([][]uint8, 7)

	for i := range game.State {
		game.State[i] = make([]uint8, 6)
	}

	gameHash := map[string]string{
		"GameId":     game.GameId,
		"GamerOneId": game.GamerOneId,
		"LastPlayer": "false",
		"State":      string(bytes.Join(game.State, []byte{})),
	}

	client.Cmd("HMSET", gameId, gameHash)

	return game, err
}

//func LoadGame(gameId string) (game Game, err error) {
//}

//func (game Game) Save() (err error) {
//}

//func (game *Game) MakeMove(gamerId string, column int) (err error) {
//}

//func (game Game) isThereWinnerAt(x int, y int) {
//}
