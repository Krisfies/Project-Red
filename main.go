package main

import "fmt"

func main() {
	var p1 Personnage
	p1.Init("Matéo", "elfe", 1, 100, 42, []string{"Potion de vie", "Potion de vie", "Potion de vie"})

	var menu int
	fmt.Println("-----------")
	fmt.Println("A quoi voulez vous accéder? \n Option 1: Afficher les informations du personnage \n Option 2: Accéder au contenu de l’inventaire \n Option 3: Quitter ")
	fmt.Println("Entrez le numéro de l'option")
	fmt.Println("-----------")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		p1.DisplayInfo()
	case 2:
		fmt.Println("Pas encore programmé bg")
	case 3:
		fmt.Println("Non, tu restes ici")
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
