package main

import (
	"fmt"
	"time"
)

func (p *Personnage) TakePot() {
	// fonction qui nous permet de prendre une potion de soin
	for _, letter := range p.inventory {
		if letter == "Potion de vie" {
			if p.lp <= (p.lpmax - 50) {
				p.lp += 50
				Slow("Glou glou glou, ça fait du bien", 2)
				Slow("\nVous avez désormais: "+Yellow, 2)
				fmt.Print(p.lp)
				Slow("/", 2)
				fmt.Print(p.lpmax)
				fmt.Print("\n" + Reset)
				p.RemoveInventory("Potion de vie")
				break
			} else if p.lp > (p.lpmax-50) && p.lp < p.lpmax {
				p.lp = p.lpmax
				Slow("Glou glou glou, ça fait du bien", 2)
				Slow("\nVous avez désormais: "+Yellow, 2)
				fmt.Print(p.lp)
				Slow("/", 2)
				fmt.Print(p.lpmax)
				fmt.Print("\n" + Reset)
				p.RemoveInventory("Potion de vie")
				break
			} else {
				Slow("Vous êtes au ", 2)
				Slow(Yellow+"maximum"+Reset, 2)
				Slow(", vous ne pouvez pas utiliser la ", 2)
				Slow(Yellow+"potion"+Reset, 2)
				break
			}
		} else {
			Slow("Tu n'as pas de ", 2)
			Slow(Yellow +"potion"+ Reset, 2)
		}
	}
	time.Sleep(1 * time.Second)
}

func (p *Personnage) PoisonPot(a *Equipement) {
	// fonction qui crée la potion poison et explique ce qu'elle fait sur un personnage
	for _, letter := range p.inventory {
		if letter == "Potion de poison" {
			Slow("Vous buvez la ", 2)
			Slow(Green+"Potion de poison"+Reset, 2)
			Slow(", ouch !", 1)
			p.RemoveInventory("Potion de poison")
			time.Sleep(1 * time.Second)
			for i := 0; i <= 3; i++ {
				p.lp -= 10
				Slow("Vos points de vie ", 2)
				Slow(Red +"diminuent: "+ Reset, 2)
				fmt.Print(p.lp)
				Slow("/", 1)
				fmt.Print(p.lpmax)
				p.Dead(a)
				time.Sleep(1 * time.Second)
			}
			break
		} else {
			Slow("\nTu n'as pas de ", 2)
			Slow(Yellow +"potion\n"+ Reset, 2)
		}
	}
	time.Sleep(1 * time.Second)
}
