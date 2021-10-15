package main

import "fmt"

func (p *Personnage) UpgradeLevel() {
	if p.exp >= p.expmax {
		Slow("\nBravo, vous passez au niveau "+Yellow,1)
		p.level += 1
		fmt.Print(p.level)
		p.exp -= p.expmax
		Slow(Reset+"\nVos PV passent de "+Yellow,1)
		fmt.Print(p.lp, "/", p.lpmax)
		p.lp += 15
		p.lpmax += 15
		Slow(Reset+" à "+Yellow,1)
		fmt.Print(p.lp, "/", p.lpmax)
		p.damage += 2
		Slow(Reset+"\nVos dégâts augmentent de ",1)
		Slow(Yellow+"2\n"+Reset,1)
		p.expmax += 5
	}
}