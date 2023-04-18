package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	if currentPurchase == 0 {
		fmt.Print("valor da compra inválido")
	}
	var soma, media, resultado float64
	for i := 0; i < len(purchaseHistory); i++ {
		soma += purchaseHistory[i]
	}
	media = soma / float64(len(purchaseHistory))
	if len(purchaseHistory) == 0 {
		resultado = currentPurchase * 0.1
	} else {
		if soma > 1000.00 {
			resultado = currentPurchase * 0.1
		}
		if soma <= 1000.00 {
			resultado = currentPurchase * 0.05
		}
		if soma <= 500.00 {
			resultado = currentPurchase * 0.2
		}
		if media > 1000.00 {
			resultado = currentPurchase * 0.2
		}
	}
	return resultado, nil
}

