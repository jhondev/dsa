package main

import "fmt"

func main() {

}

func romanToInt(input string) (int, error) {
	romans := map[string]int{
		"I":  1,
		"IV": 4,
		"IX": 9,
		"V":  5,
		"X":  10,
		"XL": 40,
		"XC": 90,
		"L":  50,
		"C":  100,
		"CD": 400,
		"CM": 900,
		"D":  500,
		"M":  1000,
	}
	var number int

	for i := 0; i < len(input); i++ {
		nexti := i + 1
		// look first for pair roman valid value if possible, example: IV, IX, etc
		if nexti < len(input) {
			rint, ok := romans[string(input[i])+string(input[nexti])]
			if ok {
				i++
				number = number + rint
				continue
			}
		}
		// if there is no pair roman value try single roman
		rint, ok := romans[string(input[i])]
		if !ok {
			return 0, fmt.Errorf("invalid roman number %v", input[i])
		}
		number = number + rint
	}

	return number, nil
}
