package q3

import "fmt"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("lista vazia")
	}
	mai := 0
	men := 0
	soma := 0
	media := 0.0
	mai = numbers[0]
	men = numbers[0]
	for i := 0; i < len(numbers); i++ {
		soma += numbers[i]
		if mai < numbers[i] {
			mai = numbers[i]
		}
		if men > numbers[i] {
			men = numbers[i]
		}
	}
	media = float64(soma-(mai+men)) / float64(len(numbers)-2)

	return mai, men, media, nil
}
