package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/01-edu/z01"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

type Personnage struct {
	// creation de la structure de notre personnage
	name      string
	class     string
	level     int
	lpmax     int
	lp        int
	inventory []string
	skill     []string
	money     int
	Equipement
}

type Equipement struct {
	Chapeau string
	Tunique string
	Bottes  string
	Arme 	string
}

type Monstre struct {
	name   string
	lp     int
	lpmax  int
	attack int
}

func main() {
	// fonction qui execute nos sous fonctions et rentre les valeur ainsi que le menu principal
	var a Equipement
	var p1 Personnage
	p1.CharCreation(&a)
	var e3 Monstre
	e3.InitGoblin("Python", 160, 8)
	var e4 Monstre
	e4.InitGoblin("Java", 200, 10)
	var e5 Monstre
	e5.InitGoblin("C++", 300, 15)
	var e6 Monstre
	e6.InitGoblin("Golang", 400, 20)
	Slow("Vous vous rendez sur la ", 1)
	fmt.Print(Yellow + "")
	Slow("Place Principale", 1)
	fmt.Print("" + Reset)
	time.Sleep(1 * time.Second)
	fmt.Printf(". ")
	time.Sleep(1 * time.Second)
	fmt.Printf(". ")
	time.Sleep(1 * time.Second)
	fmt.Printf(".")
	time.Sleep(1 * time.Second)
	p1.menu(&e3, &e4, &e5, &e6, &a)
}

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

