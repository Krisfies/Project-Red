package main

import (
	"fmt"
	"time"
	"os"
)

func (p *Personnage) Dead(a *Equipement) {
	// fonction qui verifie si le personnage est mort et le ressucite a la moitié de ses pv
	if p.lp <= 0 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		Slow("Bravo, vous êtes ", 1)
		Slow(Red+"mort."+Reset, 1)
		time.Sleep(1 * time.Second)
		Slow("\nMais ne paniquez pas, vous allez être ", 1)
		Slow(Yellow+"ressucité"+Reset, 1)
		time.Sleep(1 * time.Second)
		Slow("\nManoeuvre de ", 1)
		Slow(Yellow+"réanimation "+Reset, 1)
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
