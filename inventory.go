package main

import (
	"fmt"
	"os"
)

func (p *Personnage) AccessInventory(e3, e4, e5, e6 *Monstre, a *Equipement) {
	// fonction qui nous permet d'acceder a notre inventaire
	var LifePotion bool = false
	var PoisonPotion bool = false
	var FireBall bool = false
	var Hat bool = false
	var Tunique bool = false
	var Boots bool = false
	var Weapon bool = false
	if len(p.inventory) == 0 {
		Slow("\nLe sac est ", 1)
		Slow(Yellow+"vide"+Reset, 1)
	} else {
		fmt.Print("\n")
		for i := 0; i < len(p.inventory); i++ {
			if p.inventory[i] != " " {
				Slow("---] ", 2)
				fmt.Print(""+Yellow, p.inventory[i], Reset+"")
				Slow(" [---", 2)
				fmt.Print("\n")
				if p.inventory[i] == "Potion de vie" {
					LifePotion = true
				}
				if p.inventory[i] == "Potion de poison" {
					PoisonPotion = true
				}
				if p.inventory[i] == "Livre de sort: Boule de Feu" {
					FireBall = true
				}
				if p.inventory[i] == "Couronne en or" || p.inventory[i] == "Chapeau de paille" || p.inventory[i] == "Chapeau d'érudit" || p.inventory[i] == "Casque de mineur" || p.inventory[i] ==  "Bonnet de petite taille" || p.inventory[i] == "Chapeau de l'aventurier" {
					Hat = true
				}
				if p.inventory[i] == "Salopette rapiécée" || p.inventory[i] == "Cape en fourrure, ornée de cristaux" || p.inventory[i] == "Vieux manteau" || p.inventory[i] == "Robe de sage" || p.inventory[i] == "Veste abîmée" ||p.inventory[i] == "Tunique de l'Aventurier" {
					Tunique = true
				}
				if p.inventory[i] == "Sabot renforcé" || p.inventory[i] == "Bottes en cuir" || p.inventory[i] == "Vieille claquette" || p.inventory[i] == "Chaussure pointue" || p.inventory[i] == "Sabot en boît" || p.inventory[i] == "Bottes de l'aventurier" {
					Boots = true
				}
				if p.inventory[i] == "Véritable Couteau" {
					Weapon = true
				}
			}
		}
		fmt.Print("\n")
	}
	Slow("----------------\n", 3)
	if LifePotion {
		Slow(Yellow+"(1) "+Reset, 3)
		Slow("Prendre une ", 3)
		Slow(Yellow+"Potion de vie\n"+Reset, 3)
	}
	if PoisonPotion {
		Slow(Yellow+"(2) "+Reset, 3)
		Slow("Prendre une ", 3)
		Slow(Yellow+"Potion de poison\n"+Reset, 3)
	}
	if FireBall {
		Slow(Yellow+"(3) "+Reset, 3)
		Slow("Utiliser le Livre de sort: ", 3)
		Slow(Yellow+"Boule de Feu\n"+Reset, 3)
	}
	if Hat {
		Slow(Yellow+"(4) "+Reset, 3)
		Slow("Mettre un ", 3)
		Slow(Yellow+"Chapeau\n"+Reset, 3)
	}
	if Tunique {
		Slow(Yellow+"(5) "+Reset, 3)
		Slow("Mettre une ", 3)
		Slow(Yellow+"Tunique\n"+Reset, 3)
	}
	if Boots {
		Slow(Yellow+"(6) "+Reset, 3)
		Slow("Mettre des ", 3)
		Slow(Yellow+"Bottes\n"+Reset, 3)
	}
	if Weapon {
		Slow(Yellow+"(7)"+Reset,3)
		Slow("Equiper une",3)
		Slow(Yellow+"Arme"+Reset,3)
	}
	Slow(Yellow+"(0) "+Reset, 3)
	Slow("Arrêter de regarder dans le ", 3)
	Slow(Yellow+"Sac\n"+Reset, 3)
	Slow("----------------\n", 3)
	p.UseInventory(e3, e4, e5, e6, a)
}

