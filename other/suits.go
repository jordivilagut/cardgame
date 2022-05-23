package other

type suits int

const (
	CLUBS suits = iota
	DIAMONDS
	HEARTS
	SPADES
)

func (suit suits) toString() string {
	switch suit {
	case CLUBS:
		return "Clubs"
	case DIAMONDS:
		return "Diamonds"
	case HEARTS:
		return "Hearts"
	case SPADES:
		return "Spades"
	default:
		return ""
	}
}
