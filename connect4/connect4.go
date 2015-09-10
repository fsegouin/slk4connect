package connect4

import (
	"errors"
	"github.com/fsegouin/slk4connect/db"
	"strconv"
	"strings"
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

	var state string
	for i := range game.State {
		for j := range game.State[i] {
			state += strconv.Itoa(int(game.State[i][j]))
		}
	}

	gameHash := map[string]string{
		"GameId":     game.GameId,
		"GamerOneId": game.GamerOneId,
		"LastPlayer": "false",
		"State":      state,
	}

	client.Cmd("HMSET", gameId, gameHash)

	return game, err
}

func LoadGame(gameId string) (game Game, err error) {
	gameHash, redisErr := client.Cmd("HGETALL", gameId).Hash()

	if redisErr == nil {
		lastPlayer, _ := strconv.ParseBool(gameHash["LastPlayer"])

		var stateMap = strings.Split(gameHash["State"], "")
		var state = make([][]uint8, 7)
		for i := range state {
			state[i] = make([]uint8, 6)
			for j := range state[i] {
				val, _ := strconv.ParseUint(stateMap[i+j*7], 10, 8)
				state[i][j] = uint8(val)
			}
		}

		game := Game{
			GameId:     gameHash["GameId"],
			GamerOneId: gameHash["GamerOneId"],
			GamerTwoId: gameHash["GamerTwoId"],
			LastPlayer: lastPlayer,
			State:      state,
		}

		return game, nil
	}

	return Game{}, errors.New("Cannot find game")
}

//func (game Game) Save() (err error) {
//}

//func (game *Game) MakeMove(gamerId string, column int) (err error) {
//}

//func (game Game) isThereWinnerAt(x int, y int) {
//}
