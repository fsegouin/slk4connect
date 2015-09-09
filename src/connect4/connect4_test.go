package connect4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	aGameId         = "abcdef"
	aFirstPlayerId  = "toto"
	aSecondPlayerId = "tata"
)

func TestStartGame(t *testing.T) {
	game, error := StartGame(aGameId, aFirstPlayerId)

	assert.Nil(t, error)
	assert.NotNil(t, game)
	assert.Equal(t, len(game.State), 7, "Table should have 7 columns")
	assert.Equal(t, len(game.State[0]), 6, "Table should have 6 columns")

	for i := range game.State {
		for j := range game.State[i] {
			assert.Equal(t, game.State[i][j], uint8(0), "Table should be 0 initialized")
		}
	}
}
