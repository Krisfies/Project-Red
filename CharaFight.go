package main

import (
	"fmt"
)

func (p *Personnage) CharTurn(m *Monstre) {
	var choice int
	fmt.Print(Yellow + "")
	Slow("\n(1) ", 2)
	fmt.Print("" + Reset)
	Slow("Attaquer", 2)
	fmt.Print(Yellow + "")
	Slow("\n(2) ", 2)
	fmt.Print("" + Reset)
	Slow("Utiliser un objet\n", 2)
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		m.lp -= p.damage
		fmt.Print(Yellow + "")
		Slow(p.name, 2)
		fmt.Print("" + Reset)
		Slow(" utilise ", 2)
		fmt.Print(Yellow + "")
		Slow("Coup de poing ", 2)
		fmt.Print("" + Reset)
		Slow(" et inflige ", 2)
		fmt.Print(Red + "")
		fmt.Print(p.damage)
		Slow(" points de dégâts", 2)
		fmt.Print("" + Reset)
		Slow(" à ",2)
		Slow(Yellow + "",2)
		Slow(m.name, 2)
		fmt.Print("" + Reset)
		Slow(" il lui reste ", 2)
		fmt.Print(Yellow + "")
		fmt.Print(m.lp)
		Slow(" points de vies\n", 2)
		fmt.Print("" + Reset)
	case 2:
		if len(p.inventory) == 0 {
			Slow("\nLe sac est ", 2)
			Slow(Yellow+"vide", 2)
			Slow("\n(0) ", 2)
			Slow("Retour"+Reset, 2)
		} else {
			fmt.Print("\n")
			for i := 0; i < len(p.inventory); i++ {
				if p.inventory[i] != " " {
					Slow("---]", 2)
					fmt.Print(""+Yellow, p.inventory[i], Reset+"")
					Slow("[---\n", 2)
				}
			}
			Slow(Yellow+"\n(1) "+Reset, 2)
			Slow("Prendre une ", 2)
			Slow(Yellow+"Potion de vie", 2)
			Slow("\n(2) "+Reset, 2)
			Slow("Prendre une ", 2)
			Slow(Yellow+"Potion de poison", 2)
			Slow("\n(0) ", 2)
			Slow("Retour\n"+Reset, 2)
		}

		var use int
		fmt.Scanln(&use)
		switch use {
		case 1:
			p.TakePot()
		case 2:
			p.PoisonPot()
		case 0:
			p.CharTurn(m)
		}
	}
}
