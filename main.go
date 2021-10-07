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
	Slow("Vous vous rendez sur la place principale",1)
	time.Sleep(1 * time.Second)
	fmt.Printf(". ")
	time.Sleep(1 * time.Second)
	fmt.Printf(". ")
	time.Sleep(1 * time.Second)
	fmt.Printf(". \n")
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
	Slow("Bienvenue dans",1)
	fmt.Printf(Yellow +" ")
	Slow("Goldy",1)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("\n" + Reset)
	Slow("Vous êtes dans le menu de création de personnage\n",1)
	Slow("Pour commencer, choisissez un nom pour votre avatar: \n",1)
	fmt.Scanln(&name)
	if name == "Utilisateur" { //Easter egg n°1, le mode développeur du jeu
		class = "Utilisateur"
		level = 1
		lpmax = 9999
		lp = lpmax
		inventory = []string{"Gantelet de l'infini"}
		money = 10000
		Slow("Bienvenue Utilisateur\n",1)
		time.Sleep(300 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	} else {
		Slow("Votre personnage se nomme donc ",1)
		Slow(name,1)
		time.Sleep(2 * time.Second)
		fmt.Println("")
		Slow("Choisissez maintenant la race de ",1)
		Slow(name,1)
		Slow(" parmi:\n",1)
		Slow("\n-Humain\n",1)
		Slow("\n-Elfe\n",1)
		Slow("\n-Nain\n \n",1)
		fmt.Scanln(&class)
		if class != "Humain" && class != "Elfe" && class != "Nain" {
			Slow("Erreur, veuillez entrer une valeur correcte:\n",1)
			Slow("Humain, Elfe ou Nain\n",1)
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
		Slow(name,1)
		Slow(" est donc un ",1)
		Slow(class,1)
		Slow(", il commence avec ",1)
		fmt.Print(lp)
		Slow(" point de vie et ",1)
		fmt.Print(lpmax)
		Slow(" points de vie maximum.",1)
		fmt.Println("")
		if class == "Troll" {
			Slow("Ça t'apprendras à pas savoir écrire\n",1)
		}
		Slow(name,1)
		Slow(" est niveau 1, possède le sort Coup de poing et a ",1)
		fmt.Printf(Yellow +"")
		fmt.Print(money)
		Slow(" pièces",1)
		fmt.Println(""+ Reset)
		level = 1
		time.Sleep(3 * time.Second)
	}
	Slow("Lancement de la simulation\n",1)
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
	Slow("\n+++++++++++++++++++++++++++++++\n",3)
	Slow("Que voulez vous faire ?\n",3)
	Slow("----- \n",3)
	fmt.Print(Yellow + "")
	Slow("(1) ",3)
	fmt.Print(""+ Reset)
	Slow("Regarder la carte ",3)
	fmt.Print(Yellow + "")
	Slow("d'Identité\n",3)
	fmt.Print("" + Reset)
	Slow("----- \n",3)
	fmt.Print(Yellow + "")
	Slow("(2) ",3)
	fmt.Print(""+ Reset)
	Slow("Jeter un oeil dans le ",3)
	fmt.Print(Yellow + "")
	Slow("Sac\n",3)
	fmt.Print(""+ Reset)
	Slow("----- \n",3)
	fmt.Print(Yellow + "")
	Slow("(3) ",3)
	fmt.Print(""+ Reset)
	Slow("Se rendre dans la boutique du ",3)
	fmt.Print(Yellow + "")
	Slow("Marchand\n",3)
	fmt.Print(""+ Reset)
	Slow("----- \n",3)
	fmt.Print(Yellow + "")
	Slow("(4) ",3)
	fmt.Print(""+ Reset)
	Slow("Aller à la ",3)
	fmt.Print(Yellow + "")
	Slow("Forge\n",3)
	fmt.Print(""+ Reset)
	Slow("----- \n",3)
	fmt.Print(Yellow + "")
	Slow("(5) ",3)
	fmt.Print(""+ Reset)
	Slow("Partir en exploration dans le ",3)
	fmt.Print(Yellow + "")
	Slow("Donjon\n",3)
	fmt.Print(""+ Reset)
	Slow("----- \n",3)
	fmt.Print(Yellow +"")
	Slow("(6) ",3)
	fmt.Print(""+ Reset)
	Slow("Affronter les ",3)
	fmt.Print(Yellow + "")
	Slow("Quatres Grands\n",3)
	fmt.Print(""+ Reset)
	Slow("----- \n",3)
	fmt.Print(Yellow + "")
	Slow("(0) ",3)
	fmt.Print(""+ Reset)
	Slow("Faire un ",3)
	fmt.Print(Yellow + "")
	Slow("somme",3)
	fmt.Print(""+ Reset)
	Slow("\n-----\n",3)
	Slow("Entrez le numéro de l'option:\n",3)
	Slow("+++++++++++++++++++++++++++++++\n",3)
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.DisplayInfo()
		fmt.Print(Yellow + "")
		Slow("(0) ", 1)
		fmt.Print(""+ Reset)
		Slow("Arrêter de regarder la ",1)
		fmt.Print(Yellow +"")
		Slow("carte\n",1)
		fmt.Print(""+ Reset)
		fmt.Scanln(&retour)
		switch retour {
		case 0:
			p.menu()
		}
	case 2:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.AccessInventory()
	case 3:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		Slow("+++++++++++++++++++++Marchand+++++++++++++++++++\n",2)
		Slow("-----\n",2)
		fmt.Print(Yellow + "")
		Slow(" (1) Potion de vie ",2)
		fmt.Print(""+ Reset)
		Slow("-->",2)
		fmt.Print(Yellow + "")
		Slow("15 pièces",2)
		fmt.Print("" + Reset)
		Slow("<--",2)
		Slow("-----\n",2)
		fmt.Print(Yellow + "")
		Slow("(2) Potion de poison",2)
		fmt.Print(""+ Reset)
		Slow("-->",2)
		fmt.Print(Yellow + "")
		Slow("20 pièces",2)
		fmt.Print(""+ Reset)
		Slow("<--",2)
		Slow("-----\n",2)
		fmt.Print(Yellow + "")
		Slow("(3) Livre de sort:Boule de Feu",2)
		fmt.Print(""+ Reset)
		Slow("-->",2)
		fmt.Print(Yellow + "")
		Slow("25 pièces",2)
		fmt.Print(""+ Reset)
		Slow("<--",2)
		Slow("-----\n",2)
		fmt.Print(Yellow + "")
		Slow("(4) Fourrure de Loup",2)
		fmt.Print(""+ Reset)
		Slow("-->",2)
		fmt.Print(Yellow + "")
		Slow("5 pièces", 2)
		fmt.Print(""+ Reset)
		Slow("<--",2)
		Slow("-----\n",2)
		fmt.Print(Yellow + "")
		Slow("(5) Peau de Troll",2)
		fmt.Print(""+ Reset)
		Slow("-->",2)
		fmt.Print(Yellow + "")
		Slow("6 pièces",2)
		fmt.Print(""+ Reset)
		Slow("<--",2)
		Slow("-----\n",2)
		fmt.Print(Yellow + "")
		Slow("(6) Cuir de Sanglier",2)
		fmt.Print(""+ Reset)
		Slow("-->",2)
		fmt.Print(Yellow + "")
		Slow("4 pièces",2)
		fmt.Print(""+ Reset)
		Slow("<--",2)
		Slow("-----\n",2)
		fmt.Print(Yellow + "")
		Slow("(7) Plume de Corbac",2)
		fmt.Print(""+ Reset)
		Slow("-->",2)
		fmt.Print(Yellow +"")
		Slow("6 pièces",2)
		fmt.Print(""+ Reset)
		Slow("<--",2)
		Slow("-----\n",2)
		fmt.Print(Yellow + "")
		Slow("(0)" ,2)
		fmt.Print(""+ Reset)
		Slow("Retour la",2)
		fmt.Print(Yellow + "")
		Slow("place centrale",2)
		fmt.Print(""+ Reset)
		Slow("\n--",2)
		Slow("++++++++++++++++++++++++++++++++++++++++++++++++",2)
		p.Marchand()
	case 4:
		p.Forgeron()
	case 5:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.TrainingFight()
		p.menu()
	case 6:
		p.TheFirst()
	case 0:
		fmt.Println("Fin de la transmission")
		break
	case 7171:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		Slow("Bravo",1)
		fmt.Printf(p.name)
		Slow("tu as réussi à triompher des quatres Grands",1)
		fmt.Printf("\n")
		Slow("Ta quête s'arrête ici mon ami, mais nos voyagent se recroiseront, soit en sur.",1)

	}
}