func (p *Personnage) menu(e3, e4, e5, e6 *Monstre, a *Equipement) {
	var menu int
	var retour int
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	Slow("+++++++++++++++++++++++++++++++\n", 3)
	Slow("Que voulez vous faire ?\n", 3)
	Slow("----- \n", 3)
	if !p.Checkinv("Objet Suspicieux") {
		Slow(Yellow+"(1) "+Reset, 3)
		Slow("Aller voir l'homme mystérieux sur le ", 3)
		Slow(Yellow+"banc\n"+Reset, 3)
		Slow("----- \n", 3)
	} else {
		Slow(Yellow+"(1)"+Reset, 3)
		Slow("Aller sur le ", 3)
		Slow(Yellow+"banc\n"+Reset, 3)
		Slow("----- \n", 3)
	}
	Slow(Yellow+"(2) "+Reset, 3)
	fmt.Print("" + Reset)
	Slow("Regarder la carte ", 3)
	fmt.Print(Yellow + "")
	Slow("d'Identité\n", 3)
	fmt.Print("" + Reset)
	Slow("----- \n", 3)
	fmt.Print(Yellow + "")
	Slow("(3) ", 3)
	fmt.Print("" + Reset)
	Slow("Jeter un oeil dans le ", 3)
	fmt.Print(Yellow + "")
	Slow("Sac\n", 3)
	fmt.Print("" + Reset)
	Slow("----- \n", 3)
	fmt.Print(Yellow + "")
	Slow("(4) ", 3)
	fmt.Print("" + Reset)
	Slow("Se rendre dans la boutique du ", 3)
	fmt.Print(Yellow + "")
	Slow("Marchand\n", 3)
	fmt.Print("" + Reset)
	Slow("----- \n", 3)
	fmt.Print(Yellow + "")
	Slow("(5) ", 3)
	fmt.Print("" + Reset)
	Slow("Aller à la ", 3)
	fmt.Print(Yellow + "")
	Slow("Forge\n", 3)
	fmt.Print("" + Reset)
	Slow("----- \n", 3)
	fmt.Print(Yellow + "")
	Slow("(6) ", 3)
	fmt.Print("" + Reset)
	Slow("Partir en exploration dans le ", 3)
	fmt.Print(Yellow + "")
	Slow("Donjon\n", 3)
	fmt.Print("" + Reset)
	Slow("----- \n", 3)
	fmt.Print(Yellow + "")
	Slow("(7) ", 3)
	fmt.Print("" + Reset)
	Slow("Affronter les ", 3)
	fmt.Print(Yellow + "")
	Slow("Quatres Grands\n", 3)
	fmt.Print("" + Reset)
	Slow("----- \n", 3)
	fmt.Print(Yellow + "")
	Slow("(0) ", 3)
	fmt.Print("" + Reset)
	Slow("Quitter la ", 3)
	fmt.Print(Yellow + "")
	Slow("Simulation", 3)
	fmt.Print("" + Reset)
	Slow("\n-----\n", 3)
	Slow("Entrez le numéro de l'option:\n", 3)
	Slow("+++++++++++++++++++++++++++++++\n", 3)
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		if !p.Checkinv("Objet Suspicieux") {
			p.QuestMan(e3, e4, e5, e6)
		} else {
			Slow("Il n'y a ", 5)
			Slow(Yellow+"personne ", 5)
			Slow("ici", 5)
			time.Sleep(2 * time.Second)
		}
		p.menu(e3, e4, e5, e6, a)
	case 2:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.DisplayInfo(a)
		fmt.Print(Yellow + "")
		Slow("(0) ", 2)
		fmt.Print("" + Reset)
		Slow("Arrêter de regarder la ", 2)
		fmt.Print(Yellow + "")
		Slow("carte\n", 2)
		fmt.Print("" + Reset)
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
		Slow(Yellow+" (1) "+Reset,3)
		Slow("Potion de vie ", 3)
		Slow(Yellow+"-->"+Reset, 3)
		Slow("15 pièces ", 3)
		Slow(Yellow+"<--\n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(2) "+Reset,3)
		Slow("Potion de poison", 3)
		Slow(Yellow+"-->"+Reset, 3)
		Slow("20 pièces", 3)
		Slow(Yellow+"<--\n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(3) "+Reset,3)
		Slow("Livre de sort: Boule de Feu", 3)
		Slow(Yellow+"-->"+Reset, 3)
		Slow("25 pièces", 3)
		Slow(Yellow+"<--\n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(4) "+Reset,3)
		Slow("Fourrure de Loup", 3)
		Slow(Yellow+"-->"+Reset, 3)
		Slow("5 pièces", 3)
		Slow(Yellow+"<--\n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(5) "+Reset,3)
		Slow("Peau de Troll", 3)
		Slow(Yellow+"-->"+Reset, 3)
		Slow("6 pièces", 3)
		Slow(Yellow+"<--\n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(6) "+Reset,3)
		Slow("Cuir de Sanglier", 3)
		Slow(Yellow+"-->"+Reset, 3)
		Slow("4 pièces", 3)
		Slow(Yellow+"<--\n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(7) "+Reset,3)
		Slow("Plume de Corbac", 3)
		Slow(Yellow+"-->"+Reset, 3)
		Slow("6 pièces", 3)
		Slow(Yellow+"<--\n"+Reset, 3)
		Slow("-----\n", 3)
		Slow(Yellow+"(0) "+Reset, 3)
		Slow("Retourner à la ", 3)
		Slow(Yellow+"Place Centrale"+Reset, 3)
		Slow("\n--", 3)
		Slow("++++++++++++++++++++++++++++++++++++++++++++++++\n", 3)
		p.Marchand(e3, e4, e5, e6, a)
	case 5:
		p.Forgeron(e3, e4, e5, e6, a)
	case 6:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.TrainingFight(a)
		p.menu(e3, e4, e5, e6, a)
	case 7:
		p.TheFirst(e3, e4, e5, e6, a)
		p.TheSecond(e3, e4, e5, e6, a)
		p.TheThird(e3, e4, e5, e6, a)
		p.TheFourth(e3, e4, e5, e6, a)
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

func (p *Personnage) QuestMan(e3, e4, e5, e6 *Monstre) {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	var choice int
	if !p.Checkinv("Objet Suspicieux") {
		if !p.Checkinv("Véritable Couteau") {
				Slow(Yellow+"Bienvenue sur la Place Principale l'ami.\n"+Reset, 1)
				Slow("(1) Lui demander son nom ?\n", 1)
				Slow("(2) Lui demander ce qu'on fait la ?\n", 1)
				fmt.Scanln(&choice)
				switch choice {
				case 1:
					Slow(Yellow+"Mon nom importe peu l'ami, je ne .", 1)
					Slow("\nBonne chance l'ami.\n"+Reset, 1)
				case 2:
					Slow(Yellow+"Vois tu depuis peu les contrées de Goldy se sont fait envahir par quatres seigneur malveillants.", 1)
					Slow("\nTa quête est de les éliminer, un par un et jusqu'au dernier.\n", 1)
					Slow("\nMais avant cela tu dois te rendre dans le donjon afin d'y trouver l'arme qui te donnera la force de les battre.", 1)
					Slow("\nBonne chance l'ami.\n"+Reset, 1)
				}
		} else {
			Slow(Yellow+"Bravo l'ami ! Tu as obtenus l'artéfact.", 1)
			Slow("\nJe te conseille maintenant d'aller t'équiper en conséquence et de te lancer à la recherche des Quatres Grands"+Reset, 1)
		}
		if e3.lp <= 0 {
			if e4.lp <= 0 {
				if e5.lp <= 0 {
					Slow(Yellow + "Bien joué l'ami ! Tu as vaincu ",1)
					Slow(e5.name,1)
					Slow(", il ne t'en reste plus qu'un'" + Reset,1)
				}
				Slow(Yellow + "Bien joué l'ami ! Tu as vaincu ",1)
				Slow(e4.name,1)
				Slow(", il ne t'en reste plus que deux" + Reset,1)
			}
			Slow(Yellow + "Bien joué l'ami ! Tu as vaincu ",1)
			Slow(e3.name,1)
			Slow(", il ne t'en reste plus que trois" + Reset,1)
		}
	}
	time.Sleep(2 * time.Second)
}

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
	Slow("\n-------------------------\n", 2)
}

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

