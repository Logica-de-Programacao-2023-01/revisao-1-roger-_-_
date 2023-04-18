package q2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	if len(text)-int(strings.Count(text, " ")) == 0 {
		return 0, fmt.Errorf("frase inexistente")
	}
	partes := strings.Split(text, " ")
	soma := 0.0
	contador := 0.0
	media := 0.0
	for i := 0; i < len(partes); i++ {
		if len(partes[i]) >= 2 {
			fmt.Println(partes[i])
			soma += float64(len(partes[i]))
			contador++
		}
	}
	media = soma / contador
	return media, nil
}