func (p *Personnage) DisplayInfo() {
	// fonction nous permettant de voir les informations de notre personnage
	Slow("-------------------------\n",2)
	fmt.Print(Yellow + "")
	Slow("Nom: ",2)
	fmt.Print(""+ Reset)
	Slow(p.name,2)
	fmt.Print(Yellow + "")
	Slow("\nClasse: ",2)
	fmt.Print(""+ Reset)
	Slow(p.class,2)
	fmt.Print(Yellow + "")
	Slow("\nNiveau: ",2)
	fmt.Print(""+ Reset)
	fmt.Print(p.level)
	fmt.Print(Yellow + "")
	Slow("\nVie maximum: ",2)
	fmt.Print(""+ Reset)
	fmt.Print(p.lpmax)
	fmt.Print(Yellow + "")
	Slow("\nVie actuelle: ",2)
	fmt.Print(""+ Reset)
	fmt.Print(p.lp)
	fmt.Print(Yellow + "")
	Slow("\nContenu de l'inventaire: ",2)
	fmt.Print(""+ Reset)
	fmt.Print(p.inventory)
	fmt.Print(Yellow + "")
	Slow("\nCompétences: ",2)
	fmt.Print(""+ Reset)
	fmt.Print(p.skill)
	fmt.Print(Yellow + "")
	Slow("\nPièces: ",2)
	fmt.Print(""+ Reset)
	fmt.Print(p.money)
	Slow("\n-------------------------\n",2)
}

