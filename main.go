package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// fonction qui execute nos sous fonctions et rentre les valeur ainsi que le menu principal
	var p1 Personnage
	p1.CharCreation()
	fmt.Printf("Vous allez être redirigé vers le menu principal")
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

	fmt.Printf("Bienvenue dans le menu de création de personnage \nPour commencer, choisissez un nom pour votre avatar: \n")
	fmt.Scanln(&name)
	fmt.Println("Votre personnage se nomme donc", name)
	time.Sleep(3 * time.Second)
	fmt.Println("\nChoisissez maintenant la race de ", name, "parmi:\n-Humain\n-Elfe\n-Nain")
	fmt.Scanln(&class)
	if class != "Humain" && class != "Elfe" && class != "Nain" {
		fmt.Println("Erreur, veuillez entrer une valeur correcte:\nHumain, Elfe ou Nain")
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
	if class != "Humain" && class != "Elfe" && class != "Nain" {
		class = "Troll"
		lpmax = 50
	}
	lp = lpmax / 2
	money = 100
	inventory = []string{"Potion de vie", "Potion de vie", "Potion de vie"}
	skill = []string{"Coup de poing"}
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(name, "est donc un", class, ", il commence avec", lp, " point de vie et", lpmax, "point de vie maximum.")
	fmt.Println(name, "est niveau 1, possède le sort Coup de poing et a", money, "pièces d'or.")
	level = 1
	p.Init(name, class, level, lpmax, lp, inventory, skill, money)
}

func (p *Personnage) menu() {
	var menu int
	fmt.Println("+++++++++++++++++++++++++++++++")
	fmt.Println("A quoi voulez vous accéder:")
	fmt.Println("----- \n Afficher les informations du personnage (1)")
	fmt.Println("----- \n Accéder au contenu de l’inventaire (2)")
	fmt.Println("----- \n Voir le Marchand (3)")
	fmt.Println("----- \n Aller parler au Forgeron (4)")
	fmt.Println("----- \n S'entrainer contre un gobelin (5)")
	fmt.Println("----- \n Quitter (6) \n-----")
	fmt.Println("Entrez le numéro de l'option:")
	fmt.Println("+++++++++++++++++++++++++++++++")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		p.DisplayInfo()
	case 2:
		p.AccessInventory()
	case 3:
		fmt.Println("+++++++++++++++++++++Marchand+++++++++++++++++++")
		fmt.Println("-----\nPotion de vie --> 15 pièces <--(1)")
		fmt.Println("-----\nPotion de poison --> 20 pièces <-- (2)")
		fmt.Println("-----\nLivre de sort:Boule de Feu --> 25 pièces <-- (3)")
		fmt.Println("-----\nFourrure de Loup --> 5 pièces <--(4)")
		fmt.Println("-----\nPeau de Troll --> 6 pièces <--(5)")
		fmt.Println("-----\nCuir de Sanglier --> 4 pièces <--(6)")
		fmt.Println("-----\nPlume de Corbac --> 6 pièces <--(7)")
		fmt.Println("-----\nRetour au menu principal (8) \n--")
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
		p.Marchand()
	case 4:
		fmt.Println("------ Vous entrez dans la forge ------")
		fmt.Println("--\n Construire un Chapeau de l'aventurier (1) --")
		fmt.Println("requiert 1 plume de corbeau et 1 cuir de sanglier")
		fmt.Println("--\n Construire une Tunique de l'Aventurier (2) --")
		fmt.Println("requiert 2 Fourrure de Loup et 1 Peau de Troll")
		fmt.Println("--\n Construire les bottes de l'aventurier (3) --")
		fmt.Println("requiert 1 Fourrure de Loup et 1 Cuir de Sanglier")
		p.Forgeron()
	case 5:
		p.TrainingFight()
	}
}

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

func (p *Personnage) AccessInventory() {
	// fonction qui nous permet d'acceder a notre inventaire
	// var rep int
	if len(p.inventory) == 0 {
		fmt.Println("inventaire vide fraté")
	} else {
		fmt.Println(p.inventory)
	}
	for i := 0; i < len(p.inventory); i++ {
		if p.inventory[i] != " " {
			fmt.Println("---]", p.inventory[i], "[---")
		}
	}
	fmt.Println("___________________________________________")
	fmt.Println("Prendre une Potion de vie (1)")
	fmt.Println("Prendre une Potion de poison (2)")
	fmt.Println("Utiliser le Livre de sort : Boule de Feu (3)")
	fmt.Println("Menu précédent (4)")
	fmt.Println("___________________________________________")
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
	case 4:
		p.menu()
	}
}

func (p *Personnage) Spellbook() {
	// fonction qui nous permet d'ajouter ou repertorier les sorts (spell)
	for _, letter := range p.skill {
		if letter == ("Boule de feu") {
			fmt.Println("dsl t'a déja les boules")
		} else {
			p.skill = append(p.skill, "Boule de feu")
		}
	}
	p.menu()
}