func (a *Equipement) EquipLpBonus(item string, p *Personnage) bool {
	if item == "chapeau" {
		if a.Chapeau!= "Chapeau de l'aventurier" {
			p.AddInventory(a.Chapeau, 0)
			a.Chapeau = "Chapeau de l'aventurier"
			return true
		} else {
			return false
		}
	} else if item == "tunique" {
		if a.Tunique!= "Tunique de l'aventurier" {
			p.AddInventory(a.Tunique, 0)
			a.Tunique = "Tunique de l'aventurier"
			return true
		} else {
			return false
		}
	} else if item == "bottes" {
		if a.Bottes!= "Bottes de l'aventurier" {
			p.AddInventory(a.Bottes, 0)
			a.Bottes = "Bottes de l'aventurier"
			return true
		} else {
			return false
		}
	}
	return false
}

func (p *Personnage) TakePot() {
	// fonction qui nous permet de prendre une potion de soin
	for _, letter := range p.inventory {
		if letter == "Potion de vie" {
			if p.lp <= (p.lpmax - 50) {
				p.lp += 50
				Slow("Glou glou glou, ça fait du bien", 2)
				fmt.Print("\n")
				Slow("Vous avez désormais: ", 2)
				fmt.Print(Yellow + "")
				fmt.Print(p.lp)
				Slow("/", 2)
				fmt.Print(p.lpmax)
				fmt.Print("\n" + Reset)
				p.RemoveInventory("Potion de vie")
				break
			} else if p.lp > (p.lpmax-50) && p.lp < p.lpmax {
				p.lp = p.lpmax
				Slow("Glou glou glou, ça fait du bien", 2)
				fmt.Print("\n")
				Slow("Vous avez désormais: ", 2)
				fmt.Print(Yellow + "")
				fmt.Print(p.lp)
				Slow("/", 2)
				fmt.Print(p.lpmax)
				fmt.Print("\n" + Reset)
				p.RemoveInventory("Potion de vie")
				break
			} else {
				Slow("Vous êtes au ", 2)
				fmt.Print(Yellow + "")
				Slow("maximum", 2)
				fmt.Print("" + Reset)
				Slow(", vous ne pouvez pas utiliser la ", 2)
				fmt.Print(Yellow + "")
				Slow("potion", 2)
				fmt.Print("" + Reset)
				break
			}
		} else {
			Slow("Tu n'as pas de ", 2)
			fmt.Print(Yellow + "")
			Slow("potion", 2)
			fmt.Print("" + Reset)
		}
	}
	time.Sleep(1 * time.Second)
}

