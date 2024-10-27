package pythagoras

import (
	fuglego "github.com/yitech/fugle-go"
	"github.com/yitech/fugle-strategy/internal/strategy"
)

var _ strategy.Strategy = (*Pythagoras)(nil)

type Pythagoras struct {
	name        string
	fugleClient fuglego.APIClient
}

func NewPythagoras(name string, fugleClient *fuglego.APIClient) *Pythagoras {
	return &Pythagoras{
		name:        name,
		fugleClient: *fugleClient,
	}
}

func (p *Pythagoras) Name() string {
	return p.name
}

func (p *Pythagoras) Run() error {
	return nil
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
