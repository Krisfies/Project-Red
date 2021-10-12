package main

import (
	"fmt"
	"os"
)

func (p *Personnage) AccessInventory(e3, e4, e5, e6 *Monstre, a *Equipement) {
	// fonction qui nous permet d'acceder a notre inventaire
	if len(p.inventory) == 0 {
		Slow("\nLe sac est ", 1)
		fmt.Print(Yellow + "")
		Slow("vide", 1)
		fmt.Print("" + Reset)
	} else {
		fmt.Print("\n")
		for i := 0; i < len(p.inventory); i++ {
			if p.inventory[i] != " " {
				Slow("---] ", 2)
				fmt.Print(""+Yellow, p.inventory[i], Reset+"")
				Slow(" [---", 2)
				fmt.Print("\n")
			}
		}
		fmt.Print("\n")
	}
	Slow("----------------\n", 3)
	fmt.Print(Yellow + "")
	Slow("(1) ", 3)
	fmt.Print("" + Reset)
	Slow("Prendre une ", 3)
	fmt.Print(Yellow + "")
	Slow("Potion de vie\n", 3)
	Slow("(2) ", 3)
	fmt.Print("" + Reset)
	Slow("Prendre une ", 3)
	fmt.Print(Yellow + "")
	Slow("Potion de poison\n", 3)
	Slow("(3) ", 3)
	fmt.Print("" + Reset)
	Slow("Utiliser le Livre de sort : ", 3)
	fmt.Print(Yellow + "")
	Slow("Boule de Feu\n", 3)
	Slow("(4) ", 3)
	fmt.Print("" + Reset)
	Slow("Mettre un ", 3)
	fmt.Print(Yellow + "")
	Slow("Chapeau\n", 3)
	Slow("(5) ", 3)
	fmt.Print("" + Reset)
	Slow("Mettre une ", 3)
	fmt.Print(Yellow + "")
	Slow("Tunique\n", 3)
	Slow("(6) ", 3)
	fmt.Print("" + Reset)
	Slow("Mettre des ", 3)
	fmt.Print(Yellow + "")
	Slow("Bottes\n", 3)
	Slow("(0) ", 3)
	fmt.Print("" + Reset)
	Slow("Arrêter de regarder dans le ", 3)
	fmt.Print(Yellow + "")
	Slow("Sac\n", 3)
	fmt.Print("" + Reset)
	Slow("----------------\n", 3)
	p.UseInventory(e3, e4, e5, e6, a)
}

func (p *Personnage) SuperAccessInventory(e3, e4, e5, e6 *Monstre, a *Equipement) {
	// fonction qui nous permet d'acceder a notre inventaire
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	if len(p.inventory) == 0 {
		Slow("\nLe sac est ", 1)
		fmt.Print(Yellow + "")
		Slow("vide", 1)
		fmt.Print("" + Reset)
	} else {
		fmt.Print("\n")
		for i := 0; i < len(p.inventory); i++ {
			if p.inventory[i] != " " {
				fmt.Print("---] ")
				fmt.Print(""+Yellow, p.inventory[i], Reset+"")
				fmt.Println(" [---")
			}
		}
		fmt.Print("\n----------------\n")
	}
	fmt.Println(Yellow+"(1)"+Reset, "Prendre une", Yellow+"Potion de vie")
	fmt.Println("(2)"+Reset, "Prendre une", Yellow+"Potion de poison")
	fmt.Println("(3)"+Reset, "Utiliser le Livre de sort :", Yellow+"Boule de Feu")
	fmt.Println("(4)"+Reset, "Mettre un", Yellow+"Chapeau")
	fmt.Println("(5)"+Reset, "Mettre une", Yellow+"Tunique")
	fmt.Println("(6)"+Reset, "Mettre des", Yellow+"Bottes")
	fmt.Println("(0)"+Reset, "Arrêter de regarder dans le", Yellow+"Sac"+Reset, "\n----------------")
	p.UseInventory(e3, e4, e5, e6, a)
}

func (p *Personnage) UseInventory(e3 , e4 , e5 , e6 *Monstre, a *Equipement) {
	// fonction qui nous permet d'interagir avec les éléments de l'inventaire
	var use int
	fmt.Scanln(&use)
	switch use {
	case 1:
		p.TakePot()
		p.SuperAccessInventory(e3, e4, e5, e6, a)
	case 2:
		p.PoisonPot()
		p.SuperAccessInventory(e3, e4, e5, e6, a)
	case 3:
		p.Spellbook(e3, e4, e5, e6, a)
		p.SuperAccessInventory(e3, e4, e5, e6, a)
	case 4:
		if p.Checkinv("Chapeau de l'aventurier") {
			if a.EquipLpBonus("chapeau", p) {
				p.lp +=15
				p.lpmax+= 15
				Slow("Vous avez désormais +15 points de vies maximum avec le couvre chef !\n", 2)
			}
			p.RemoveInventory("Chapeau de l'Aventurier")
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		} else if !p.Checkinv("Chapeau de L'aventurier") {
			Slow("Tu l'as déja équiper !", 2)
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		}
	case 5:
		if p.Checkinv("Tunique de l'Aventurier") {
			if a.EquipLpBonus("tunique", p) {
				p.lp +=20
				p.lpmax+= 20
				Slow("Vous avez désormais +20 points de vies maximum grâce a la tunique !\n", 2)
			}
			p.RemoveInventory("Tunique de l'Aventurier")
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		} else if !p.Checkinv("Tunique de L'aventurier"){
			Slow("Tu l'as déja équiper !", 2)
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		}
	case 6:
		if p.Checkinv("Bottes de l'Aventurier") {
			if a.EquipLpBonus("bottes",p) {
				p.lp +=10
				p.lpmax+= 10
				Slow("Vous avez désormais +10 points de vies maximum grâce aux bottes !\n", 2)
			}
			p.RemoveInventory("Bottes de l'Aventurier")
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		} else if !p.Checkinv("Bottes de L'aventurier"){
			Slow("Tu l'as déja équiper !", 2)
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		}
	case 0:
		p.menu(e3, e4, e5, e6, a)
	}
}

func (p *Personnage) Spellbook(e3, e4, e5, e6 *Monstre, a *Equipement) {
	// fonction qui nous permet d'ajouter ou repertorier les sorts (spell)
	for _, letter := range p.skill {
		if letter == ("Boule de feu") {
			Slow("Tu connais déjà ce ", 2)
			fmt.Print(Yellow + "")
			Slow("sort", 2)
			fmt.Print("" + Reset)
		} else {
			p.skill = append(p.skill, "Boule de feu")
		}
	}
	p.menu(e3, e4, e5, e6, a)
}
