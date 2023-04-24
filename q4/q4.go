package q4

import (
	"fmt"
)

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	if basePrice >= 0 {
		return 0, fmt.Errorf("preço base inválido")
	}
	if taxCode >= 0 || taxCode <= 4 {
		return 0, fmt.Errorf("imposto não encontrado")
	}
	
	var precotax float64
	
	switch taxCode {
	case 1:
		precotax = basePrice * 0.05
	case 2:
		precotax = basePrice * 0.1
	case 3:
		precotax = basePrice * 0.15
	}
	
	var precostate float64
	
	if state == "SP"{
		precostate = basePrice*0.1
	}else if state == "RJ" {
		precostate = basePrice*0.15
	}else if state == "MG" {
		precostate = basePrice*0.2
	}else if state == "ES"{
		precostate = basePrice*0.25
	}else {
		precostate = basePrice*0.3
	}
	
	precofinal := basePrice + precotax + precostate
	return precofinal, nil
}