func (p *Personnage) PoisonPot() {
	// fonction qui crée la potion poison et explique ce qu'elle fait sur un personnage
	for _, letter := range p.inventory {
		if letter == "Potion de poison" {
			Slow("Vous buvez la ", 2)
			fmt.Print(Green + "")
			Slow("Potion de poison", 2)
			fmt.Print("" + Reset)
			Slow(", ouch !", 1)
			p.RemoveInventory("Potion de poison")
			time.Sleep(1 * time.Second)
			p.lp -= 10
			Slow("Vos points de vie ", 2)
			fmt.Print(Red + "")
			Slow("diminuent: ", 2)
			fmt.Print("" + Reset)
			fmt.Print(p.lp)
			Slow("/", 1)
			fmt.Print(p.lpmax)
			time.Sleep(1 * time.Second)
			Slow("Vos points de vie ", 2)
			fmt.Print(Red + "")
			Slow("diminuent: ", 2)
			fmt.Print("" + Reset)
			p.lp -= 10
			fmt.Print(p.lp)
			Slow("/", 1)
			fmt.Print(p.lpmax)
			fmt.Println(p.lp, "/", p.lpmax, "PV")
			time.Sleep(1 * time.Second)
			Slow("Vos points de vie ", 2)
			fmt.Print(Red + "")
			Slow("diminuent: ", 2)
			fmt.Print("" + Reset)
			p.lp -= 10
			fmt.Print(p.lp)
			Slow("/", 1)
			fmt.Print(p.lpmax)
			fmt.Println(p.lp, "/", p.lpmax, "PV")
			break
		} else {
			Slow("\nTu n'as pas de ", 2)
			fmt.Print(Yellow + "")
			Slow("potion", 2)
			fmt.Print("\n" + Reset)
			break
		}
	}
	time.Sleep(1 * time.Second)
}

