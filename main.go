package main

import (
	"fmt"
	"time"
)

func main() {
	// fonction qui execute nos sous fonctions et rentre les valeur ainsi que le menu principal
	var p1 Personnage
	p1.Init("Bijoux", "elfe", 1, 100, 42, []string{"Potion de vie", "Potion de vie", "Potion de vie"}, []string{"Coup de poing"})
	var menu int
	fmt.Println("+++++++++++++++++++++++++++++++")
	fmt.Println("A quoi voulez vous accéder:")
	fmt.Println("----- \n Afficher les informations du personnage (1)")
	fmt.Println("----- \n Accéder au contenu de l’inventaire (2)")
	fmt.Println("----- \n Voir le Marchand (3)")
	fmt.Println("----- \n Quitter (4) \n-----")
	fmt.Println("Entrez le numéro de l'option:")
	fmt.Println("+++++++++++++++++++++++++++++++")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		p1.DisplayInfo()
	case 2:
		p1.AccessInventory()
	case 3:
		p1.Marchand()
	case 4:
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
}

func (p *Personnage) Init(name string, class string, level int, lpmax int, lp int, inventory []string, skill []string) {
	// initialisation de notre personnage
	p.name = name
	p.class = class
	p.level = level
	p.lpmax = lpmax
	p.lp = lp
	p.inventory = inventory
	p.skill = skill
}

func (p *Personnage) AccessInventory() {
	// fonction qui nous permet d'acceder a notre inventaire
	if len(p.inventory) == 0 {
		fmt.Println("inventaire vide fraté")
	} else {
		fmt.Println(p.inventory)
	}
	retour()
}

func retour() {
	// fonction qui nous permet de retourner au menu précédent
	var rep int
	fmt.Println("tapez 1 pour retourner zo menu précédent")
	fmt.Scanln(&rep)
	if rep == 1 {
		main()
	}
}

func (p *Personnage) spellbook(item string) {
	// fonction qui nous permet d'ajouter ou repertorier les sorts (spell)
	h := &p.skill
	for _, letter := range p.skill {
		if letter == ("Boule de feu") {
			fmt.Println("dsl t'a déja les boules")
		} else {
			*h = append(*h, item)
		}
	}
	retour()
}

func (p *Personnage) Marchand() {
	// fonction affichant le menu du marchand , et les ajoute a notre inventaire
	var menum int
	fmt.Println("-----------------Marchand-------------------")
	fmt.Println("Tapez 1 pour obtenir une Potion de vie ;)")
	fmt.Println("Tapez 2 pour obtenir une Potion de poison ;(")
	fmt.Println("Tapez 3 pour tenter d'obtenir boule de feu ")
	fmt.Println("Tapez 4 pour retourner au menu précédent ")
	fmt.Println("--------------------------------------------")
	fmt.Scanln(&menum)
	switch menum {
	case 1:
		p.AddInventory("Potion de vie")
		p.AccessInventory()
	case 2:
		p.AddInventory("Potion de poison")
		p.AccessInventory()
	case 3:
		p.spellbook("boule de feu")
	case 4:
		retour()
	}
}

func (p *Personnage) DisplayInfo() {
	// fonction nous permettant de voir les informations de notre personnage
	fmt.Println("-----------")
	fmt.Println("Nom:", p.name)
	fmt.Println("Classe:", p.class)
	fmt.Println("Niveau:", p.level)
	fmt.Println("Vie maximum:", p.lpmax)
	fmt.Println("Vie actuelle:", p.lp)
	fmt.Println("Contenu de l'inventaire:", p.inventory)
	fmt.Println("skill :", p.skill)
	fmt.Println("-----------")
	retour()

}

func (p *Personnage) AddInventory(item string) {
	p.inventory = append(p.inventory, item)
}

func (p *Personnage) TakePot() {
	// fonction qui nous permet de prendre une potion de soin
	for _, letter := range p.inventory {
		if letter == "Potion de vie" {
			if p.lp <= (p.lpmax - 50) {
				p.lp += 50
				p.inventory[len(p.inventory)-1] = ""
				break
			} else if p.lp > (p.lpmax-50) && p.lp < p.lpmax {
				p.lp = p.lpmax
				p.inventory[len(p.inventory)-1] = ""
				break
			} else {
				fmt.Println("Vous êtes full")
				break
			}
		}
	}
}

func (p *Personnage) Dead() {
	// fonction qui verifie si le personnage est mort et le ressussite a la moitié de ses pv
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
			time.Sleep(1 * time.Second)
			fmt.Println(p.lp, "/", p.lpmax)
			p.lp -= 10
			fmt.Println(p.lp, "/", p.lpmax)
			time.Sleep(1 * time.Second)
			p.lp -= 10
			fmt.Println(p.lp, "/", p.lpmax)
			time.Sleep(1 * time.Second)
			p.lp -= 10
			fmt.Println(p.lp, "/", p.lpmax)
		}
	}
}
