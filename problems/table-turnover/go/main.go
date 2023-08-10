package main

// - We consider a round table with seats numbered clockwise
// - The number of seats around the table is given by the integer 'seats'
// - We consider a customer who is sitting at this table at seat #'customer'
// - In front of each seat may be: an starter [E], a main course [P] and one dessert [D] on the table
// - The costumer can turn the table to the left or right, one seat at a time
// - The meal ends when the costumer has eaten a starter, a main course and a dessert, in that order

// Implement function that:
// - takes as inputs the integers `seats` and `customer` and the string `dishes`,
//		composed of the characters `E, P, D, N` and representing the dishes layout
// - returns the minimum number of rotations at the end of the meal, as an integer

// We colnsider:
// - 0 < seats < 20
// - 0 <= number of characters in `dishes` < 20

// examples
// seats: 7, customer: 1, dishes: ENPDNNN = 4 rotations

func main() {

}

// epd is the eaten order so loop through it
// for each course calculate the shorts path from customer position
// acumulate turns
// having dishes positions in a map can facilitate the calculation
func CalculateTurns(seats int, customer int, dishes string) int {
	if seats != len(dishes) {
		return 0
	}

	dishesInfo := make(map[rune]int, 0)
	for idx, dish := range dishes {
		if dish == 'N' {
			continue
		}
		dishesInfo[dish] = idx
	}

	total := 0

	for _, dish := range []rune{'E', 'P', 'D'} {
		position := dishesInfo[dish]
		rotations, reverse := calculateRotations(seats, customer, position)
		if rotations == 0 {
			continue
		}
		total += rotations
		rotate(dishesInfo, rotations, reverse, seats)
	}

	return total
}

func calculateRotations(seats int, customer int, position int) (int, bool) {
	turnsClockwise := (customer - position + seats) % seats
	turnsCounterClockwise := (position - customer + seats) % seats
	return min(turnsClockwise, turnsCounterClockwise)
}

func rotate(dishesInfo map[rune]int, rotations int, reverse bool, seats int) {
	for dish, p := range dishesInfo {
		dishesInfo[dish] = newPosition(p, rotations, reverse, seats)
	}
}

func newPosition(p int, rotations int, reverse bool, seats int) int {
	var newP int
	lastP := seats - 1
	if reverse {
		newP = p - rotations
		if newP < 0 {
			newP = seats + newP // +(-X)
		}
	} else {
		newP = p + rotations
		if newP > lastP {
			newP = newP - seats
		}
	}

	return newP
}

// bool represents reverse/counter
func min(a, b int) (int, bool) {
	if a < b {
		return a, false
	}
	return b, true
}