func (p *Personnage) Dead(a *Equipement) {
	// fonction qui verifie si le personnage est mort et le ressucite a la moitié de ses pv
	if p.lp == 0 {
		Slow("Bravo, vous êtes ", 1)
		fmt.Print(Red + "")
		Slow("mort.", 1)
		fmt.Print("" + Reset)
		fmt.Print("\n")
		time.Sleep(1 * time.Second)
		Slow("Mais ne paniquez pas, vous allez être ", 1)
		fmt.Print(Yellow + "")
		Slow("ressucité", 1)
		fmt.Print("" + Reset)
		fmt.Print("\n")
		time.Sleep(1 * time.Second)
		Slow("Manoeuvre de ", 1)
		fmt.Print(Yellow + "")
		Slow("réanimation ", 1)
		fmt.Print("" + Reset)
		Slow("en cours", 1)
		time.Sleep(1 * time.Second)
		Slow(". ", 1)
		time.Sleep(1 * time.Second)
		Slow(". ", 1)
		time.Sleep(1 * time.Second)
		Slow(".", 1)
		fmt.Print("\n")
		time.Sleep(1 * time.Second)
		p.lp = p.lpmax / 2
		p.DisplayInfo(a)
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

func (p *Personnage) Marchand(e3, e4, e5, e6 *Monstre, a *Equipement) {
	// fonction affichant le menu du marchand , et les ajoute a notre inventaire
	var menum int
	Slow("Vous avez", 1)
	fmt.Print(Yellow + "")
	fmt.Print(p.money)
	Slow("pièces\n"+Reset, 1)
	fmt.Scanln(&menum)
	switch menum {
	case 1:
		p.AddInventory("Potion de vie", 15)
		p.Marchand(e3, e4, e5, e6, a)
	case 2:
		p.AddInventory("Potion de poison", 20)
		p.Marchand(e3, e4, e5, e6, a)
	case 3:
		p.AddInventory("Livre de sort: Boule de Feu", 25)
		p.Marchand(e3, e4, e5, e6, a)
	case 4:
		p.AddInventory("Fourrure de Loup", 5)
		p.Marchand(e3, e4, e5, e6, a)
	case 5:
		p.AddInventory("Peau de Troll", 5)
		p.Marchand(e3, e4, e5, e6, a)
	case 6:
		p.AddInventory("Cuir de Sanglier", 4)
		p.Marchand(e3, e4, e5, e6, a)
	case 7:
		p.AddInventory("Plume de Corbac", 6)
		p.Marchand(e3, e4, e5, e6, a)
	case 0:
		p.menu(e3, e4, e5, e6, a)
	}
}

func (p *Personnage) AddInventory(item string, price int) {
	if p.money >= price {
		p.money -= price
		p.inventory = append(p.inventory, item)
	} else {
		Slow("Tu n'as pas assez ", 2)
		fmt.Print(Yellow + "")
		Slow("d'argent ", 2)
		fmt.Print("" + Reset)
		Slow("pour acheter cet objet", 1)
	}
}

func (p *Personnage) RemoveInventory(item string) {
	var index int = -1
	for i := range p.inventory {
		if p.inventory[i] == item {
			index = i
		}
	}
	if index != -1 {
		p.inventory = append(p.inventory[:index], p.inventory[index+1:]...)
	}
}

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
	Slow("Votre inventaire: ", 3)
	for i := 0; i < len(p.inventory); i++ {
		if p.inventory[i] != " " {
			Slow("---] ", 2)
			fmt.Print(""+Yellow, p.inventory[i], Reset+"")
			Slow(" [---", 2)
			fmt.Print("\n")
		}
	}
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
			p.Forgeron(e3, e4, e5, e6, a)
		} else {
			Slow("Tu te moques de moi ? Regarde ton inventaire l'ami", 3)
			time.Sleep(2 * time.Second)
			p.Forgeron(e3, e4, e5, e6, a)
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
				p.Forgeron(e3, e4, e5, e6, a)
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
			p.Forgeron(e3, e4, e5, e6, a)
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
			p.Forgeron(e3, e4, e5, e6, a)
		} else {
			Slow("Tu te moques de moi ? Regarde ton inventaire l'ami", 2)
			time.Sleep(2 * time.Second)
			p.Forgeron(e3, e4, e5, e6, a)
		}
	case 0:
		p.menu(e3, e4, e5, e6, a)
	}
}

func (p *Personnage) Checkinv(item string) bool {
	var founditem bool = false
	for _, letter := range p.inventory {
		if letter == item {
			founditem = true
		}
	}
	return founditem
}

func (m *Monstre) InitGoblin(name string, lpmax int, attack int) {
	// initialisation de notre personnage
	m.name = name
	m.lpmax = lpmax
	m.lp = lpmax / 2
	m.attack = attack
}

func (p *Personnage) GoblinPattern(m *Monstre, att int) {
	p.lp -= m.attack * att
	Slow(Yellow+m.name+ Reset, 2)
	Slow(" attaque ", 2)
	Slow(Yellow+p.name+Reset, 2)
	Slow(" et lui inflige ", 2)
	fmt.Print(Red+"")
	fmt.Print(m.attack)
	Slow(" points de dégâts"+Reset,2)
	Slow(", il lui reste ", 2)
	fmt.Print(Yellow+"")
	fmt.Print(p.lp)
	Slow(" points de vies "+Reset,2)
	Slow("sur un total de ", 2)
	fmt.Print(Yellow+"")
	fmt.Print(p.lpmax)
	fmt.Println(""+Reset)
}

