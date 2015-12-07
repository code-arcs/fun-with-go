package codearcs

type Game struct {
	rolls [21]int
	current int
}

func (g *Game) Roll (pins int) {
	g.rolls[g.current] = pins
	g.current++
}

func (g *Game) GetScore () int {
	result := 0
	i := 0
	for frame := 0; frame < 10; frame++ {
		if g.isStrike(i) {
			result += 10 + g.strikeBonus(i)
			i++
		} else if g.isSpare(i) {
			result += 10 + g.spareBonus(i)
			i += 2
		} else {
			result += g.frameScore(i)
			i += 2
		}
	}
	return result
}

func (g *Game) isStrike(i int) bool {
	return (g.rolls[i] == 10)
}

func (g *Game) isSpare(i int) bool {
	return (g.rolls[i]+g.rolls[i+1] == 10)
}

func (g *Game) spareBonus(i int) int {
	return g.rolls[i+2]
}

func (g *Game) strikeBonus(i int) int {
	return g.rolls[i+1] + g.rolls[i+2]
}

func (g *Game) frameScore(i int) int {
	return g.rolls[i] + g.rolls[i+1]
}