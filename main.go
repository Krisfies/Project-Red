package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/01-edu/z01"
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
	if name == "Utilisateur" { //Easter egg n°1, un peu le mode développeur du jeu
		class = "Utilisateur"
		level = 1
		lpmax = 9999
		lp = lpmax
		inventory = []string{"Gantelet de l'infini"}
		money = 10000
		var welcoming string = "Bienvenue Utilisateur"
		for _, letter := range welcoming {
			time.Sleep(100 * time.Millisecond)
			z01.PrintRune(letter)
		}
		time.Sleep(200 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	} else {
		fmt.Println("Votre personnage se nomme donc", name)
		time.Sleep(2 * time.Second)
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
		if class != "Humain" && class != "Elfe" && class != "Nain" { //Easter egg n°3
			class = "Troll"
			lpmax = 50
		}
		lp = lpmax / 2
		money = 100
		inventory = []string{"Potion de vie", "Potion de vie", "Potion de vie"}
		skill = []string{"Coup de poing"}
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println(name, "est donc un", class, ", il commence avec", lp, " point de vie et", lpmax, "point de vie maximum.")
		if class == "Troll" {
			fmt.Println("Ça t'apprendra à faire le malin")
		}
		fmt.Println(name, "est niveau 1, possède le sort Coup de poing et a", money, "pièces d'or.")
		level = 1
		time.Sleep(3 * time.Second)
	}
	p.Init(name, class, level, lpmax, lp, inventory, skill, money)
}

func (p *Personnage) menu() {
	var menu int
	var retour int
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("+++++++++++++++++++++++++++++++")
	fmt.Println("A quoi voulez vous accéder:")
	fmt.Println("----- \n Afficher les informations du personnage (1)")
	fmt.Println("----- \n Accéder au contenu de l’inventaire (2)")
	fmt.Println("----- \n Voir le Marchand (3)")
	fmt.Println("----- \n Aller parler au Forgeron (4)")
	fmt.Println("----- \n S'entrainer contre un gobelin (5)")
	fmt.Println("----- \n Affronter les Quatres Grands (6)")
	fmt.Println("----- \n Quitter (7) \n-----")
	fmt.Println("Entrez le numéro de l'option:")
	fmt.Println("+++++++++++++++++++++++++++++++")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.DisplayInfo()
		fmt.Println("Retour au menu principal (1)")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			p.menu()
		}
	case 2:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.AccessInventory()
	case 3:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
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
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("------ Vous entrez dans la forge ------")
		fmt.Println("\n-- Construire un Chapeau de l'aventurier (1) --")
		fmt.Println("requiert 1 plume de corbeau et 1 cuir de sanglier")
		fmt.Println("\n-- Construire une Tunique de l'Aventurier (2) --")
		fmt.Println("requiert 2 Fourrure de Loup et 1 Peau de Troll")
		fmt.Println("\n-- Construire les bottes de l'aventurier (3) --")
		fmt.Println("requiert 1 Fourrure de Loup et 1 Cuir de Sanglier")
		fmt.Printf("\n-- Retourner au menu principal (4) --\n")
		fmt.Println("\nVotre inventaire:")
		for i := 0; i < len(p.inventory); i++ {
			if p.inventory[i] != " " {
				fmt.Println("---]", p.inventory[i], "[---")
			}
		}
		p.Forgeron()
	case 5:
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.TrainingFight()
		p.menu()
	case 6:
		p.TheFirst()
	case 7:
		fmt.Println("Fin de la transmission")
		break
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
	Equipement
}

type Equipement struct {
	Chapeau   string
	Tunique   string
	Maxbottes string
}

