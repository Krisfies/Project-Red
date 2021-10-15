package main

import (
	"os"
	"time"
	"fmt"
)

func (p *Personnage) TheFirst(e3, e4, e5, e6 *Monstre, a *Equipement) {
	var enter int
	if !p.Checkinv("Objet Suspicieux") {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		Slow("Vous vous dirigez vers l'enceinte d'un chateau, les pierres qui le forment sont sombres et un tapis rouge flamboyant marque l'entrée de la bâtisse\n",1)
		Slow("Vous vous apprêtez à rentrer lorsque vous ressentez une aura terrifiante provenant tout droit du chateau\n\n",1)
		Slow(Yellow+"(1) "+Reset,2)
		Slow("Continuer\n",2)
		Slow(Yellow+"(2) "+Reset,2)
		Slow("Rebrousser chemin\n",2)
		fmt.Scanln(&enter)
	} else {
		enter = 1
	}
	switch enter {
	case 1:
		Slow("Vous prenez votre courage à deux mains et vous ",1)
		Slow(Yellow+"pénétrez dans le chateau\n"+Reset,1)
		if e3.lp <= 0 {
			p.TheSecond(e3, e4, e5, e6, a)
		} else {
			Slow("Le Grand Python se tiens devant vous\n", 1)
			time.Sleep(1 * time.Second)
			var turn int
			for i := 0; i <= 9999; i++ {
				turn++
				os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
				Slow("Tour ", 1)
				fmt.Println(turn)
				p.CharTurn(e3, a)
				if e3.lp <= 0 {
					Slow("\nVous avez vaincu le Grand ", 1)
					Slow(e3.name, 1)
					Slow(" !", 1)
					Slow("\nVous gagnez ",1)
					Slow(Yellow+"15 points d'expérience"+Reset,2)
					p.exp += 15
					p.UpgradeLevel()
					time.Sleep(2 * time.Second)
					p.TheSecond(e3, e4, e5, e6, a)
				}
				time.Sleep(1 * time.Second)
				Slow("\nC'est à l'ennemi !\n", 2)
				p.GoblinPattern(e3, 1)
				if p.lp <= 0 {
					Slow(e3.name, 1)
					Slow(" vous a battu", 1)
					p.Dead(a)
					break
				}
				time.Sleep(3 * time.Second)
			}
		}
	case 2:
		p.menu(e3, e4, e5, e6, a)
	}
}

func (p *Personnage) TheSecond(e3, e4, e5, e6 *Monstre, a *Equipement) {
	if e4.lp <= 0 {
		p.TheThird(e3, e4, e5, e6, a)
	} else {
		Slow("Le Grand Java se tiens devant vous\n", 1)
		time.Sleep(1 * time.Second)
		var turn2 int
		for i := 0; i <= 9999; i++ {
			turn2++

			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			Slow("Tour ", 2)
			fmt.Print(turn2)
			p.CharTurn(e4,a)
			if e4.lp <= 0 {
				Slow("\nVous avez vaincu le Grand ", 1)
				Slow(e4.name, 1)
				Slow(" !", 1)
				Slow("\nVous gagnez ",1)
				Slow(Yellow+"20 points d'expérience"+Reset,2)
				p.exp += 20
				p.UpgradeLevel()
				time.Sleep(2 * time.Second)
				p.TheThird(e3, e4, e5, e6, a)
			}
			time.Sleep(1 * time.Second)
			Slow("\nC'est à l'ennemi !\n", 2)
			p.GoblinPattern(e4, 1)
			if p.lp <= 0 {
				Slow(e4.name, 1)
				Slow(" vous a battu", 1)
				p.Dead(a)
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}

func (p *Personnage) TheThird(e3, e4, e5, e6 *Monstre, a *Equipement) {
	if e5.lp <= 0 {
		p.TheFourth(e3, e4, e5, e6, a)
	} else {
		Slow("Le Grand C++ se tiens devant vous\n", 1)
		time.Sleep(1 * time.Second)
		var turn3 int
		for i := 0; i <= 9999; i++ {
			turn3++

			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			Slow("Tour ", 2)
			fmt.Print(turn3)
			p.CharTurn(e5,a)
			if e5.lp <= 0 {
				Slow("\nVous avez vaincu le Grand ", 1)
				Slow(e5.name, 1)
				Slow(" !", 1)
				Slow("\nVous gagnez ",1)
				Slow(Yellow+"25 points d'expérience"+Reset,2)
				p.exp += 25
				p.UpgradeLevel()
				time.Sleep(2 * time.Second)
				p.TheFourth(e3, e4, e5, e6, a)
			}
			time.Sleep(1 * time.Second)
			Slow("\nC'est à l'ennemi !\n", 2)
			p.GoblinPattern(e5, 1)
			if p.lp <= 0 {
				Slow(e5.name, 1)
				Slow(" vous a battu", 1)
				p.Dead(a)
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}

func (p *Personnage) TheFourth(e3, e4, e5, e6 *Monstre, a *Equipement) {
	if e6.lp <= 0 {
		Slow("Vous avez déjà vaincu les Quatres Grands, il n'y a plus personne ici",1)
	} else {
		Slow("Le plus Grand des Grand, Golang, se tiens devant vous\n", 1)
		time.Sleep(1 * time.Second)
		var turn4 int
		for i := 0; i <= 9999; i++ {
			turn4++

			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			Slow("Tour ", 2)
			fmt.Println(turn4)
			p.CharTurn(e6,a)
			if e6.lp <= 0 {
				Slow("Vous avez vaincu Golang, le plus Grand des Grands.\n", 1)
				Slow("Dans son dernier souffle il lâche quelque chose, vous le ramassez.", 4)
				p.AddInventory("Objet suspicieux", 0)
				Slow("\nVous gagnez ",1)
				Slow(Yellow+"30 points d'expérience"+Reset,2)
				p.exp += 30
				p.UpgradeLevel()
				time.Sleep(2 * time.Second)
				break
			}
			time.Sleep(1 * time.Second)
			Slow("\nC'est à l'ennemi !\n", 2)
			p.GoblinPattern(e6, 1)
			if p.lp <= 0 {
				Slow(e6.name, 1)
				Slow(" vous a battu", 1)
				p.Dead(a)
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}
