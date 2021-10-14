package main

import (
	"fmt"
)

func (p *Personnage) DisplayInfo(a * Equipement) {
	// fonction nous permettant de voir les informations de notre personnage
	Slow("-------------------------\n", 2)
	fmt.Print(Yellow + "")
	Slow("Nom: ", 2)
	fmt.Print("" + Reset)
	Slow(p.name, 2)
	fmt.Print(Yellow + "")
	Slow("\nClasse: ", 2)
	fmt.Print("" + Reset)
	Slow(p.class, 2)
	fmt.Print(Yellow + "")
	Slow("\nNiveau: ", 2)
	fmt.Print("" + Reset)
	fmt.Print(p.level)
	fmt.Print(Yellow + "")
	Slow("\nVie actuelle: ", 2)
	fmt.Print("" + Reset)
	fmt.Print(p.lp)
	fmt.Print(Yellow + "")
	Slow("\nVie maximum: ", 2)
	fmt.Print("" + Reset)
	fmt.Print(p.lpmax)
	fmt.Print(Yellow + "")
	Slow("\nContenu de l'inventaire: ", 2)
	fmt.Print("" + Reset)
	fmt.Print(p.inventory)
	fmt.Print(Yellow + "")
	Slow("\nCompétences: ", 2)
	fmt.Print("" + Reset)
	fmt.Print(p.skill)
	fmt.Print(Yellow + "")
	Slow("\nPièces: ", 2)
	fmt.Print("" + Reset)
	fmt.Println(p.money)
	Slow(Yellow+"\nChapeau: "+Reset,2)
	Slow(a.Chapeau,2)
	Slow(Yellow+"\nTunique: "+Reset,2)
	Slow(a.Tunique,2)
	Slow(Yellow+"\nBottes: "+Reset,2)
	Slow(a.Bottes,2)
	if a.Arme == "Véritable Couteau" {
		Slow(Yellow+"\nArme:"+Reset,2)
		Slow(a.Arme,2)
	}
	Slow("\n-------------------------\n", 2)
}
