package main

import (
	"fmt"
	"os"
	"time"
)

func (p *Personnage) CharCreation(a *Equipement) {
	var name string
	var class string
	var level int
	var lpmax int
	var lp int
	var inventory []string
	var skill []string
	var money int
	var chapeau string
	var tunique string
	var bottes string

	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	Slow("Bienvenue dans", 1)
	fmt.Printf(Yellow + " ")
	Slow("Goldy", 1)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("\n" + Reset)
	Slow("Vous êtes dans le menu de ", 1)
	fmt.Print(Yellow + "")
	Slow("création de personnage\n", 1)
	fmt.Print("" + Reset)
	Slow("Pour commencer, choisissez un ", 1)
	fmt.Print(Yellow + "")
	Slow("nom ", 1)
	fmt.Print("" + Reset)
	Slow("pour votre ", 1)
	fmt.Print(Yellow + "")
	Slow("avatar", 1)
	fmt.Print("" + Reset)
	Slow(": \n", 1)
	fmt.Scanln(&name)
	if name == "Utilisateur" { //Easter egg n°1, le mode développeur du jeu
		class = "Utilisateur"
		level = 1
		lpmax = 9999
		lp = lpmax
		inventory = []string{"Véritable Couteau"}
		money = 10000
		Slow("Bienvenue Utilisateur\n", 1)
		time.Sleep(300 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	} else {
		Slow("Votre personnage se nomme donc ", 1)
		fmt.Print(Yellow + "")
		Slow(name, 1)
		fmt.Print("" + Reset)
		time.Sleep(1 * time.Second)
		Slow("\nChoisissez maintenant la ", 1)
		fmt.Print(Yellow + "")
		Slow("race ", 1)
		fmt.Print("" + Reset)
		Slow("de ", 1)
		fmt.Print(Yellow + "")
		Slow(name, 1)
		fmt.Print("" + Reset)
		Slow(" parmi:\n", 1)
		fmt.Print(Yellow + "")
		Slow("\n-Humain\n", 1)
		Slow("\n-Elfe\n", 1)
		Slow("\n-Nain\n \n", 1)
		fmt.Print("" + Reset)
		fmt.Scanln(&class)
		if class != "Humain" && class != "Elfe" && class != "Nain" {
			Slow("Erreur, veuillez entrer une valeur correcte:\n", 1)
			Slow("Humain, Elfe ou Nain\n", 1)
			fmt.Scanln(&class)
		}
		switch class {
		case "Humain":
			class = "Humain"
			lpmax = 100
			chapeau = "Chapeau de paille"
			tunique = "Vieux manteau"
			bottes = "Vieille claquette"
		case "Elfe":
			class = "Elfe"
			lpmax = 80
			chapeau = "Chapeau d'érudit"
			tunique = "Robe de sage"
			bottes = "Chaussure pointue"
		case "Nain":
			class = "Nain"
			lpmax = 120
			chapeau = "Casque de mineur"
			tunique = "Salopette rapiécée"
			bottes = "Sabot renforcé"
		}
		if class != "Humain" && class != "Elfe" && class != "Nain" { //Easter egg n°3
			class = "Troll"
			lpmax = 50
			chapeau = "Bonnet de petite taille"
			tunique = "Veste abîmée"
			bottes = "Sabot en boît"

		}
		lp = lpmax / 2
		money = 100
		inventory = []string{"Potion de vie", "Potion de vie", "Potion de vie"}
		skill = []string{"Coup de poing"}
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Printf(Yellow + "")
		Slow(name, 1)
		fmt.Print("" + Reset)
		Slow(" est donc un ", 1)
		fmt.Printf(Yellow + "")
		Slow(class, 1)
		fmt.Print("" + Reset)
		Slow(", il commence avec ", 1)
		fmt.Printf(Yellow + "")
		fmt.Print(lp)
		Slow(" points de vie ", 1)
		fmt.Print("" + Reset)
		Slow("et ", 1)
		fmt.Printf(Yellow + "")
		fmt.Print(lpmax)
		Slow(" points de vie maximum.\n", 1)
		fmt.Print("" + Reset)
		if class == "Troll" {
			Slow("Ça t'apprendras à pas savoir écrire\n", 1)
		}
		fmt.Printf(Yellow + "")
		Slow(name, 1)
		fmt.Print("" + Reset)
		Slow(" est ", 1)
		fmt.Printf(Yellow + "")
		Slow("niveau 1, ", 1)
		fmt.Print("" + Reset)
		Slow("possède le sort ", 1)
		fmt.Printf(Yellow + "")
		Slow("Coup de poing ", 1)
		fmt.Print("" + Reset)
		Slow("et a ", 1)
		fmt.Printf(Yellow + "")
		fmt.Print(money)
		Slow(" pièces", 1)
		fmt.Print("" + Reset)
		level = 1
		time.Sleep(3 * time.Second)
	}
	Slow("\nLancement de la simulation\n", 1)
	a.Chapeau = chapeau
	a.Tunique = tunique
	a.Bottes = bottes
	p.Init(name, class, level, lpmax, lp, inventory, skill, money)
}

func (p *Personnage) Init(name string, class string, level int, lpmax int, lp int, inventory []string, skill []string, money int) {
	// initialisation de notre personnage
	p.name = name
	p.class = class
	p.level = level
	p.lpmax = lpmax
	p.lp = lp
	p.inventory = inventory
	p.skill = skill
	p.money = money
}