func (p *Personnage) Marchand() {
	// fonction affichant le menu du marchand , et les ajoute a notre inventaire
	var menum int
	fmt.Scanln(&menum)
	switch menum {
	case 1:
		p.AddInventory("Potion de vie", 15)
		fmt.Println("il vous reste", p.money, "pièces")
		p.Marchand()
	case 2:
		p.AddInventory("Potion de poison", 20)
		fmt.Println("il vous reste", p.money, "pièces")
		p.Marchand()
	case 3:
		p.AddInventory("Livre de sort: Boule de Feu", 25)
		fmt.Println("il vous reste", p.money, "pièces")
		p.Marchand()
	case 4:
		p.AddInventory("Fourrure de Loup", 5)
		fmt.Println("il vous reste", p.money, "pièces")
		p.Marchand()
	case 5:
		p.AddInventory("Peau de Troll", 6)
		fmt.Println("il vous reste", p.money, "pièces")
		p.Marchand()
	case 6:
		p.AddInventory("Cuir de Sanglier", 4)
		fmt.Println("il vous reste", p.money, "pièces")
		p.Marchand()
	case 7:
		p.AddInventory("Plume de Corbac", 6)
		fmt.Println("il vous reste", p.money, "pièces")
		p.Marchand()
	case 8:
		p.menu()
	}
}

func (p *Personnage) DisplayInfo() {
	// fonction nous permettant de voir les informations de notre personnage
	fmt.Println("-------------------------")
	fmt.Println("Nom:", p.name)
	fmt.Println("Classe:", p.class)
	fmt.Println("Niveau:", p.level)
	fmt.Println("Vie maximum:", p.lpmax)
	fmt.Println("Vie actuelle:", p.lp)
	fmt.Println("Contenu de l'inventaire:", p.inventory)
	fmt.Println("skill :", p.skill)
	fmt.Println("pessos :", p.money)
	fmt.Println("-------------------------")
	p.menu()
}

func (p *Personnage) AddInventory(item string, price int) {
	if p.money >= price {
		p.money -= price
		p.inventory = append(p.inventory, item)
	} else {
		fmt.Println("Tu n'a pas assez d'argent pour acheter cet objet")
	}
}

func (p *Personnage) TakePot() {
	// fonction qui nous permet de prendre une potion de soin
	for _, letter := range p.inventory {
		if letter == "Potion de vie" {
			if p.lp <= (p.lpmax - 50) {
				p.lp += 50
				fmt.Println("Glou glou glou, ça fait du bien")
				p.inventory[len(p.inventory)-1] = ""
				break
			} else if p.lp > (p.lpmax-50) && p.lp < p.lpmax {
				p.lp = p.lpmax
				fmt.Println("Glou glou glou, ça fait du bien")
				p.inventory[len(p.inventory)-1] = ""
				break
			} else {
				fmt.Println("Vous êtes full, vous ne pouvez pas utiliser la potion")
				break
			}
		}
	}
}

func (p *Personnage) Dead() {
	// fonction qui verifie si le personnage est mort et le ressucite a la moitié de ses pv
	if p.lp == 0 {
		fmt.Printf("Bravo, vous êtes mort. \n")
		time.Sleep(2 * time.Second)
		fmt.Printf("Mais ne paniquez pas, vous allez être ressuciter \n")
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

func (p *Personnage) PoisonPot() {
	// fonction qui crée la potion poison et explique ce qu'elle fait sur un personnage
	for _, letter := range p.inventory {
		if letter == "Potion de poison" {
			fmt.Println("Vous buvez la potion de poison, ouch")
			p.inventory[len(p.inventory)-1] = ""
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
		}
	}
}

func (p *Personnage) Forgeron() {
	var enclume int
	fmt.Scanln(&enclume)
	switch enclume {
	case 1:
		for _, letter := range p.inventory {
			if letter != "Plume de Cobrac" && letter != "Cuir de Sanglier" {
				fmt.Println("vous n'avez pas les ressources nécéssaires !")
			} else {
				p.AddInventory("Chapeau de l'Aventurier", 5)
			}
		}
	case 2:
		for _, letter := range p.inventory {
			if letter != "Fourrure de Loup" && letter != "Peau de Troll" {
				fmt.Println("vous n'avez pas les ressources nécéssaires !")
			} else {
				p.AddInventory("Tunique de l'Aventurier", 5)
			}
		}
	case 3:
		for _, letter := range p.inventory {
			if letter != "Fourrure de Loup" && letter != "Cuir de Sanglier" {
				fmt.Println("vous n'avez pas les ressources nécéssaires !")
			} else {
				p.AddInventory("Bottes de l'Aventurier", 5)
			}
		}
	}
	p.menu()
}

type Monstre struct {
	name   string
	lp     int
	lpmax  int
	attack int
}

func (m *Monstre) InitGoblin(name string, lpmax int, attack int) {
	// initialisation de notre personnage
	m.name = name
	m.lpmax = lpmax
	m.lp = lpmax / 2
	m.attack = attack
}

func (p *Personnage) GoblinPattern(m *Monstre) {
	p.lp -= m.attack
	fmt.Println(m.name, "attaque", p.name, "et lui inflige")
}

func (p *Personnage) CharTurn() {
	var choice int
	fmt.Println("Attaquer (1)")
	fmt.Println("Utiliser un objet (2)")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println(p.name, "utilise Attaque Basique et inflige 5 points de dégâts à Y, il lui reste X PV")
	case 2:
		for i := 0; i < len(p.inventory); i++ {
			if p.inventory[i] != " " {
				fmt.Println("---]", p.inventory[i], "[---")

			}
		}
	}
}

func (p *Personnage) TrainingFight() {
	var e1 Monstre
	e1.InitGoblin("Gobelin d'entrainement", 40, 5)
	var turn int
	for i := 0; i <= 9999; i++ {
		turn++
		fmt.Println("Tour", turn)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est au joueur !")
		time.Sleep(1 * time.Second)
		fmt.Println("C'est à l'ennemi !")
		p.GoblinPattern(&e1)
		time.Sleep(1 * time.Second)
	}
}
