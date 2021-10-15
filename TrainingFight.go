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
	damage := m.attack * att
	p.lp -= damage
	Slow(Yellow+m.name+ Reset, 2)
	Slow(" attaque ", 2)
	Slow(Yellow+p.name+Reset, 2)
	Slow(" et lui inflige "+Red, 2)
	fmt.Print(damage)
	Slow(" points de dégâts"+Reset,2)
	Slow(", il lui reste "+Yellow, 2)
	fmt.Print(p.lp)
	Slow(" points de vies "+Reset,2)
	Slow("sur un total de "+Yellow, 2)
	fmt.Print(p.lpmax)
	fmt.Println(""+Reset)
}

func (p *Personnage) TrainingFight(a *Equipement) {
	var e1 Monstre
	var e2 Monstre
	// var entier int 
	n := rand.Intn(4)
	fmt.Println(n)
	if n == 3 {
		Slow("Un mimic sauvage apparaît", 1)
		time.Sleep(3 * time.Second)
		var turn int
		e2.InitGoblin("Mimic", 80, 5)
		for i := 0; i <= 9999; i++ {
			turn++
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			Slow("Tour ", 1)
			fmt.Print(turn)
			time.Sleep(1 * time.Second)
			p.CharTurn(&e2,a)
			if e2.lp <= 0 {
				Slow("Vous avez vaincu le Mimic ", 1)
				Slow("\nVous gagnez ",1)
				Slow(Yellow+"10 points d'expérience"+Reset,2)
				p.exp += 10
				p.Loot()
				p.UpgradeLevel()
				time.Sleep(1 * time.Second)
				
				if !p.Checkinv("Véritable couteau") {
					p.AddInventory("Véritable couteau", 0)
					Slow("\nVous le fouillez et obtenez un objet étrange...", 5)
					time.Sleep(3 * time.Second)
				} else {
					p.Loot()
					Slow("Le Mimic est rempli de pièce, vous vous servez autant que vous le pouvez et remportez 50 pièces",1)
				}
				break
			}
			time.Sleep(1 * time.Second)
			Slow("\n\nC'est à l'ennemi !\n", 2)
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
			p.CharTurn(&e1,a)
			if e1.lp <= 0 {
				Slow("Vous avez vaincu ", 2)
				Slow(e1.name, 2)
				Slow("\nVous gagnez ",1)
				Slow(Yellow+"5 points d'expérience"+Reset,2)
				p.exp += 5
				p.Loot()
				p.UpgradeLevel()
				time.Sleep(2 * time.Second)
				break
			}
			time.Sleep(1 * time.Second)
			Slow("\n\nC'est à l'ennemi !\n", 2)
			if turn%3 == 0 || turn == 3 {
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
