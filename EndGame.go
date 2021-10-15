package main

import (
	"os"
)

func (p *Personnage) EndGame() {
	Slow("\nFélicitations",5)
	Slow(p.name,5)
	Slow("\n Tu as terminé le jeu.",5)
	os.Exit(7171)
}