func (p *Personnage) CharTurn(m *Monstre) {
	var choice int
	var damage int = 5
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
		if p.Checkinv("Véritable Couteau") {
			m.lp = 0
			Slow("\nVous avez tué ", 2)
			Slow(m.name, 2)
		} else {
			m.lp -= damage
			fmt.Print(Yellow + "")
			Slow(p.name, 2)
			fmt.Print("" + Reset)
			Slow(" utilise ", 2)
			fmt.Print(Yellow + "")
			Slow("Coup de poing ", 2)
			fmt.Print("" + Reset)
			Slow(" et inflige ", 2)
			fmt.Print(Red + "")
			fmt.Print(damage)
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
		}
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
					Slow("---] ", 2)
					fmt.Print(""+Yellow, p.inventory[i], Reset+"")
					Slow(" [---\n", 2)
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

func (p *Personnage) TrainingFight(a *Equipement) {
	var e1 Monstre
	var e2 Monstre
	n := rand.Intn(5)
	if n == 4 {
		Slow("Un mimic sauvage apparaît", 2)
		time.Sleep(3 * time.Second)
		var turn int
		e2.InitGoblin("Mimic", 80, 5)
		for i := 0; i <= 9999; i++ {
			turn++
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			Slow("Tour ", 2)
			fmt.Print(turn)
			time.Sleep(1 * time.Second)
			Slow("\nC'est au joueur !", 2)
			p.CharTurn(&e2)
			if e2.lp <= 0 {
				Slow("Vous avez vaincu le Mimic ", 2)
				time.Sleep(1 * time.Second)
				Slow("vous le fouillez et obtenez un objet étrange...", 5)
				time.Sleep(3 * time.Second)
				if !p.Checkinv("Véritable couteau") {
					p.AddInventory("Véritable couteau", 0)
				}
				break
			}
			time.Sleep(1 * time.Second)
			Slow("\nC'est à l'ennemi !\n", 2)
			p.GoblinPattern(&e2, 1)
			if p.lp <= 0 {
				Slow(e2.name, 2)
				Slow(" vous a battu", 2)
				p.Dead(a)
				break
			}
			time.Sleep(3 * time.Second)
		}
	} else {
		e1.InitGoblin("Gobelin d'entrainement", 40, 5)
		Slow("Le ", 1)
		Slow(e1.name, 1)
		Slow(" est prêt à se battre", 1)
		time.Sleep(3 * time.Second)
		var turn int
		for i := 0; i <= 9999; i++ {
			turn++
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			Slow("Tour ", 2)
			fmt.Print(turn)
			time.Sleep(1 * time.Second)
			Slow("\nC'est au joueur !\n", 2)
			p.CharTurn(&e1)
			if e1.lp <= 0 {
				Slow("Vous avez vaincu ", 2)
				Slow(e1.name, 2)
				time.Sleep(2 * time.Second)
				break
			}
			time.Sleep(1 * time.Second)
			Slow("C'est à l'ennemi !\n", 2)
			if turn%3 == 3 || turn == 3 {
				p.GoblinPattern(&e1, 2)
			} else {
				p.GoblinPattern(&e1, 1)
			}
			if p.lp <= 0 {
				Slow(e1.name, 2)
				Slow(" vous a battu\n", 2)
				p.Dead(a)
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}

func (p *Personnage) TheFirst(e3, e4, e5, e6 *Monstre, a *Equipement) {
	if e3.lp <= 0 {
		p.TheSecond(e3, e4, e5, e6, a)
	} else {
		Slow("Le Grand Python se présente devant vous\n", 1)
		time.Sleep(1 * time.Second)
		var turn int
		for i := 0; i <= 9999; i++ {
			turn++
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			Slow("Tour", 1)
			fmt.Println(turn)
			Slow("C'est au joueur !", 2)
			p.CharTurn(e3)
			if e3.lp <= 0 {
				Slow("\nVous avez vaincu le Grand ", 1)
				Slow(e3.name, 1)
				Slow(" !", 1)
				time.Sleep(2 * time.Second)
				break
			}
			time.Sleep(1 * time.Second)
			Slow("\nC'est à l'ennemi !\n", 2)
			p.GoblinPattern(e3, 1)
			if p.lp <= 0 {
				Slow(e3.name, 1)
				Slow(" vous a battu", 1)
				p.Dead(a)
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}

func (p *Personnage) TheSecond(e3, e4, e5, e6 *Monstre, a *Equipement) {
	if e4.lp <= 0 {
		p.TheThird(e3, e4, e5, e6, a)
	} else {
		Slow("Le Grand Java se présente devant vous\n", 1)
		time.Sleep(1 * time.Second)
		var turn2 int
		for i := 0; i <= 9999; i++ {
			turn2++

			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			Slow("Tour", 2)
			fmt.Print(turn2)
			Slow("\nC'est au joueur !", 2)
			p.CharTurn(e4)
			if e4.lp <= 0 {
				Slow("\nVous avez vaincu le Grand ", 1)
				Slow(e4.name, 1)
				Slow(" !", 1)
				time.Sleep(2 * time.Second)
				break
			}
			time.Sleep(1 * time.Second)
			Slow("\nC'est à l'ennemi !\n", 2)
			p.GoblinPattern(e4, 1)
			if p.lp <= 0 {
				Slow(e4.name, 1)
				Slow(" vous a battu", 1)
				p.Dead(a)
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}

func (p *Personnage) TheThird(e3, e4, e5, e6 *Monstre, a *Equipement) {
	if e5.lp <= 0 {
		p.TheFourth(e3, e4, e5, e6, a)
	} else {
		Slow("Le Grand C++ se présente devant vous\n", 1)
		time.Sleep(1 * time.Second)
		var turn3 int
		for i := 0; i <= 9999; i++ {
			turn3++

			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			Slow("Tour", 2)
			fmt.Print(turn3)
			fmt.Println("\nC'est au joueur !", 2)
			p.CharTurn(e5)
			if e5.lp <= 0 {
				Slow("\nVous avez vaincu le Grand ", 1)
				Slow(e5.name, 1)
				Slow(" !", 1)
				time.Sleep(2 * time.Second)
				break
			}
			time.Sleep(1 * time.Second)
			Slow("\nC'est à l'ennemi !\n", 2)
			p.GoblinPattern(e5, 1)
			if p.lp <= 0 {
				Slow(e5.name, 1)
				Slow(" vous a battu", 1)
				p.Dead(a)
				break

			}
			time.Sleep(3 * time.Second)
		}
	}
}

func (p *Personnage) TheFourth(e3, e4, e5, e6 *Monstre, a *Equipement) {
	if e6.lp <= 0 {
		Slow("Vous avez déjà vaincu les Quatres Grands, il n'y a plus personne ici",1)
	} else {
		Slow("Le plus Grand des Grand, Golang, se présente devant vous\n", 1)
		time.Sleep(1 * time.Second)
		var turn4 int
		for i := 0; i <= 9999; i++ {
			turn4++

			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")

			Slow("Tour", 2)
			fmt.Println(turn4)
			Slow("C'est au joueur !", 2)
			p.CharTurn(e6)
			if e6.lp <= 0 {
				Slow("Vous avez vaincu Golang, le plus Grand des Grands.\n", 1)
				Slow("Dans son dernier souffle il lâche un objet.", 4)
				p.AddInventory("Objet suspicieux", 0)
				time.Sleep(2 * time.Second)
				break
			}
			time.Sleep(1 * time.Second)
			Slow("\nC'est à l'ennemi !\n", 2)
			p.GoblinPattern(e6, 1)
			if p.lp <= 0 {
				Slow(e6.name, 1)
				Slow(" vous a battu", 1)
				p.Dead(a)
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}

func Slow(s string, v int) {
	if v == 1 {
		for _, letter := range s {
			time.Sleep(50 * time.Millisecond)
			z01.PrintRune(letter)
		}
	} else if v == 2 {
		for _, letter := range s {
			time.Sleep(25 * time.Millisecond)
			z01.PrintRune(letter)
		}
	} else if v == 3 {
		for _, letter := range s {
			time.Sleep(1 * time.Millisecond)
			z01.PrintRune(letter)
		}
	} else if v == 4 {
		for _, letter := range s {
			time.Sleep(100 * time.Millisecond)
			z01.PrintRune(letter)
		}
	} else if v == 5 {
		for _, letter := range s {
			time.Sleep(200 * time.Millisecond)
			z01.PrintRune(letter)
		}
	}
}
