package main

import (
	"fmt"
)

func (p *Personnage) CharTurn(m *Monstre, a *Equipement) {
	var choice int
	Slow("\nC'est au joueur !", 1)
	Slow("\n\nIl reste "+Yellow,1)
	fmt.Print(m.lp)
	Slow("/",1)
	fmt.Print(m.lpmax)
	Slow(Reset+" PV à "+Yellow,1)
	Slow(m.name,1)
	Slow(Reset+"\nIl reste "+Yellow,1)
	fmt.Print(p.lp)
	Slow("/",1)
	fmt.Print(p.lpmax)
	Slow(Reset+" PV à "+Yellow,1)
	Slow(p.name,1)
	Slow("\n\n(1) "+Reset, 1)
	Slow("Attaquer", 1)
	Slow(Yellow+"\n(2) "+Reset, 1)
	Slow("Utiliser un objet\n", 1)
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		m.lp -= p.damage
		fmt.Print(Yellow + "")
		Slow(p.name, 1)
		Slow(Reset+" utilise ", 1)
		Slow(Yellow+"Coup de poing "+Reset, 1)
		Slow(" et inflige "+Red, 1)
		fmt.Print(p.damage)
		Slow(" points de dégâts"+Reset, 1)
		Slow(" à "+Yellow,1)
		Slow(m.name, 1)
		Slow(Reset+" il lui reste "+Yellow, 1)
		fmt.Print(m.lp)
		Slow(" points de vies\n"+Reset, 1)
	case 2:
		if len(p.inventory) == 0 {
			Slow("\nLe sac est ", 1)
			Slow(Yellow+"vide", 1)
			Slow("\n(0) ", 1)
			Slow("Retour"+Reset, 1)
		} else {
			fmt.Print("\n")
			for i := 0; i < len(p.inventory); i++ {
				if p.inventory[i] != " " {
					Slow("---]", 2)
					fmt.Print(""+Yellow, p.inventory[i], Reset+"")
					Slow("[---\n", 2)
				}
			}
			Slow(Yellow+"\n(1) "+Reset, 1)
			Slow("Prendre une ", 1)
			Slow(Yellow+"Potion de vie", 1)
			Slow("\n(2) "+Reset, 1)
			Slow("Prendre une ", 1)
			Slow(Yellow+"Potion de poison", 1)
			Slow("\n(0) ", 1)
			Slow("Retour\n"+Reset, 1)
		}

		var use int
		fmt.Scanln(&use)
		switch use {
		case 1:
			p.TakePot()
		case 2:
			p.PoisonPot(a)
		case 0:
			p.CharTurn(m, a)
		}
	}
}
