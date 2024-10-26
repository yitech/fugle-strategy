package pythagoras

import "github.com/yitech/fugle-strategy/internal/strategy"

var _ strategy.Strategy = (*Pythagoras)(nil)

type Pythagoras struct {
	name string
}

func (p *Pythagoras) Name() string {
	return p.name
}

func (p *Pythagoras) Family() string {
	return "Pythagoras"
}

func (p *Pythagoras) Status() string {
	return "OK"
}

func (p *Pythagoras) Destroy() {
	// destroy
}
