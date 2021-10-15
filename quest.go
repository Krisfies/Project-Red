package main

import (
	"os"
	"fmt"
	"time"
)

func (p *Personnage) QuestMan(e3, e4, e5, e6 *Monstre, a *Equipement) {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	var choice int
	if !p.Checkinv("Objet Suspicieux") {
		if !p.Checkinv("Véritable Couteau") || a.Arme == "Véritable Couteau" {
				Slow(Yellow+"Bienvenue sur la Place Principale l'ami.\n"+Reset, 1)
				Slow("(1) Lui demander son nom ?\n", 1)
				Slow("(2) Lui demander ce qu'on fait la ?\n", 1)
				fmt.Scanln(&choice)
				switch choice {
				case 1:
					Slow(Yellow+"Mon nom importe peu l'ami, je ne suis qu'un personnage de plus dans cette histoire, mon rôle est malheureusement insignifiant.", 1)
					Slow("\nMais si tu pouvais me rendre un service, peut être pourrais-je me libérer de mes chaines ?\n"+Reset,1)
				case 2:
					Slow(Yellow+"Il y a quelques temps, Goldy le pays où tout est jaune, s'est fait envahir par quatres seigneurs malveillants", 1)
					Slow("\nJe te confie la mission de les éliminer, je te récompenserais en conséquence", 1)
					Slow("\nMais avant cela tu dois te rendre dans le donjon afin d'y trouver l'arme qui te donnera la force de les battre.", 1)
					Slow("\nPourquoi n'y vais-je pas moi ? m'attaquer à eux pourrait supprimer mon existence du programme haha !"+Reset,1)
					Slow("\nL'homme est pris d'un fous rire. Puis reprend quelques instants plus tard",1)
					Slow(Yellow+"\nBonne chance l'ami !\n"+Reset, 1)
				}
		} else {
			Slow(Yellow+"Bravo l'ami ! Tu as obtenus l'artéfact.", 1)
			Slow("\nJe te conseille maintenant d'aller t'équiper en conséquence et de te lancer à la recherche des Quatres Grands"+Reset, 1)
		}
	}
	time.Sleep(2 * time.Second)
}
