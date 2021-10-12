package main

import (
	"fmt"
	"time"
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

func (p *Personnage) Checkinv(item string) bool {
	var founditem bool = false
	for _, letter := range p.inventory {
		if letter == item {
			founditem = true
		}
	}
	return founditem
}
