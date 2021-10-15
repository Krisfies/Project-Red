package main

import (
	"fmt"
)

func (p *Personnage) DisplayInfo(a * Equipement) {
	// fonction nous permettant de voir les informations de notre personnage
	Slow("-------------------------\n", 2)
	Slow(Yellow+"Nom: "+Reset, 2)
	Slow(p.name, 2)
	Slow(Yellow+"\nClasse: "+Reset, 2)
	Slow(p.class, 2)
	Slow(Yellow+"\nNiveau: "+Reset, 2)
	fmt.Print(p.level)
	Slow(Yellow+"\nExpérience: "+Reset,2)
	fmt.Print(p.exp, "/", p.expmax)
	Slow(Yellow+"\nVie actuelle: "+Reset, 2)
	fmt.Print(p.lp)
	Slow(Yellow+"\nVie maximum: "+Reset, 2)
	fmt.Print(p.lpmax)
	Slow(Yellow+"\nContenu de l'inventaire: "+Reset, 2)
	fmt.Print(p.inventory)
	Slow(Yellow+"\nCompétences: "+Reset, 2)
	fmt.Print(p.skill)
	Slow(Yellow+"\nPièces: "+Reset, 2)
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
