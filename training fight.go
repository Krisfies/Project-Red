package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

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
