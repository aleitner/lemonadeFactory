package factory

type LemonadeType int

const (
	PLAIN LemonadeType = iota
	PINK
	LIME
)

type Lemonade struct {
	lemonadeType LemonadeType
	amount int
	size int
}

func NewLemonade(lemonadeType LemonadeType, size int) *Lemonade {
	return &Lemonade{
		lemonadeType: lemonadeType,
		size: size,
		amount: size,
	}
}

func (l *Lemonade) Drink() {
	if l.amount > 0 {
		l.amount--
	}
}

