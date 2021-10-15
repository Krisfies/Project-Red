package main

import (
	"fmt"
	"os"
	"time"
)

func (p *Personnage) menu(e3, e4, e5, e6 *Monstre, a *Equipement) {
	var menu int
	var retour int
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	Slow("+++++++++++++++++++++++++++++++\n", 3)
	Slow("Que voulez vous faire ?\n", 3)
	Slow("----- \n", 3)
	if !p.Checkinv("Objet Suspicieux") {
		Slow(Yellow+"(1) "+Reset, 3)
		Slow("Aller voir l'homme en jaune sur le ", 3)
		Slow(Yellow+"banc\n"+Reset, 3)
		Slow("----- \n", 3)
	} else {
		Slow(Yellow+"(1) "+Reset, 3)
		Slow("Aller sur le ", 3)
		Slow(Yellow+"banc jaune\n"+Reset, 3)
		Slow("----- \n", 3)
	}
	Slow(Yellow+"(2) "+Reset, 3)
	Slow("Regarder la carte ", 3)
	Slow(Yellow+"d'Identité\n"+Reset, 3)
	Slow("----- \n", 3)
	Slow(Yellow+"(3) "+Reset, 3)
	Slow("Jeter un oeil dans le ", 3)
	Slow(Yellow+"Sac\n"+Reset, 3)
	Slow("----- \n", 3)
	Slow(Yellow+"(4) "+Reset, 3)
	Slow("Se rendre dans la boutique du ", 3)
	Slow(Yellow+"Marchand\n"+Reset, 3)
	Slow("----- \n", 3)
	Slow(Yellow+"(5) "+Reset, 3)
	Slow("Aller à la ", 3)
	Slow(Yellow+"Forge\n"+Reset, 3)
	Slow("----- \n", 3)
	Slow(Yellow+"(6) "+Reset, 3)
	Slow("Partir en exploration dans le ", 3)
	Slow(Yellow+"Donjon\n"+Reset, 3)
	Slow(Yellow+"(7) "+Reset, 3)
	Slow("Affronter les ", 3)
	Slow(Yellow+"Quatres Grands\n"+Reset, 3)
	Slow("----- \n", 3)
	Slow(Yellow+"(0) "+Reset, 3)
	Slow("Quitter la ", 3)
	Slow(Yellow+"Simulation"+Reset, 3)
	Slow("\n-----\n", 3)
	Slow("Entrez le numéro de l'option:\n", 3)
	Slow("+++++++++++++++++++++++++++++++\n", 3)
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		if !p.Checkinv("Objet Suspicieux") {
			p.QuestMan(e3, e4, e5, e6, a)
		} else {
			Slow("Il n'y a ", 5)
			Slow(Yellow+"personne "+Reset, 5)
			Slow("ici", 5)
			time.Sleep(2 * time.Second)
		}
		p.menu(e3, e4, e5, e6, a)
	case 2:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.DisplayInfo(a)
		Slow(Yellow+"(0) "+Reset, 2)
		Slow("Arrêter de regarder la ", 2)
		Slow(Yellow+"carte\n"+Reset, 2)
		fmt.Scanln(&retour)
		switch retour {
		case 0:
			p.menu(e3, e4, e5, e6, a)
		}
	case 3:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.AccessInventory(e3, e4, e5, e6, a)
	case 4:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		Slow("+++++++++++++++++++++Marchand+++++++++++++++++++\n", 2)
		Slow("-----\n", 3)
		Slow(Yellow+"(1) "+Reset,3)
		Slow("Potion de vie ", 3)
		Slow(Yellow+" --> "+Reset, 3)
		Slow("15 pièces ", 3)
		Slow(Yellow+" <-- \n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(2) "+Reset,3)
		Slow("Potion de poison", 3)
		Slow(Yellow+" --> "+Reset, 3)
		Slow("20 pièces", 3)
		Slow(Yellow+" <-- \n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(3) "+Reset,3)
		Slow("Livre de sort: Boule de Feu", 3)
		Slow(Yellow+" --> "+Reset, 3)
		Slow("25 pièces", 3)
		Slow(Yellow+" <-- \n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(4) "+Reset,3)
		Slow("Fourrure de Loup", 3)
		Slow(Yellow+" --> "+Reset, 3)
		Slow("5 pièces", 3)
		Slow(Yellow+" <-- \n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(5) "+Reset,3)
		Slow("Peau de Troll", 3)
		Slow(Yellow+" --> "+Reset, 3)
		Slow("6 pièces", 3)
		Slow(Yellow+" <-- \n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(6) "+Reset,3)
		Slow("Cuir de Sanglier", 3)
		Slow(Yellow+" --> "+Reset, 3)
		Slow("4 pièces", 3)
		Slow(Yellow+" <-- \n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(7) "+Reset,3)
		Slow("Plume de Corbac", 3)
		Slow(Yellow+" --> "+Reset, 3)
		Slow("6 pièces", 3)
		Slow(Yellow+" <-- \n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(0) "+Reset, 3)
		Slow("Retourner à la ", 3)
		Slow(Yellow+"Place Centrale"+Reset, 3)
		Slow("\n-----", 3)
		Slow("\n++++++++++++++++++++++++++++++++++++++++++++++++\n", 3)
		p.Marchand(e3, e4, e5, e6, a)
	case 5:
		p.Forgeron(e3, e4, e5, e6, a)
	case 6:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.TrainingFight(a)
		p.menu(e3, e4, e5, e6, a)
	case 7:
		p.TheFirst(e3, e4, e5, e6, a)
		p.menu(e3, e4, e5, e6, a)
	case 0:
		fmt.Println("Fin de la transmission")
		break
	case 7171:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		Slow("Bravo", 1)
		fmt.Printf(p.name)
		Slow("tu as réussi à triompher des quatres Grands", 1)
		fmt.Printf("\n")
		Slow("Ta quête s'arrête ici l'ami, mais nos voyagent se recroiseront, soit en sur.", 1)
	}
}
