package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleAddBonus() {
	card := types.Card{Balance: 10_000_00, Active: true, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1002465
}
func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 10_000_00, Active: false, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1000000
}
func ExampleAddBonus_negative() {
	card := types.Card{Balance: -10_000_00, Active: true, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: -1000000
}
func ExampleAddBonus_limit() {
	card := types.Card{Balance: 10_000_000_00, Active: true, MinBalance: 10_000_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1000000000
}
