package q5

import "fmt"

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	var bases = []string{"C", "K", "F"}
	para1 := 0
	para2 := 0
	for i := 0; i < 3; i++ {
		if fromScale != bases[i] {
			para1++
		}
		if toScale != bases[i] {
			para2++
		}
	}
	if para1 != 0 || len(fromScale) == 0 {
		return 0, fmt.Errorf("escala inválida")
	}
	if para2 != 0 || len(toScale) == 0 {
		return 0, fmt.Errorf("escala inválida")
	}

	temperatura := 0.0

	switch fromScale {
	case "C":
		switch toScale {
		case "F":
			temperatura = temp*9/5 + 32
		case "K":
			temperatura = temp + 273.15
		}
	case "K":
		switch toScale {
		case "F":
			temperatura = (temp-273.15)*9/5 + 32
		case "C":
			temperatura = temp - 273.15
		}
	case "F":
		switch toScale {
		case "C":
			temperatura = (temp - 32) * 5 / 9
		case "K":
			temperatura = (temp-32)*5/9 + 273.15
		}
	}

	return temperatura, nil
}
