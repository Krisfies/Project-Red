package main

import (
	"math/rand"
)

func (p *Personnage) Loot() {
	loot := rand.Intn(3)
	if loot == 0 {
		Slow("\nVous ne remportez rien de plus",2)
	} else if loot == 1 {
		p.money += 25
		Slow("\nVous remportez ",2)
		Slow(Yellow+"25 pièces "+Reset,2)
		Slow("supplémentaires",2)
	} else if loot == 2 {
		p.exp += 15
		Slow("\nVous remportez ",2)
		Slow(Blue+"15 exp "+Reset,2)
		Slow("supplémentaires",2)
	}
}