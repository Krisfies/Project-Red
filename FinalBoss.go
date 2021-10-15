package main

import (
	"time"
	"os"
	"fmt"
)

func (p *Personnage) FinalBoss(e7 *Monstre, a *Equipement) {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	var hand int
	Slow("En pénétrant à l'intérieur du chateau vous ressentez une aura maléfique bien plus forte et bien plus terrifiante que celle des Qatres Grands",6)
	Slow("\nVous remarquez aussi que l'intérieur du chateau est entièrement coloré en jaune", 6)
	Slow("\nVous réussissez à distinguer une silhouette s'approchant au loin",6)
	Slow("\nElle aussi est ",5)
	Slow(Yellow+"jaune"+Reset, 5)
	Slow("\nLa silhouette semble s'adresser à vous:",6)
	Slow(Yellow+"\nJe dois te féliciter l'ami, tu as réussi à vaincre les Quatres Grands, c'est un exploit remarquable.",6)
	Slow("\nJe me suis libéré des liens qui m'entravaient, j'ai repris possession de Goldy et maintenant plus personne ne pourra se mettre en travers de mon chemin",6)
	Slow("\nEt ce, grâce à toi",6)
	Slow("\nIl est temps de te donner ta récompense"+Reset,6)
	Slow("\nL'Homme s'approche et tend sa main attendant que vous lui la serriez",6)
	Slow("\n\nLui serrer la main ?",6)
	Slow(Yellow+"\n(1)"+Reset,4)
	Slow("Oui",4)
	Slow(Yellow+"\n(2)"+Reset,4)
	Slow("Non\n",4)
	fmt.Scanln(&hand)
	switch hand {
	case 1:
		Slow("\nA l'instant où vos mains se touchent, vous sentez votre énergie vitale grandement diminuer",6)
		Slow("\nVous tombez à genoux",6)
		p.lp = 1
		Slow("\nIl vous restes "+Yellow,6)
		fmt.Print(p.lp, "/", p.lpmax)
		Slow(Reset+" PV",6)
		Slow(Yellow+"\nHo, tu es plus résistant que je ne le pensais ",6)
		Slow(p.name,6)
		Slow(Reset+"\nL'Homme semble encore plus jaune qu'il ne l'était",6)
		Slow(Yellow+"\nRelève toi, et bats toi si tu l'oses !"+Reset,6)
	case 2:
		Slow(Yellow+"\nJe vois...",6)
		Slow("\nCela m'attriste mais je vais devoir employer la manière forte"+Reset,6)
	}
	time.Sleep(1 * time.Second)
	Slow("\nL'Homme en jaune se prépare au combat",6)
	var turn4 int
	Loop:
	for i := 0; i <= 9999; i++ {
		turn4++
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		Slow("Tour ", 2)
		fmt.Println(turn4)
		p.CharTurn(e7,a)
		if e7.lp <= 0 {
			Slow("Vous avez vaincu l'Homme en jaune\n", 1)
			Slow("Il disparaît dans un flash lumineux, emportant avec lui sa couleur", 4)
			Slow("\nVous gagnez ",1)
			Slow(Yellow+"100 points d'expérience"+Reset,2)
			p.exp += 100
			p.UpgradeLevel()
			p.UpgradeLevel()
			p.UpgradeLevel()
			p.UpgradeLevel()
			time.Sleep(2 * time.Second)
			p.EndGame()
		}
		time.Sleep(1 * time.Second)
		Slow("\nC'est à l'ennemi !\n", 2)
		p.GoblinPattern(e7, 1)
		if p.lp <= 0 {
			Slow(Yellow+e7.name, 1)
			Slow(" a eu raison de vous."+Reset, 1)
			p.Dead(a)
			break Loop
		}
		time.Sleep(3 * time.Second)
	}
}
