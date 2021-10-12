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
				fmt.Print("\n")
				Slow("Vous avez désormais: ", 2)
				fmt.Print(Yellow + "")
				fmt.Print(p.lp)
				Slow("/", 2)
				fmt.Print(p.lpmax)
				fmt.Print("\n" + Reset)
				p.RemoveInventory("Potion de vie")
				break
			} else if p.lp > (p.lpmax-50) && p.lp < p.lpmax {
				p.lp = p.lpmax
				Slow("Glou glou glou, ça fait du bien", 2)
				fmt.Print("\n")
				Slow("Vous avez désormais: ", 2)
				fmt.Print(Yellow + "")
				fmt.Print(p.lp)
				Slow("/", 2)
				fmt.Print(p.lpmax)
				fmt.Print("\n" + Reset)
				p.RemoveInventory("Potion de vie")
				break
			} else {
				Slow("Vous êtes au ", 2)
				fmt.Print(Yellow + "")
				Slow("maximum", 2)
				fmt.Print("" + Reset)
				Slow(", vous ne pouvez pas utiliser la ", 2)
				fmt.Print(Yellow + "")
				Slow("potion", 2)
				fmt.Print("" + Reset)
				break
			}
		} else {
			Slow("Tu n'as pas de ", 2)
			fmt.Print(Yellow + "")
			Slow("potion", 2)
			fmt.Print("" + Reset)
		}
	}
	time.Sleep(1 * time.Second)
}

func (p *Personnage) PoisonPot() {
	// fonction qui crée la potion poison et explique ce qu'elle fait sur un personnage
	for _, letter := range p.inventory {
		if letter == "Potion de poison" {
			Slow("Vous buvez la ", 2)
			fmt.Print(Green + "")
			Slow("Potion de poison", 2)
			fmt.Print("" + Reset)
			Slow(", ouch !", 1)
			p.RemoveInventory("Potion de poison")
			time.Sleep(1 * time.Second)
			p.lp -= 10
			Slow("Vos points de vie ", 2)
			fmt.Print(Red + "")
			Slow("diminuent: ", 2)
			fmt.Print("" + Reset)
			fmt.Print(p.lp)
			Slow("/", 1)
			fmt.Print(p.lpmax)
			time.Sleep(1 * time.Second)
			Slow("Vos points de vie ", 2)
			fmt.Print(Red + "")
			Slow("diminuent: ", 2)
			fmt.Print("" + Reset)
			p.lp -= 10
			fmt.Print(p.lp)
			Slow("/", 1)
			fmt.Print(p.lpmax)
			fmt.Println(p.lp, "/", p.lpmax, "PV")
			time.Sleep(1 * time.Second)
			Slow("Vos points de vie ", 2)
			fmt.Print(Red + "")
			Slow("diminuent: ", 2)
			fmt.Print("" + Reset)
			p.lp -= 10
			fmt.Print(p.lp)
			Slow("/", 1)
			fmt.Print(p.lpmax)
			fmt.Println(p.lp, "/", p.lpmax, "PV")
			break
		} else {
			Slow("\nTu n'as pas de ", 2)
			fmt.Print(Yellow + "")
			Slow("potion", 2)
			fmt.Print("\n" + Reset)
			break
		}
	}
	time.Sleep(1 * time.Second)
}
