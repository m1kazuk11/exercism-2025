package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "king", "queen", "jack":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	myCard1 := ParseCard(card1)
	myCard2 := ParseCard(card2)
	myCardNum := myCard1 + myCard2
	dealerCardNum := ParseCard(dealerCard)
	switch {
	case myCard1 == 11 && myCard2 == 11:
		return "P"
	case myCardNum == 21 && (dealerCardNum == 10 || dealerCardNum == 11):
		return "S"
	case myCardNum == 21 && dealerCardNum != 10:
		return "W"
	case myCardNum >= 17 && myCardNum <= 20:
		return "S"
	case myCardNum >= 12 && myCardNum <= 16 && dealerCardNum >= 7:
		return "H"
	case myCardNum >= 12 && myCardNum <= 16:
		return "S"
	case myCardNum <= 11:
		return "H"
	}
	return "Invalid Cards"
}