func (p *Personnage) AccessInventory() {
	// fonction qui nous permet d'acceder a notre inventaire
	if len(p.inventory) == 0 {
		Slow("\nLa sacoche est ",1)
		fmt.Print(Yellow + "")
		Slow("vide",1)
		fmt.Print(""+ Reset)
	} else {
		fmt.Print("\n")
		for i := 0; i < len(p.inventory); i++ {
			if p.inventory[i] != " " {
				Slow("---] ",2) 
				fmt.Print(""+ Yellow, p.inventory[i], Reset +"")
				Slow(" [---",2)
				fmt.Print("\n")
			}
		}
		fmt.Printf("\n")
	}
	Slow("----------------\n",2)
	fmt.Print(Yellow + "")
	Slow("(1) ",2)
	fmt.Print("" + Reset)
	Slow("Prendre une ",2)
	fmt.Print(Yellow + "")
	Slow("Potion de vie\n",2)
	Slow("(2) ",2)
	fmt.Print("" + Reset)
	Slow("Prendre une ",2)
	fmt.Print(Yellow + "")
	Slow("Potion de poison\n",2)
	Slow("(3) ",2)
	fmt.Print("" + Reset)
	Slow("Utiliser le Livre de sort : ",2)
	fmt.Print(Yellow + "")
	Slow("Boule de Feu\n",2)
	Slow("(4) ",2)
	fmt.Print("" + Reset)
	Slow("Mettre le ",2)
	fmt.Print(Yellow + "")
	Slow("Chapeau\n",2)
	Slow("(5) ",2)
	fmt.Print("" + Reset)
	Slow("Mettre la ",2)
	fmt.Print(Yellow + "")
	Slow("Tunique\n",2)
	Slow("(6) ",2)
	fmt.Print("" + Reset)
	Slow("Mettre les ",2)
	fmt.Print(Yellow + "")
	Slow("Bottes\n",2)
	Slow("(0) ",2)
	fmt.Print(""+ Reset)
	Slow("Arrêter de regarder dans le ",2)
	fmt.Print(Yellow + "")
	Slow("Sac\n",2)
	fmt.Print(""+ Reset)
	Slow("----------------\n",2)
	p.UseInventory()
}

