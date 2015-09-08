package connect4

import (
	"db"
	"log"
)

var (
	client = db.GetInstance()
)

type Game struct {
	gameId     string
	gamerOneId string
	gamerTwoId string
	lastPlayer bool
	state      [7][6]byte
}

func StartGame() (game Game, err error) {

}

func LoadGame(gameId string) (game Game, err error) {

}

func (game Game) Save() (err error) {

}

func (game *Game) MakeMove(gamerId string, column int) (err error) {
	//isThereWinnerAt
}

func (game Game) isThereWinnerAt(x int, y int) {

}
