package main

import (
	"os"
	"fmt"
	"time"
)

func (p *Personnage) Forgeron(e3, e4, e5, e6 *Monstre, a *Equipement) {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	Slow("------ Vous entrez dans la ", 3)
	fmt.Print(Yellow + "")
	Slow("Forge ", 3)
	fmt.Print("" + Reset)
	Slow("------", 3)
	fmt.Print("\n \n")
	Slow("-- ", 3)
	fmt.Print(Yellow + "")
	Slow("(1) ", 3)
	fmt.Print("" + Reset)
	Slow("Construire un ", 3)
	fmt.Print(Yellow + "")
	Slow("Chapeau de l'aventurier ", 3)
	fmt.Print("" + Reset)
	Slow("--", 3)
	fmt.Print("\n")
	Slow("Requiert ", 3)
	fmt.Print(Yellow + "")
	Slow("1 Plume de corbeau ", 3)
	fmt.Print("" + Reset)
	Slow("et ", 3)
	fmt.Print(Yellow + "")
	Slow("1 Cuir de sanglier", 3)
	fmt.Print("" + Reset)
	fmt.Print("\n \n")
	Slow("-- ", 3)
	fmt.Print(Yellow + "")
	Slow("(2) ", 3)
	fmt.Print("" + Reset)
	Slow("Construire une ", 3)
	fmt.Print(Yellow + "")
	Slow("Tunique de l'Aventurier ", 3)
	fmt.Print("" + Reset)
	Slow("--", 3)
	fmt.Print("\n")
	Slow("Requiert ", 3)
	fmt.Print(Yellow + "")
	Slow("2 Fourrure de Loup ", 3)
	fmt.Print("" + Reset)
	Slow("et ", 3)
	fmt.Print(Yellow + "")
	Slow("1 Peau de Troll", 3)
	fmt.Print("" + Reset)
	fmt.Print("\n \n")
	Slow("-- ", 3)
	fmt.Print(Yellow + "")
	Slow("(3) ", 3)
	fmt.Print("" + Reset)
	Slow("Construire les ", 3)
	fmt.Print(Yellow + "")
	Slow("Bottes de l'aventurier ", 3)
	fmt.Print("" + Reset)
	Slow("--", 3)
	fmt.Print("\n")
	Slow("Requiert ", 3)
	fmt.Print(Yellow + "")
	Slow("1 Fourrure de Loup ", 3)
	fmt.Print("" + Reset)
	Slow("et ", 3)
	fmt.Print(Yellow + "")
	Slow("1 Cuir de Sanglier", 3)
	fmt.Print("" + Reset)
	fmt.Print("\n \n")
	Slow("-- ", 3)
	fmt.Print(Yellow + "")
	Slow("(0) ", 3)
	fmt.Print("" + Reset)
	Slow("Retourner à la ", 3)
	fmt.Print(Yellow + "")
	Slow("Place Principale ", 3)
	fmt.Print("" + Reset)
	Slow("--", 3)
	fmt.Print("\n \n")
	Slow("Votre inventaire: \n", 3)
	for i := 0; i < len(p.inventory); i++ {
		if p.inventory[i] != " " {
			Slow("---] ", 2)
			fmt.Print(""+Yellow, p.inventory[i], Reset+"")
			Slow(" [---", 2)
			fmt.Print("\n")
		}
	}
	p.UseForgeron(e3, e4, e5, e6, a)
}

