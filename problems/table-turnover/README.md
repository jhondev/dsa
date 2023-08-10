# Table Turnover

A restaurant wants to calculate how many times a customer should turn the tray depending on their seat and the dishes layout

- We consider a round table with seats numbered clockwise
- The number of seats around the table is given by the integer 'seats'
- We consider a customer who is sitting at this table at seat #'customer'
- in front of each seat may be: an starter [E], a main course [P] and one dessert [D] on the table
- The costumer can turn the table to the left or right, one seat at a time
- The meal ends when the costumer has eaten a starter, a main course and a dessert, in that order

Implement the function `CalculateTurns(seats int, customer int, dishes string) int` which:

- takes as inputs the integers `seats` and `customer` and the string `dishes`, composed of the characters `E, P, D, N` and representing the dishes layout
- returns the minimum number of rotations at the end of the meal, as an integer

We colnsider:

- 0 < seats < 20
- 0 <= number of characters in `dishes` < 20

## Reference

https://www.theforkmanager.com/blog/restaurant-management/table-turnover-how-calculate-it-and-optimize-it