func (p *Personnage) UseInventory() {
	// fonction qui nous permet d'interagir avec les éléments de l'inventaire
	var use int
	fmt.Scanln(&use)
	switch use {
	case 1:
		p.TakePot()
		p.UseInventory()
	case 2:
		p.PoisonPot()
		p.UseInventory()
	case 3:
		p.Spellbook()
		p.UseInventory()
	case 4:
		if p.Checkinv("Chapeau de l'aventurier") {
			p.Chapeau = "Chapeau de l'aventurier"
			p.RemoveInventory("Chapeau de l'Aventurier")
			fmt.Println("Vous avez désormais +15 hp avec le couvre chef !")
			p.UseInventory()
		} else if !p.Checkinv("Chapeau de L'aventurier") {
			fmt.Println("Tu l'as déja équiper !")
			p.UseInventory()
		}
	case 5:
		if !p.Checkinv("Tunique de l'Aventurier") {
			p.Tunique = "Tunique de l'Aventurier"
			p.lpmax += 20
			fmt.Println("Vous avez désormais +20 hp grace a la tunique !")
			p.RemoveInventory("Tunique de l'Aventurier")
			p.UseInventory()
		} else {
			fmt.Println("Tu l'as déja équiper !")
			p.UseInventory()
		}
	case 6:
		if !p.Checkinv("Bottes de l'Aventurier") {
			p.Bottes = "Bottes de l'Aventurier"
			p.lpmax += 10
			fmt.Println("Vous avez désormais +10 hp grace a la paire !")
			p.RemoveInventory("Bottes de l'Aventurier")
			p.UseInventory()
		} else {
			fmt.Println("Tu l'as déja équiper !")
			p.UseInventory()
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
				fmt.Println("   :~:  ")
				fmt.Println("   | |    ")
				fmt.Println("  .' `.   	 ")
				fmt.Println(".'     `. ")
				fmt.Println("|       | ")
				fmt.Println(" `.._..' ")
				fmt.Println("Glou glou glou, ça fait du bien")
				p.RemoveInventory("Potion de vie")
				break
			} else if p.lp > (p.lpmax-50) && p.lp < p.lpmax {
				p.lp = p.lpmax
				fmt.Println("Glou glou glou, ça fait du bien")
				p.RemoveInventory("Potion de vie")
				break
			} else {
				fmt.Println("Vous êtes full, vous ne pouvez pas utiliser la potion")
				break
			}
		} else {
			Slow("T'as pas l'objet fils de tes parents",2)
		}
	}
}

func (p *Personnage) PoisonPot() {
	// fonction qui crée la potion poison et explique ce qu'elle fait sur un personnage
	for _, letter := range p.inventory {
		if letter == "Potion de poison" {
			fmt.Println("Vous buvez la potion de poison, ouch")
			p.RemoveInventory("Potion de poison")
			time.Sleep(1 * time.Second)
			fmt.Println(p.lp, "/", p.lpmax, "PV")
			p.lp -= 10
			fmt.Println(p.lp, "/", p.lpmax, "PV")
			time.Sleep(1 * time.Second)
			p.lp -= 10
			fmt.Println(p.lp, "/", p.lpmax, "PV")
			time.Sleep(1 * time.Second)
			p.lp -= 10
			fmt.Println(p.lp, "/", p.lpmax, "PV")
		} else {
			Slow("T'as pas l'objet fils de tes parents",2)
		}
	}
}

func (p *Personnage) Dead() {
	// fonction qui verifie si le personnage est mort et le ressucite a la moitié de ses pv
	if p.lp == 0 {
		fmt.Printf("Bravo, vous êtes mort. \n")
		time.Sleep(2 * time.Second)
		fmt.Printf("Mais ne paniquez pas, vous allez être ressucité \n")
		time.Sleep(2 * time.Second)
		fmt.Printf("Manoeuvre de réanimation en cours")
		time.Sleep(1 * time.Second)
		fmt.Printf(". ")
		time.Sleep(1 * time.Second)
		fmt.Printf(". ")
		time.Sleep(1 * time.Second)
		fmt.Printf(". \n")
		time.Sleep(1 * time.Second)
		p.lp = p.lpmax / 2
		p.DisplayInfo()
	}
}

func (p *Personnage) Spellbook() {
	// fonction qui nous permet d'ajouter ou repertorier les sorts (spell)
	for _, letter := range p.skill {
		if letter == ("Boule de feu") {
			fmt.Println("Tu as déjà les boules")

		} else {
			p.skill = append(p.skill, "Boule de feu")
		}
	}
	p.menu()
}

