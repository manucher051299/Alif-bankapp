package card

import (
	"bank/pkg/bank/types"
)

const withdrawLimit = 20_000_00
const depositLimit = 50_000_00

func Withdraw(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}
	if card.Balance < amount {
		return
	}
	if amount < 0 {
		return
	}
	if amount > withdrawLimit {
		return
	}
	card.Balance -= amount
}
func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}
	if amount > depositLimit {
		return
	}
	if amount < 0 {
		return
	}
	card.Balance += amount
}
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}
	if card.Balance <= 0 {
		return
	}
	bonus := int(card.MinBalance) * percent / 100 * daysInMonth / daysInYear
	if bonus > 5_000_00 {
		return
	}
	card.Balance += types.Money(bonus)
}
func Total(cards []types.Card) types.Money {
	var sum types.Money
	for _, card := range cards {
		if card.Balance <= 0 || !card.Active {
			continue
		}
		sum += card.Balance
	}
	return sum
}
func PaymentSources(cards []types.Card) []types.PaymentSource {
	// TODO: ваш код
	var a []types.PaymentSource
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			a = append(a, types.PaymentSource{Type: "card", Number: "5058 xxxx xxxx 8888", Balance: card.Balance})
		}
	}
	return a
}
