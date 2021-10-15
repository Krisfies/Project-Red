package main

import (

)

func (a *Equipement) ExchangeEquip(item string, p *Personnage) bool {
	if item == "chapeau" {
		if a.Chapeau!= "Chapeau de l'aventurier" {
			p.AddInventory(a.Chapeau, 0)
			p.RemoveInventory("Chapeau de l'aventurier")
			a.Chapeau = "Chapeau de l'aventurier"
			return true
		} else {
			return false
		}
	} else if item == "tunique" {
		if a.Tunique!= "Tunique de l'aventurier" {
			p.AddInventory(a.Tunique, 0)
			p.RemoveInventory("Tunique de l'aventurier")
			a.Tunique = "Tunique de l'aventurier"
			return true
		} else {
			return false
		}
	} else if item == "bottes" {
		if a.Bottes!= "Bottes de l'aventurier" {
			p.AddInventory(a.Bottes, 0)
			p.RemoveInventory("Bottes de l'aventurier")
			a.Bottes = "Bottes de l'aventurier"
			return true
		} else {
			return false
		}
	}
	return false
}
