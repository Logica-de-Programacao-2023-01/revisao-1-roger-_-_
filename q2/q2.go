package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	caracteresp := []string{"@", "!", ".", ",", "#", "$", "%", "*", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, caracter := range caracteresp {
		text = strings.ReplaceAll(text, caracter, "")
	}

	palavras := strings.Fields(text)

	if len(palavras) == 0 {
		return 0, fmt.Errorf("frase inexistente")
	}
	total_letras := 0
	for _, palavra := range palavras {
		total_letras += len(palavra)
	}
	var media float64
	media = float64(total_letras) / float64(len(palavras))
	return media, nil
}
