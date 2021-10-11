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
}

type Monstre struct {
	name   string
	lp     int
	lpmax  int
	attack int
}

func main() {
	// fonction qui execute nos sous fonctions et rentre les valeur ainsi que le menu principal
	var p1 Personnage
	p1.CharCreation()
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
	p1.menu()
}

func (p *Personnage) CharCreation() {
	var name string
	var class string
	var level int
	var lpmax int
	var lp int
	var inventory []string
	var skill []string
	var money int

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
		case "Elfe":
			class = "Elfe"
			lpmax = 80
		case "Nain":
			class = "Nain"
			lpmax = 120
		}
		if class != "Humain" && class != "Elfe" && class != "Nain" { //Easter egg n°3
			class = "Troll"
			lpmax = 50
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

func (p *Personnage) menu() {
	var menu int
	var retour int
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	Slow("+++++++++++++++++++++++++++++++\n", 3)
	Slow("Que voulez vous faire ?\n", 3)
	Slow("----- \n", 3)
	if p.Checkinv("Objet Suspicieux") == false {
		Slow(Yellow + "(1) " + Reset,3)
		Slow("Aller voir l'homme mystérieux sur le ",3)
		Slow(Yellow + "banc\n",3)
	} else {
		Slow(Yellow + "(1)" + Reset,3)
		Slow("Aller sur le ",3)
		Slow(Yellow + "banc\n",3)
	}
	Slow("(2) " + Reset, 3)
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
		if p.Checkinv("Objet Suspicieux") == false {
			p.QuestMan()
		}
		Slow("Il n'y a ",5)
		Slow(Yellow + "personne ",5)
		Slow("ici",5)
		time.Sleep(2 * time.Second)
		p.menu()
	case 2:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.DisplayInfo()
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
			p.menu()
		}
	case 3:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.AccessInventory()
	case 4:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		Slow("+++++++++++++++++++++Marchand+++++++++++++++++++\n", 2)
		Slow("-----\n", 2)
		fmt.Print(Yellow + "")
		Slow(" (1) Potion de vie ", 2)
		fmt.Print("" + Reset)
		Slow("-->", 2)
		fmt.Print(Yellow + "")
		Slow("15 pièces", 2)
		fmt.Print("" + Reset)
		Slow("<--", 2)
		Slow("-----\n", 2)
		fmt.Print(Yellow + "")
		Slow("(2) Potion de poison", 2)
		fmt.Print("" + Reset)
		Slow("-->", 2)
		fmt.Print(Yellow + "")
		Slow("20 pièces", 2)
		fmt.Print("" + Reset)
		Slow("<--", 2)
		Slow("-----\n", 2)
		fmt.Print(Yellow + "")
		Slow("(3) Livre de sort: Boule de Feu", 2)
		fmt.Print("" + Reset)
		Slow("-->", 2)
		fmt.Print(Yellow + "")
		Slow("25 pièces", 2)
		fmt.Print("" + Reset)
		Slow("<--", 2)
		Slow("-----\n", 2)
		fmt.Print(Yellow + "")
		Slow("(4) Fourrure de Loup", 2)
		fmt.Print("" + Reset)
		Slow("-->", 2)
		fmt.Print(Yellow + "")
		Slow("5 pièces", 2)
		fmt.Print("" + Reset)
		Slow("<--", 2)
		Slow("-----\n", 2)
		fmt.Print(Yellow + "")
		Slow("(5) Peau de Troll", 2)
		fmt.Print("" + Reset)
		Slow("-->", 2)
		fmt.Print(Yellow + "")
		Slow("6 pièces", 2)
		fmt.Print("" + Reset)
		Slow("<--", 2)
		Slow("-----\n", 2)
		fmt.Print(Yellow + "")
		Slow("(6) Cuir de Sanglier", 2)
		fmt.Print("" + Reset)
		Slow("-->", 2)
		fmt.Print(Yellow + "")
		Slow("4 pièces", 2)
		fmt.Print("" + Reset)
		Slow("<--", 2)
		Slow("-----\n", 2)
		fmt.Print(Yellow + "")
		Slow("(7) Plume de Corbac", 2)
		fmt.Print("" + Reset)
		Slow("-->", 2)
		fmt.Print(Yellow + "")
		Slow("6 pièces", 2)
		fmt.Print("" + Reset)
		Slow("<--", 2)
		Slow("-----\n", 2)
		fmt.Print(Yellow + "")
		Slow("(0)", 2)
		fmt.Print("" + Reset)
		Slow("Retourner à la ", 2)
		fmt.Print(Yellow + "")
		Slow("place centrale", 2)
		fmt.Print("" + Reset)
		Slow("\n--", 2)
		Slow("++++++++++++++++++++++++++++++++++++++++++++++++", 2)
		p.Marchand()
	case 5:
		p.Forgeron()
	case 6:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.TrainingFight()
		p.menu()
	case 7:
		p.TheFirst()
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

func(p *Personnage) QuestMan() {
	var choice int
	var rappel1 bool = false
	if p.Checkinv("Objet Suspicieux") == false {
		if p.Checkinv("Véritable Couteau") == false {
			if rappel1 == false {
				Slow(Yellow + "Bienvenue sur la Place Principale l'ami.\n" + Reset,1)
				Slow("(1) Lui demander son nom ?",1)
				Slow("(2) Lui demander ce qu'on fait la ?",1)
				fmt.Scanln(&choice)
				switch choice {
				case 1:
					Slow(Yellow + "Mon nom importe peu l'ami ! Tu es la pour bien plus.\n Vois tu depuis peu les contrées de Goldy se sont fait envahir par quatres seigneur malveillants.",1)
					Slow("\nTa quête est de les éliminer, un par un et jusqu'au dernier.\n",1)
					Slow("\nMais avant cela tu dois te rendre dans le donjon afin d'y trouver l'arme qui te donnera la force de les battre.",1)
					Slow("\nBonne chance l'ami.\n" + Reset,1)
					rappel1 = true
				case 2:
					Slow(Yellow + "Vois tu depuis peu les contrées de Goldy se sont fait envahir par quatres seigneur malveillants.",1)
					Slow("\nTa quête est de les éliminer, un par un et jusqu'au dernier.\n",1)
					Slow("\nMais avant cela tu dois te rendre dans le donjon afin d'y trouver l'arme qui te donnera la force de les battre.",1)
					Slow("\nBonne chance l'ami.\n" + Reset,1)
					rappel1 = true
				}
			} else if rappel1 == true {
				Slow(Yellow + "Tu dois te rendre dans le donjon pour obtenir un artéfact qui te permettra de vaincre les Quatres Grands." + Reset,2)
			}
		} else {
			Slow(Yellow + "Bravo l'ami ! Tu as obtenus l'artéfact.",1)
			Slow("\nJe te conseille maintenant d'aller t'équiper en conséquence et de te lancer à la recherche des Quatres Grands" + Reset,1)
		}
	}
	time.Sleep(2 * time.Second)
	p.menu()
}

func (p *Personnage) DisplayInfo() {
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
	fmt.Print(p.money)
	Slow("\n-------------------------\n", 2)
}

func (p *Personnage) AccessInventory() {
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
	p.UseInventory()
}

func (p *Personnage) SuperAccessInventory() {
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
	p.UseInventory()
}

func (p *Personnage) UseInventory() {
	// fonction qui nous permet d'interagir avec les éléments de l'inventaire
	var use int
	fmt.Scanln(&use)
	switch use {
	case 1:
		p.TakePot()
		p.SuperAccessInventory()
	case 2:
		p.PoisonPot()
		p.SuperAccessInventory()
	case 3:
		p.Spellbook()
		p.SuperAccessInventory()
	case 4:
		if p.Checkinv("Chapeau de l'aventurier") && p.lpmax > p.lpmax+15 {
			p.Chapeau = "Chapeau de l'aventurier"
			p.lpmax += 15
			p.RemoveInventory("Chapeau de l'Aventurier")
			Slow("Vous avez désormais +15 hp avec le couvre chef !\n", 2)
			p.SuperAccessInventory()
		} else if !p.Checkinv("Chapeau de L'aventurier") {
			Slow("Tu l'as déja équiper !", 2)
			p.SuperAccessInventory()
		}
	case 5:
		if !p.Checkinv("Tunique de l'Aventurier") && p.lpmax > p.lpmax+20 {
			p.Tunique = "Tunique de l'Aventurier"
			p.lpmax += 20
			Slow("Vous avez désormais +20 hp grace a la tunique !\n", 2)
			p.RemoveInventory("Tunique de l'Aventurier")
			p.SuperAccessInventory()
		} else {
			Slow("Tu l'as déja équiper !", 2)
			p.SuperAccessInventory()
		}
	case 6:
		if !p.Checkinv("Bottes de l'Aventurier") && p.lpmax > p.lpmax+10 {
			p.Bottes = "Bottes de l'Aventurier"
			p.lpmax += 10
			Slow("Vous avez désormais +10 hp grace a la paire !\n", 2)
			p.RemoveInventory("Bottes de l'Aventurier")
			p.SuperAccessInventory()
		} else {
			Slow("Tu l'as déja équiper !", 2)
			p.SuperAccessInventory()
		}
	case 0:
		p.menu()
	}
}

func (p *Personnage) TakePot() {
	// fonction qui nous permet de prendre une potion de soin
	for _, letter := range p.inventory {
		if letter == "Potion de vie" {
			if p.lp <= (p.lpmax - 50) {
				p.lp += 50
<<<<<<< HEAD
				Slow("Glou glou glou, ça fait du bien",2)
=======
				// fmt.Println("   :~:  ")
				// fmt.Println("   | |    ")
				// fmt.Println("  .' `.   	 ") On garde ou pas ?
				// fmt.Println(".'     `. ")
				// fmt.Println("|       | ")
				// fmt.Println(" `.._..' ")
				Slow("Glou glou glou, ça fait du bien", 2)
>>>>>>> 78b31279b37ae775c8acaf3aefc57bbc07565c90
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

func (p *Personnage) Dead() {
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
		p.DisplayInfo()
	}
}

func (p *Personnage) Spellbook() {
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
	p.menu()
}

func (p *Personnage) Marchand() {
	// fonction affichant le menu du marchand , et les ajoute a notre inventaire
	var menum int
	Slow("Vous avez", 1)
	fmt.Print(Yellow + "")
	fmt.Print(p.money)
	Slow("pièces", 1)
	fmt.Print("" + Reset)
	fmt.Scanln(&menum)
	switch menum {
	case 1:
		p.AddInventory("Potion de vie", 15)
		p.Marchand()
	case 2:
		p.AddInventory("Potion de poison", 20)
		p.Marchand()
	case 3:
		p.AddInventory("Livre de sort: Boule de Feu", 25)
		p.Marchand()
	case 4:
		p.AddInventory("Fourrure de Loup", 5)
		p.Marchand()
	case 5:
		p.AddInventory("Peau de Troll", 5)
		p.Marchand()
	case 6:
		p.AddInventory("Cuir de Sanglier", 4)
		p.Marchand()
	case 7:
		p.AddInventory("Plume de Corbac", 6)
		p.Marchand()
	case 0:
		p.menu()
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

func (p *Personnage) Forgeron() {
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
			p.Forgeron()
		} else {
			Slow("Tu te moques de moi ? Regarde ton inventaire l'ami", 3)
			time.Sleep(2 * time.Second)
			p.Forgeron()
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
				p.Forgeron()
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
			p.Forgeron()
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
			p.Forgeron()
		} else {
			Slow("Tu te moques de moi ? Regarde ton inventaire l'ami", 2)
			time.Sleep(2 * time.Second)
			p.Forgeron()
		}
	case 0:
		p.menu()
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
	Slow(m.name, 2)
	Slow(" attaque ", 2)
	Slow(p.name, 2)
	Slow(" et lui inflige ", 2)
	fmt.Print(m.attack)
	Slow(" points de dégâts, il lui reste", 2)
	fmt.Print(p.lp)
	Slow(" points de vies sur un total de", 2)
	fmt.Print(p.lpmax)
	fmt.Print("\n")
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
	Slow("Utiliser un objet", 2)
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		if p.RealKnife() {
			m.lp = 0
			Slow("\nVous avez tué ", 2)
			Slow(m.name, 2)
		} else {
			m.lp -= damage
			fmt.Print(Yellow + "")
			Slow(p.name, 2)
			fmt.Print("" + Reset)
			Slow(" utilise", 2)
			fmt.Print(Yellow + "")
			Slow(" Coup de poing ", 2)
			fmt.Print("" + Reset)
			Slow(" et inflige ", 2)
			fmt.Print(Red + "")
			fmt.Print(damage)
			fmt.Print("" + Reset)
			Slow(" points de dégâts à ", 2)
			fmt.Print(Yellow + "")
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
			fmt.Print(Yellow + "")
			Slow("vide", 2)
			Slow("\n(0) ", 2)
			Slow("Retour", 2)
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
			fmt.Print(Yellow + "")
			Slow("(1) ", 2)
			fmt.Print("" + Reset)
			Slow("Prendre une ", 2)
			fmt.Print(Yellow + "")
			Slow("Potion de soin", 2)
			fmt.Print("" + Reset)
			Slow("\n(2) ", 2)
			Slow("Prendre une ", 2)
			fmt.Print(Yellow + "")
			Slow("Potion de poison", 2)
			Slow("\n(0) ", 2)
			Slow("Retour", 2)
			fmt.Print("" + Reset)
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

func (p *Personnage) RealKnife() bool {
	var couteau bool
	for _, letter := range p.inventory {
		if letter == "Véritable couteau" {
			couteau = true
		} else {
			couteau = false
		}
	}
	return couteau
}

func (p *Personnage) TrainingFight() {
	var e1 Monstre
	var e2 Monstre
	n := rand.Intn(10)
	if n == 9 {
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
			Slow("\nC'est à l'ennemi !", 2)
			p.GoblinPattern(&e2, 1)
			if p.lp <= 0 {
				Slow(e2.name, 2)
				Slow(" vous a battu", 2)
				p.Dead()
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
				Slow("Vous avez vaincu", 2)
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
<<<<<<< HEAD
				Slow(e1.name,2)
				Slow(" vous a battu\n",2)
=======
				Slow(e1.name, 2)
				Slow(" vous a battu", 2)
>>>>>>> 78b31279b37ae775c8acaf3aefc57bbc07565c90
				p.Dead()
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}

func (p *Personnage) TheFirst() {
	var e3 Monstre
	e3.InitGoblin("Python", 80, 8)
	Slow("Le Grand Python se présente devant vous\n",1)
	time.Sleep(1 * time.Second)
	var turn int
	for i := 0; i <= 9999; i++ {
		turn++
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
<<<<<<< HEAD
		
		Slow("Tour",1)
		fmt.Println(turn)
		Slow("C'est au joueur !",2)
		p.CharTurn(&e3)
		if e3.lp <= 0 {
			Slow("\nVous avez vaincu le Grand ",1)
			Slow(e3.name,1)
			Slow(" !",1)
=======
		Slow("Le Grand Python se présente devant vous", 1)
		Slow("Tour", 1)
		fmt.Print(turn)
		time.Sleep(1 * time.Second)
		Slow("C'est au joueur !", 2)
		p.CharTurn(&e3)
		if e3.lp <= 0 {
			fmt.Print("\n")
			Slow("Vous avez vaincu le Grand ", 1)
			Slow(e3.name, 1)
			Slow(" !", 1)
>>>>>>> 78b31279b37ae775c8acaf3aefc57bbc07565c90
			time.Sleep(2 * time.Second)
			p.TheSecond()
			break
		}
		time.Sleep(1 * time.Second)
<<<<<<< HEAD
		Slow("\nC'est à l'ennemi !\n",2)
=======
		fmt.Print("\n")
		Slow("C'est à l'ennemi !", 2)
>>>>>>> 78b31279b37ae775c8acaf3aefc57bbc07565c90
		p.GoblinPattern(&e3, 1)
		if p.lp <= 0 {
			Slow(e3.name, 1)
			Slow(" vous a battu", 1)
			p.Dead()
		}
		time.Sleep(3 * time.Second)
	}
}

func (p *Personnage) TheSecond() {
	var e4 Monstre
	e4.InitGoblin("Java", 100, 10)
	Slow("Le Grand Java se présente devant vous\n",1)
	time.Sleep(1 * time.Second)
	var turn2 int
	for i := 0; i <= 9999; i++ {
		turn2++

		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
<<<<<<< HEAD
		
		Slow("Tour",2)
		fmt.Print(turn2)
		Slow("\nC'est au joueur !",2)
		p.CharTurn(&e4)
		if e4.lp <= 0 {
			Slow("\nVous avez vaincu le Grand ",1)
			Slow(e4.name,1)
			Slow(" !",1)
=======
		Slow("Le Grand Java se présente devant vous", 1)
		Slow("Tour", 2)
		fmt.Print(turn2)
		time.Sleep(1 * time.Second)
		Slow("C'est au joueur !", 2)
		p.CharTurn(&e4)
		if e4.lp <= 0 {
			fmt.Print("\n")
			Slow("Vous avez vaincu le Grand ", 1)
			Slow(e4.name, 1)
			Slow(" !", 1)
>>>>>>> 78b31279b37ae775c8acaf3aefc57bbc07565c90
			time.Sleep(2 * time.Second)
			p.TheThird()
		}
		time.Sleep(1 * time.Second)
<<<<<<< HEAD
		Slow("\nC'est à l'ennemi !\n",2)
		p.GoblinPattern(&e4, 1)
		if p.lp <= 0 {
			Slow(e4.name,1)
			Slow(" vous a battu",1)
=======
		fmt.Print("\n")
		Slow("C'est à l'ennemi !", 2)
		p.GoblinPattern(&e4, 1)
		if p.lp <= 0 {
			fmt.Print("\n")
			Slow(e4.name, 1)
			Slow(" vous a battu", 1)
>>>>>>> 78b31279b37ae775c8acaf3aefc57bbc07565c90
			p.Dead()
		}
		time.Sleep(3 * time.Second)
	}
}

func (p *Personnage) TheThird() {
	var e5 Monstre
	e5.InitGoblin("C++", 150, 15)
	Slow("Le Grand C++ se présente devant vous\n",1)
	time.Sleep(1 * time.Second)
	var turn3 int
	for i := 0; i <= 9999; i++ {
		turn3++

		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
<<<<<<< HEAD
		Slow("Tour",2)
		fmt.Print(turn3)
		fmt.Println("\nC'est au joueur !",2)
		p.CharTurn(&e5)
		if e5.lp <= 0 {
			Slow("\nVous avez vaincu le Grand ",1)
			Slow(e5.name,1)
			Slow(" !",1)
=======
		Slow("Le Grand C++ se présente devant vous", 1)
		Slow("Tour", 2)
		fmt.Print(turn3)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est au joueur !", 2)
		p.CharTurn(&e5)
		if e5.lp <= 0 {
			fmt.Print("\n")
			Slow("Vous avez vaincu le Grand ", 1)
			Slow(e5.name, 1)
			Slow(" !", 1)
>>>>>>> 78b31279b37ae775c8acaf3aefc57bbc07565c90
			time.Sleep(2 * time.Second)
			p.TheFourth()
		}
		time.Sleep(1 * time.Second)
		Slow("\nC'est à l'ennemi !\n", 2)
		p.GoblinPattern(&e5, 1)
		if p.lp <= 0 {
			Slow(e5.name, 1)
			Slow(" vous a battu", 1)
			p.Dead()

		}
		time.Sleep(3 * time.Second)
	}
}

func (p *Personnage) TheFourth() {
	var e6 Monstre
	e6.InitGoblin("Golang", 200, 20)
	Slow("Le plus Grand des Grand, Golang, se présente devant vous\n",1)
	time.Sleep(1 * time.Second)
	var turn4 int
	for i := 0; i <= 9999; i++ {
		turn4++

		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
<<<<<<< HEAD
		Slow("Tour",2)
		fmt.Println(turn4)
		Slow("C'est au joueur !",2)
		p.CharTurn(&e6)
		if e6.lp <= 0 {
			Slow("Vous avez vaincu Golang, le plus Grand des Grands.\n",1)
			Slow("Dans son dernier souffle il lâche un objet.",4)
			p.AddInventory("Objet suspicieux",0)
=======
		Slow("Le plus Grand des Grand, Golang, se présente devant vous", 1)
		Slow("Tour", 2)
		fmt.Print(turn4)
		time.Sleep(1 * time.Second)
		Slow("C'est au joueur !", 2)
		p.CharTurn(&e6)
		if e6.lp <= 0 {
			Slow("Vous avez vaincu Golang, le plus Grand des Grands.", 1)
			fmt.Print("\n")
			Slow("Dans son dernier souffle il lâche un objet.", 4)
			p.AddInventory("Objet suspicieux", 0)
>>>>>>> 78b31279b37ae775c8acaf3aefc57bbc07565c90
			time.Sleep(2 * time.Second)
			break
		}
		time.Sleep(1 * time.Second)
<<<<<<< HEAD
		Slow("\nC'est à l'ennemi !\n",2)
=======
		fmt.Print("\n")
		Slow("C'est à l'ennemi !", 2)
>>>>>>> 78b31279b37ae775c8acaf3aefc57bbc07565c90
		p.GoblinPattern(&e6, 1)
		if p.lp <= 0 {
			Slow(e6.name, 1)
			Slow(" vous a battu", 1)
			p.Dead()
			break
		}
		time.Sleep(3 * time.Second)
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
