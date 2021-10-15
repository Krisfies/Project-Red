package main

import (
	"fmt"

)

func (p *Personnage) Marchand(e3, e4, e5, e6, e7 *Monstre, a *Equipement) {
	// fonction affichant le menu du marchand , et les ajoute a notre inventaire
	var menum int
	Slow("Vous avez "+Yellow, 1)
	fmt.Print(p.money)
	Slow(" pièces\n"+Reset, 1)
	fmt.Scanln(&menum)
	switch menum {
	case 1:
		p.AddInventory("Potion de vie", 15)
		p.Marchand(e3, e4, e5, e6, e7, a)
	case 2:
		p.AddInventory("Potion de poison", 20)
		p.Marchand(e3, e4, e5, e6, e7, a)
	case 3:
		p.AddInventory("Livre de sort: Boule de Feu", 25)
		p.Marchand(e3, e4, e5, e6, e7, a)
	case 4:
		p.AddInventory("Fourrure de Loup", 5)
		p.Marchand(e3, e4, e5, e6, e7, a)
	case 5:
		p.AddInventory("Peau de Troll", 5)
		p.Marchand(e3, e4, e5, e6, e7, a)
	case 6:
		p.AddInventory("Cuir de Sanglier", 4)
		p.Marchand(e3, e4, e5, e6, e7, a)
	case 7:
		p.AddInventory("Plume de Corbac", 6)
		p.Marchand(e3, e4, e5, e6, e7, a)
	case 0:
		p.menu(e3, e4, e5, e6, e7, a)
	}
}
