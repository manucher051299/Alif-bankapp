package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

/*
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
*/
func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active:  true,
		},
	}))
	// Output:
	// 100000
	// 300000
	// 0
	// 0
}

func ExamplePaymentSources() {
	PaymentSources([]types.Card{
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 1_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 9999",
			Balance: 2_000_00,
			Active:  true,
		},
	})

	// Output:
	//5058 xxxx xxxx 8888
	//5058 xxxx xxxx 9999
}
