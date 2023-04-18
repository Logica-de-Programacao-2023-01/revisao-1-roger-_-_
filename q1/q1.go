package q1

import "fmt"

// Uma loja de roupas precisa de uma função para verificar
//se um cliente é elegível para um desconto com base
//em seu histórico de compras.

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	if currentPurchase == 0 {
		return 0, fmt.Errorf("valor da compra inválido")
	}

	var soma, media, resultado float64

	for i := 0; i < len(purchaseHistory); i++ {
		soma += purchaseHistory[i]
	}

	media = soma / float64(len(purchaseHistory))

	if len(purchaseHistory) == 0 {
		resultado = currentPurchase * 0.1
	}
	if soma > 1000 {
		resultado = currentPurchase * 0.1
	} else if soma <= 1000 {
		resultado = currentPurchase * 0.05
	} else if soma <= 500 {
		resultado = currentPurchase * 0.02
	}
	if media > 1000 {
		resultado = currentPurchase * 0.2
	}

	return resultado, nil
}