func (p *Personnage) Marchand() {
	// fonction affichant le menu du marchand , et les ajoute a notre inventaire
	var menum int
	fmt.Println("Vous avez", p.money, "pièces")
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
		fmt.Println("Tu n'a pas assez d'argent pour acheter cet objet")
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
		fmt.Println("------ Vous entrez dans la forge ------")
		fmt.Println("\n-- Construire un Chapeau de l'aventurier (1) --")
		fmt.Println("requiert 1 plume de corbeau et 1 cuir de sanglier")
		fmt.Println("\n-- Construire une Tunique de l'Aventurier (2) --")
		fmt.Println("requiert 2 Fourrure de Loup et 1 Peau de Troll")
		fmt.Println("\n-- Construire les bottes de l'aventurier (3) --")
		fmt.Println("requiert 1 Fourrure de Loup et 1 Cuir de Sanglier")
		fmt.Printf("\n-- Retourner au menu principal (0) --\n")
		fmt.Println("\nVotre inventaire:")
		for i := 0; i < len(p.inventory); i++ {
			if p.inventory[i] != " " {
				fmt.Println("---]", p.inventory[i], "[---")
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
			fmt.Println("Vous êtes désormais en possession de Chapeau de l'aventurier")
			time.Sleep(2 * time.Second)
			p.Forgeron()
		} else {
			fmt.Println("Tu te moques de moi ? Regarde ton inventaire l'ami")
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
				fmt.Println("Vous êtes désormais en possession de la Tunique de l'aventurier")
				time.Sleep(2 * time.Second)
				p.Forgeron()
			} else {
				p.AddInventory("Fourrure de Loup", 5)
			}
		} else {
			fmt.Println("Tu te moques de moi ? Regarde ton inventaire l'ami")
			time.Sleep(2 * time.Second)
			p.Forgeron()
		}
	case 3:
		if p.Checkinv("Fourrure de Loup") && p.Checkinv("Cuir de Sanglier") {
			p.RemoveInventory("Fourrure de Loup")
			p.RemoveInventory("Cuir de Sanglier")
			p.AddInventory("Bottes de l'Aventurier", 0)
			fmt.Println("Vous êtes désormais en possession des Bottes de l'aventurier")
			time.Sleep(2 * time.Second)
			p.Forgeron()
		} else {
			fmt.Println("Tu te moques de moi ? Regarde ton inventaire l'ami")
			time.Sleep(2 * time.Second)
			p.Forgeron()
		}
	case 0:
		p.menu()
	}
}

func (p *Personnage) Checkinv(item string) bool {
	var founditem bool
	for _, letter := range p.inventory {
		if letter == item {
			founditem = true
		}
	}
	if founditem {
		return true
	} else {
		return false
	}
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
	fmt.Println(m.name, "attaque", p.name, "et lui inflige", m.attack, " points de dégâts, il lui reste", p.lp, "PV")

}