func (p *Personnage) UseForgeron(e3, e4, e5, e6 *Monstre, a *Equipement)  {
	var enclume int
	fmt.Scanln(&enclume)
	switch enclume {
	case 1:
		if p.Checkinv("Plume de Corbac") && p.Checkinv("Cuir de Sanglier") && !p.Checkinv("Chapeau de l'aventurier") {
			p.AddInventory("Chapeau de l'aventurier", 5)
			p.RemoveInventory("Plume de Corbac")
			p.RemoveInventory("Cuir de Sanglier")
			Slow("Vous êtes désormais en possession du ", 3)
			fmt.Print(Yellow + "")
			Slow("Chapeau de l'aventurier", 3)
			fmt.Print("" + Reset)
			time.Sleep(2 * time.Second)
			p.SuperAccessForgeron(e3, e4, e5, e6, a)
		} else {
			Slow("Tu te moques de moi ? Regarde ton inventaire l'ami", 3)
			time.Sleep(2 * time.Second)
			p.SuperAccessForgeron(e3, e4, e5, e6, a)
		}
	case 2:
		if p.Checkinv("Fourrure de Loup") && p.Checkinv("Peau de Troll") {
			p.RemoveInventory("Fourrure de Loup")
			if p.Checkinv("Fourrure de Loup") {
				p.RemoveInventory("Fourrure de Loup")
				p.RemoveInventory("Peau de Troll")
				p.AddInventory("Tunique de l'Aventurier", 5)
				Slow("Vous êtes désormais en possession de la ", 2)
				fmt.Print(Yellow + "")
				Slow("Tunique de l'aventurier", 2)
				fmt.Print("" + Reset)
				time.Sleep(2 * time.Second)
				p.SuperAccessForgeron(e3, e4, e5, e6, a)
			} else {
				p.AddInventory("Fourrure de Loup", 5)
			}
		} else {
			Slow("Tu te moques de moi ? Regarde ton ", 3)
			fmt.Print(Yellow + "")
			Slow("inventaire ", 3)
			fmt.Print("" + Reset)
			Slow("l'ami", 2)
			time.Sleep(2 * time.Second)
			p.SuperAccessForgeron(e3, e4, e5, e6, a)
		}
	case 3:
		if p.Checkinv("Fourrure de Loup") && p.Checkinv("Cuir de Sanglier") {
			p.RemoveInventory("Fourrure de Loup")
			p.RemoveInventory("Cuir de Sanglier")
			p.AddInventory("Bottes de l'Aventurier", 0)
			Slow("Vous êtes désormais en possession des ", 2)
			fmt.Print(Yellow + "")
			Slow("Bottes de l'aventurier", 2)
			fmt.Print("" + Reset)
			time.Sleep(2 * time.Second)
			p.SuperAccessForgeron(e3, e4, e5, e6, a)
		} else {
			Slow("Tu te moques de moi ? Regarde ton inventaire l'ami", 2)
			time.Sleep(2 * time.Second)
			p.SuperAccessForgeron(e3, e4, e5, e6, a)
		}
	case 0:
		p.menu(e3, e4, e5, e6, a)
	}
}


func (p *Personnage) SuperAccessForgeron(e3, e4, e5, e6 *Monstre, a *Equipement) {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("------ Vous entrez dans la ", Yellow+"Forge "+Reset,"------\n")
	fmt.Println("--", Yellow+"(1) "+Reset,"Construire un", Yellow+"Chapeau de l'aventurier "+Reset)
	fmt.Println("Requiert",Yellow + "1 Plume de corbeau"+Reset, "et", Yellow+"1 Cuir de sanglier\n"+Reset)
	fmt.Println("--",Yellow + "(2)"+ Reset, "Construire une",Yellow + "Tunique de l'Aventurier"+Reset, "--")
	fmt.Println("Requiert", Yellow +"2 Fourrure de Loup"+ Reset,"et",Yellow + "1 Peau de Troll\n" + Reset)
	fmt.Println("--", Yellow + "(3)"+Reset, "Construire les",Yellow + "Bottes de l'aventurier"+ Reset,"--")
	fmt.Println("Requiert", Yellow + "1 Fourrure de Loup"+ Reset, "et", Yellow + "1 Cuir de Sanglier\n"+ Reset)
	fmt.Println("--",Yellow +"(0)"+ Reset,"Retourner à la",Yellow + "Place Principale"+ Reset,"--\n")
	Slow("Votre inventaire:\n", 3)
	for i := 0; i < len(p.inventory); i++ {
		if p.inventory[i] != " " {
			fmt.Println("---]"+Yellow, p.inventory[i],Reset +"[---")
		}
	}
	p.UseForgeron(e3, e4, e5, e6, a)
}