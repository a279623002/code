package memento

import "testing"

func TestGame_Status(t *testing.T) {
	game := &Game{
		hp: 10,
		mp: 10,
	}
	game.Status()
	progress := game.Save()

	game.Play(-2, -5)
	game.Status()

	game.Load(progress)
	game.Status()
}