func (p *Personnage) CharTurn(m *Monstre) {
	var choice int
	var damage int = 5
	fmt.Println("Attaquer (1)")
	fmt.Println("Utiliser un objet (2)")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		if p.RealKnife() == true {
			m.lp = 0
		} else if p.name == "Utilisateur" {
			m.lp -= damage * 10
		} else {
			m.lp -= damage
		}
		fmt.Println(p.name, "utilise Coup de poing et inflige", damage, "points de dégâts à", m.name, "il lui reste", m.lp, "PV")
	case 2:
		for i := 0; i < len(p.inventory); i++ {
			if p.inventory[i] != " " {
				fmt.Println("---]", p.inventory[i], "[---")

			}
		}
		fmt.Println("Prendre une potion de soin (1)")
		fmt.Println("Prendre une potion de poison (2)")
		fmt.Println("Retour (0)")
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
		fmt.Println("Un mimic sauvage apparaît")
		time.Sleep(3 * time.Second)
		var turn int
		e2.InitGoblin("Mimic", 80, 5)
		for i := 0; i <= 9999; i++ {
			turn++
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			fmt.Println("Tour", turn)
			time.Sleep(1 * time.Second)
			fmt.Println("C'est au joueur !")
			p.CharTurn(&e2)
			if e2.lp <= 0 {
				var mimic string = "Vous avez vaincu le mimic "
				for _, letter := range mimic {
					time.Sleep(25 * time.Millisecond)
					z01.PrintRune(letter)
				}
				time.Sleep(1 * time.Second)
				var strange1 string = "vous le fouillez"
				var strange2 string = " et obtenez "
				var strange3 string = "un objet étrange"
				for _, letter := range strange1 {
					time.Sleep(100 * time.Millisecond)
					z01.PrintRune(letter)
				}
				for _, letter := range strange2 {
					time.Sleep(200 * time.Millisecond)
					z01.PrintRune(letter)
				}
				for _, letter := range strange3 {
					time.Sleep(300 * time.Millisecond)
					z01.PrintRune(letter)
				}
				time.Sleep(3 * time.Second)
				if p.Checkinv("Véritable couteau") == false {
					p.AddInventory("Véritable couteau", 0)
				}
				break
			}
			time.Sleep(1 * time.Second)
			fmt.Println("\nC'est à l'ennemi !")
			p.GoblinPattern(&e2, 1)
			if p.lp <= 0 {
				fmt.Println(e2.name, "vous a battu")
				p.Dead()
				break
			}
			time.Sleep(3 * time.Second)
		}
	} else {
		e1.InitGoblin("Gobelin d'entrainement", 40, 5)
		fmt.Println("Le", e1.name, "estprêt à se battre")
		time.Sleep(3 * time.Second)
		var turn int
		for i := 0; i <= 9999; i++ {
			turn++
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			fmt.Println("Tour", turn)
			time.Sleep(1 * time.Second)
			fmt.Println("C'est au joueur !")
			p.CharTurn(&e1)
			if e1.lp <= 0 {
				fmt.Println("\nVous avez vaincu", e1.name)
				time.Sleep(2 * time.Second)
				break
			}
			time.Sleep(1 * time.Second)
			fmt.Println("\nC'est à l'ennemi !")
			if turn%3 == 3 || turn == 3 {
				p.GoblinPattern(&e1, 2)
			} else {
				p.GoblinPattern(&e1, 1)
			}
			if p.lp <= 0 {
				fmt.Println(e1.name, "vous a battu")
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
	var turn int
	for i := 0; i <= 9999; i++ {
		turn++
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Le Grand Python se présente devant vous")
		fmt.Println("Tour", turn)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est au joueur !")
		p.CharTurn(&e3)
		if e3.lp <= 0 {
			fmt.Println("\nVous avez vaincu le Grand", e3.name, "!")
			time.Sleep(2 * time.Second)
			p.TheSecond()
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("\nC'est à l'ennemi !")
		p.GoblinPattern(&e3, 1)
		if p.lp <= 0 {
			fmt.Println(e3.name, "vous a battu")
			p.Dead()
			break
		}
		time.Sleep(3 * time.Second)
	}
}

func (p *Personnage) TheSecond() {
	var e4 Monstre
	e4.InitGoblin("Java", 100, 10)
	var turn2 int
	for i := 0; i <= 9999; i++ {
		turn2++

		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Le Grand Java se présente devant vous")
		fmt.Println("Tour", turn2)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est au joueur !")
		p.CharTurn(&e4)
		if e4.lp <= 0 {
			fmt.Println("\nVous avez vaincu le Grand", e4.name, "!")
			time.Sleep(2 * time.Second)
			p.TheThird()
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("\nC'est à l'ennemi !")
		p.GoblinPattern(&e4, 1)
		if p.lp <= 0 {
			fmt.Println(e4.name, "vous a battu")
			p.Dead()
			break
		}
		time.Sleep(3 * time.Second)
	}
}

func (p *Personnage) TheThird() {
	var e5 Monstre
	e5.InitGoblin("C++", 150, 15)
	var turn3 int
	for i := 0; i <= 9999; i++ {
		turn3++

		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Le Grand C++ se présente devant vous")
		fmt.Println("Tour", turn3)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est au joueur !")
		p.CharTurn(&e5)
		if e5.lp <= 0 {
			fmt.Println("\nVous avez vaincu le Grand", e5.name, "!")
			time.Sleep(2 * time.Second)
			p.TheFourth()
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("\nC'est à l'ennemi !")
		p.GoblinPattern(&e5, 1)
		if p.lp <= 0 {
			fmt.Println(e5.name, "vous a battu")
			p.Dead()
			break
		}
		time.Sleep(3 * time.Second)
	}
}

func (p *Personnage) TheFourth() {
	var e6 Monstre
	e6.InitGoblin("Golang", 200, 20)
	var turn4 int
	for i := 0; i <= 9999; i++ {
		turn4++

		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Le plus Grand des Grand, Golang, se présente devant vous")
		fmt.Println("Tour", turn4)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est au joueur !")
		p.CharTurn(&e6)
		if e6.lp <= 0 {
			var Golang1 string = "Vous avez vaincu Golang, le plus Grand des Grands."
			for _, letter := range Golang1 {
				time.Sleep(50 * time.Millisecond)
				z01.PrintRune(letter)
			}
			fmt.Printf("\n")
			var Golang2 string = "Dans son dernier souffle il vous murmure ces quelques chiffres:"
			for _, letter := range Golang2 {
				time.Sleep(50 * time.Millisecond)
				z01.PrintRune(letter)
			}
			var lastwords string = "7 1 7 1"
			for _, letter := range lastwords{
				time.Sleep(200 * time.Millisecond)
				z01.PrintRune(letter)
			}
			time.Sleep(2 * time.Second)
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("\nC'est à l'ennemi !")
		p.GoblinPattern(&e6, 1)
		if p.lp <= 0 {
			fmt.Println(e6.name, "vous a battu")
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
	}
}