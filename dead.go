package main

import (
	"fmt"
	"time"
)

func (p *Personnage) Dead(a *Equipement) {
	// fonction qui verifie si le personnage est mort et le ressucite a la moitié de ses pv
	if p.lp == 0 {
		Slow("Bravo, vous êtes ", 1)
		fmt.Print(Red + "")
		Slow("mort.", 1)
		fmt.Print("" + Reset)
		fmt.Print("\n")
		time.Sleep(1 * time.Second)
		Slow("Mais ne paniquez pas, vous allez être ", 1)
		fmt.Print(Yellow + "")
		Slow("ressucité", 1)
		fmt.Print("" + Reset)
		fmt.Print("\n")
		time.Sleep(1 * time.Second)
		Slow("Manoeuvre de ", 1)
		fmt.Print(Yellow + "")
		Slow("réanimation ", 1)
		fmt.Print("" + Reset)
		Slow("en cours", 1)
		time.Sleep(1 * time.Second)
		Slow(". ", 1)
		time.Sleep(1 * time.Second)
		Slow(". ", 1)
		time.Sleep(1 * time.Second)
		Slow(".", 1)
		fmt.Print("\n")
		time.Sleep(1 * time.Second)
		p.lp = p.lpmax / 2
		p.DisplayInfo(a)
	}
}
