package main

import (
	"fmt"
	"time"
)

func main() {
	var p1 Personnage
	p1.Init("Matéo", "elfe", 1, 100, 30, []string{"Potion de vie", "Potion de vie", "Potion de vie"})

	var menu int
	fmt.Println("-----------")
	fmt.Println("A quoi voulez vous accéder? \n Option 1: Afficher les informations du personnage \n Option 2: Accéder au contenu de l’inventaire \n Option 3: Voir le Marchand\n Option 4: Quitter ")
	fmt.Println("Entrez le numéro de l'option")
	fmt.Println("-----------")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		p1.DisplayInfo()
		p1.Dead()
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
	name      string
	class     string
	level     int
	lpmax     int
	lp        int
	inventory []string
}

func (p *Personnage) Init(name string, class string, level int, lpmax int, lp int, inventory []string) {
	p.name = name
	p.class = class
	p.level = level
	p.lpmax = lpmax
	p.lp = lp
	p.inventory = inventory
}

func (p *Personnage) AccessInventory() {
	var rep int
	if len(p.inventory) == 0 {
		fmt.Println("inventaire vide fraté")
	}
	for i := 0; i < len(p.inventory); i++ {
		fmt.Println("---]", p.inventory[i], "[---")
	}
	fmt.Println("tapez 1 pour retourner zo menu précédent")
	fmt.Scanln(&rep)
	if rep == 1 {
		main()
	}

}
func (p *Personnage) Marchand() {
	var menum int
	fmt.Println("-----------------Marchand-------------------")
	fmt.Println("tapez 1 pour obtenir une Potion de vie ;)")
	fmt.Println("tapez 2 pour obtenir une Potion de poison ;(")
	fmt.Println(" tapez 3 pour retourner au menu précédent ")
	fmt.Println("____________________________________________")
	fmt.Scanln(&menum)
	switch menum {
	case 1:
		p.inventory = append(p.inventory, "Potion de vie")
		p.AccessInventory()
	case 2:
		p.inventory = append(p.inventory, "Potion de poison")
		p.AccessInventory()
	case 3:
		main()
	}
}

func (p Personnage) DisplayInfo() {
	fmt.Println("-----------")
	fmt.Println("Nom:", p.name)
	fmt.Println("Classe:", p.class)
	fmt.Println("Niveau", p.level)
	fmt.Println("Vie maximum", p.lpmax)
	fmt.Println("Vie actuelle", p.lp)
	fmt.Println("Inventaire", p.inventory)
	fmt.Println("-----------")
}

func (p *Personnage) TakePot() {
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
	if p.lp == 0 {
		fmt.Printf("Bravo, vous êtes mort. \n")
		time.Sleep(2 * time.Second)
		fmt.Printf("Mais ne paniquez pas, vous allez être ressuciter \n")
		time.Sleep(2 * time.Second)
		fmt.Println("Manoeuvre de réanimation en cours. . .")
		time.Sleep(3 * time.Second)
		p.lp = p.lpmax / 2
		p.DisplayInfo()
	}
}
