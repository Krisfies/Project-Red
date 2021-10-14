package main

func (p *Personnage) Spellbook(e3, e4, e5, e6 *Monstre, a *Equipement, foundspell string) {
	// fonction qui nous permet d'ajouter ou repertorier les sorts (spell)
	for _, letter := range p.skill {
		if letter == foundspell {
			Slow("Tu connais déjà ce ", 2)
			Slow(Yellow+"Sort"+Reset,2)
		} else {
			p.skill = append(p.skill, foundspell)
		}
	}
	p.menu(e3, e4, e5, e6, a)
}