func (p *Personnage) SuperAccessInventory(e3, e4, e5, e6 *Monstre, a *Equipement) {
	// fonction qui nous permet d'acceder a notre inventaire
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	var LifePotion bool = false
	var PoisonPotion bool = false
	var FireBall bool = false
	var Hat bool = false
	var Tunique bool = false
	var Boots bool = false
	var Weapon bool = false
	if len(p.inventory) == 0 {
		Slow("\nLe sac est ", 1)
		Slow(Yellow+"vide"+Reset, 1)
	} else {
		fmt.Print("\n")
		for i := 0; i < len(p.inventory); i++ {
			if p.inventory[i] != " " {
				Slow("---] ", 2)
				fmt.Print(""+Yellow, p.inventory[i], Reset+"")
				Slow(" [---", 2)
				fmt.Print("\n")
				if p.inventory[i] == "Potion de vie" {
					LifePotion = true
				}
				if p.inventory[i] == "Potion de poison" {
					PoisonPotion = true
				}
				if p.inventory[i] == "Livre de sort: Boule de Feu" {
					FireBall = true
				}
				if p.inventory[i] == "Couronne en or" || p.inventory[i] == "Chapeau de paille" || p.inventory[i] == "Chapeau d'érudit" || p.inventory[i] == "Casque de mineur" || p.inventory[i] ==  "Bonnet de petite taille"{
					Hat = true
				}
				if p.inventory[i] == "Salopette rapiécée" || p.inventory[i] == "Cape en fourrure, ornée de cristaux" || p.inventory[i] == "Vieux manteau" || p.inventory[i] == "Robe de sage" || p.inventory[i] == "Veste abîmée" {
					Tunique = true
				}
				if p.inventory[i] == "Sabot renforcé" || p.inventory[i] == "Bottes en cuir" || p.inventory[i] == "Vieille claquette" || p.inventory[i] == "Chaussure pointue" || p.inventory[i] == "Sabot en boît" {
					Boots = true
				}
				if p.inventory[i] == "Véritable Couteau" {
					Weapon = true
				}
			}
		}
		fmt.Println("\n----------------")
	}
	if LifePotion {
		fmt.Println(Yellow+"(1)"+Reset, "Prendre une", Yellow+"Potion de vie"+Reset)
	}
	if PoisonPotion {
		fmt.Println(Yellow+"(2)"+Reset, "Prendre une", Yellow+"Potion de poison"+Reset)
	}
	if FireBall {
		fmt.Println(Yellow+"(3)"+Reset, "Utiliser le Livre de sort :", Yellow+"Boule de Feu"+Reset)
	}
	if Hat {
		fmt.Println(Yellow+"(4)"+Reset, "Mettre un", Yellow+"Chapeau"+Reset)
	}
	if Tunique {
		fmt.Println(Yellow+"(5)"+Reset, "Mettre une", Yellow+"Tunique"+Reset)
	}
	if Boots {
		fmt.Println(Yellow+"(6)"+Reset, "Mettre des", Yellow+"Bottes"+Reset)
	}
	if Weapon {
		fmt.Println(Yellow+"(7)"+Reset, "Equiper une", Yellow+"Arme"+Reset)
	}
	fmt.Println(Yellow+"(0)"+Reset, "Arrêter de regarder dans le", Yellow+"Sac"+Reset, "\n----------------")
	p.UseInventory(e3, e4, e5, e6, a)
}

func (p *Personnage) UseInventory(e3 , e4 , e5 , e6 *Monstre, a *Equipement) {
	// fonction qui nous permet d'interagir avec les éléments de l'inventaire
	var use int
	fmt.Scanln(&use)
	switch use {
	case 1:
		p.TakePot()
		p.SuperAccessInventory(e3, e4, e5, e6, a)
	case 2:
		p.PoisonPot(a)
		p.SuperAccessInventory(e3, e4, e5, e6, a)
	case 3:
		p.Spellbook(e3, e4, e5, e6, a, "Boule de Feu")
		p.SuperAccessInventory(e3, e4, e5, e6, a)
	case 4:
		if p.Checkinv("Chapeau de l'aventurier") {
			if a.ExchangeEquip("chapeau", p) {
				p.lp +=15
				p.lpmax+= 15
				Slow("Vous avez désormais +15 points de vies maximum avec le couvre chef !\n", 2)
			}
			p.RemoveInventory("Chapeau de l'Aventurier")
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		} else if !p.Checkinv("Chapeau de L'aventurier") {
			Slow("Tu n'as pas cet obkjet dans ton inventaire !", 2)
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		}
	case 5:
		if p.Checkinv("Tunique de l'Aventurier") {
			if a.ExchangeEquip("tunique", p) {
				p.lp +=20
				p.lpmax+= 20
				Slow("Vous avez désormais +20 points de vies maximum grâce a la tunique !\n", 2)
			}
			p.RemoveInventory("Tunique de l'Aventurier")
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		} else if !p.Checkinv("Tunique de L'aventurier"){
			Slow("Tu n'as pas cet obkjet dans ton inventaire !", 2)
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		}
	case 6:
		if p.Checkinv("Bottes de l'Aventurier") {
			if a.ExchangeEquip("bottes",p) {
				p.lp +=10
				p.lpmax+= 10
				Slow("Vous avez désormais +10 points de vies maximum grâce aux bottes !\n", 2)
			}
			p.RemoveInventory("Bottes de l'Aventurier")
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		} else if !p.Checkinv("Bottes de L'aventurier"){
			Slow("Tu n'as pas cet obkjet dans ton inventaire !", 2)
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		}
	case 7:
		if p.Checkinv("Véritable Couteau") {
			p.RemoveInventory("Véritable Couteau")
			a.Arme = "Véritable Couteau"
			p.damage += 15
			Slow("Vous équipez le Véritable Couteau\n",2)
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		} else  if !p.Checkinv("Véritable Couteau") {
			Slow("Tu n'as pas cet obkjet dans ton inventaire !",2)
			p.SuperAccessInventory(e3, e4, e5, e6, a)
		}
	case 0:
		p.menu(e3, e4, e5, e6, a)
	}
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
