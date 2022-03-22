package factory

type LemonadeFactory struct {
	lemonadeType LemonadeType
	bottleSize int
	ProductionCount int
}

func NewLemonadeFactory(size int, lemonadeType LemonadeType) *LemonadeFactory {
	return &LemonadeFactory{
		lemonadeType: lemonadeType,
		bottleSize: size,
	}
}

func (lf *LemonadeFactory) NewLemonade() *Lemonade {
	lf.ProductionCount++
	return NewLemonade(lf.lemonadeType, lf.bottleSize)
}