type Monstre struct {
	name   string
	lp     int
	lpmax  int
	attack int
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
	if len(p.inventory) == 0 {
		fmt.Println("Inventaire vide fraté")
	} else {
		fmt.Printf("\n")
		for i := 0; i < len(p.inventory); i++ {
			if p.inventory[i] != " " {
				fmt.Println("---]", p.inventory[i], "[---")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("----------------")
	fmt.Println("Prendre une Potion de vie (1)")
	fmt.Println("Prendre une Potion de poison (2)")
	fmt.Println("Utiliser le Livre de sort : Boule de Feu (3)")
	fmt.Println("Menu précédent (4)")
	fmt.Println("----------------")
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
		}
	}
}

func (p *Personnage) Forgeron() {
	var enclume int
	fmt.Scanln(&enclume)
	switch enclume {
	case 1:
		if p.Checkinv("Plume de Corbac") == true && p.Checkinv("Cuir de Sanglier") == true && p.Checkinv("Chapeau de l'aventurier") == false {
			p.AddInventory("Chapeau de l'aventurier", 5)
			p.RemoveInventory("Plume de Corbac")
			p.RemoveInventory("Cuir de Sanglier")
			fmt.Println("Vous êtes désormais en possession de Chapeau de l'aventurier")
			p.Forgeron()
		} else {
			fmt.Println("Tu te moques de moi ? Regarde ton inventaire l'ami")
			p.Forgeron()
		}
	case 2:
		if p.Checkinv("Fourrure de Loup") == true && p.Checkinv("Peau de Troll") == true {
			p.RemoveInventory("Fourrure de Loup")
			if p.Checkinv("Fourrure de Loup") == true {
				p.RemoveInventory("Fourrure de Loup")
				p.RemoveInventory("Peau de Troll")
				p.AddInventory("Tunique de l'Aventurier", 5)
				p.Forgeron()
			} else {
				p.AddInventory("Fourrure de Loup", 5)
			}
		} else {
			fmt.Println("Tu te moques de moi ? Regarde ton inventaire l'ami")
			p.Forgeron()
		}
	case 3:
		if p.Checkinv("Fourrure de Loup") == true && p.Checkinv("Cuir de Sanglier") == true {
			p.RemoveInventory("Fourrure de Loup")
			p.RemoveInventory("Cuir de Sanglier")
			p.AddInventory("Bottes de l'Aventurier", 0)
			p.Forgeron()
		} else {
			fmt.Println("Tu te moques de moi ? Regarde ton inventaire l'ami")
			p.Forgeron()
		}
	case 4:
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
	if founditem == true {
		return true
	} else {
		return false
	}
}

//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU
//VRAI COUTEAU

func (p *Personnage) RealKnife(m *Monstre) {
	for i := 0; i < len(p.inventory); i++ {
		if p.inventory[i] != " " {
			m.lp = 0
		}
	}
	fmt.Println("Vous avez tué", m.name)
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
		if p.name == "Utilisateur" {
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
		fmt.Println("Retour (3)")
		var use int
		fmt.Scanln(&use)
		switch use {
		case 1:
			p.TakePot()
		case 2:
			p.PoisonPot()
		case 3:
			p.CharTurn(m)
		}
	}
}

func (p *Personnage) TrainingFight() {
	var e1 Monstre
	var e2 Monstre
	n := rand.Intn(10)
	if n == 9 {
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
					time.Sleep(400 * time.Millisecond)
					z01.PrintRune(letter)
				}
				time.Sleep(3 * time.Second)
				p.AddInventory("Véritable couteau", 0)
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
		fmt.Println("Tour", turn)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est au joueur !")
		p.CharTurn(&e3)
		if e3.lp <= 0 {
			fmt.Println("\nVous avez vaincu", e3.name)
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
		fmt.Println("Tour", turn2)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est au joueur !")
		p.CharTurn(&e4)
		if e4.lp <= 0 {
			fmt.Println("\nVous avez vaincu", e4.name)
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
		fmt.Println("Tour", turn3)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est au joueur !")
		p.CharTurn(&e5)
		if e5.lp <= 0 {
			fmt.Println("\nVous avez vaincu", e5.name)
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
		fmt.Println("Tour", turn4)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est au joueur !")
		p.CharTurn(&e6)
		if e6.lp <= 0 {
			fmt.Println("\nVous avez vaincu", e6.name)
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
