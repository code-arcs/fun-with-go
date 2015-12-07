package codearcs

import "testing"

func rollMany(g *Game, rolls int, pins int) {
	for i := 1; i <= rolls; i++ {
		g.Roll(pins)
	}
}

func TestGutterGame(t *testing.T) {
	game := Game{}

	rollMany(&game, 0, 20);

	if game.GetScore() != 0 {
		t.Error("Score was incorrect", 0, game.GetScore())
	}
}

func TestOnePinGame(t *testing.T) {
	game := Game{}

	rollMany(&game, 1, 20);

	if game.GetScore() != 20 {
		t.Error("Score was incorrect", 20, game.GetScore())
	}
}

func TestAddSpareBonus(t *testing.T) {
	game := Game{}

	game.Roll(4)
	game.Roll(6)
	rollMany(&game, 18, 4);

	var expected = 4 + 6 + 4 + 18 * 4
	if game.GetScore() != expected {
		t.Error("Score was incorrect",  expected, game.GetScore())
	}
}

func TestAddStrikeBonus(t *testing.T) {
	game := Game{}

	game.Roll(10)
	rollMany(&game, 18, 4);

	var expected = 10 + 4 + 4 + 18 * 4
	if game.GetScore() != expected {
		t.Error("Score was incorrect",  expected, game.GetScore())
	}
}

func TestPerfectGame(t *testing.T) {
	game := Game{}

	rollMany(&game, 12, 10);

	var expected = 300
	if game.GetScore() != expected {
		t.Error("Score was incorrect",  expected, game.GetScore())
	}
}
