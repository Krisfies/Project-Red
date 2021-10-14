package main

import (
	"math/rand"
)

func (p *Personnage) Loot() {
	loot := rand.Intn(3)
	if loot == 1 {
		Slow("Vous ne remportez rien",2)
	} else if loot == 2 {
		p.money += 25
		Slow("Vous remportez 25 pièces supplémentaires",2)
	} else if loot == 3 {
		p.exp += 15
		Slow("Vous remportez 15 exp supplémentaires",2)
	}
}