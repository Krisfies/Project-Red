package main

import (
	"math/rand"
)

func (p *Personnage) Loot() {
	loot := rand.Intn(3)
	if loot == 0 {
		Slow("\nVous ne remportez rien",2)
	} else if loot == 1 {
		p.money += 25
		Slow("\nVous remportez 25 pièces supplémentaires",2)
	} else if loot == 2 {
		p.exp += 15
		Slow("\nVous remportez 15 exp supplémentaires",2)
	}
}