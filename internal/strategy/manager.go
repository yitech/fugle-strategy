package strategy

import (
	"fmt"
)

type StrategyManager struct {
	Strategies map[string]Strategy
}

type Notify struct {
	Name   string
	Status string
}

func NewStrategyManager() *StrategyManager {
	return &StrategyManager{
		Strategies: make(map[string]Strategy),
	}
}

func (m *StrategyManager) AddStrategy(name string, strategy Strategy) error {
	// check if name already exists
	if _, ok := m.Strategies[name]; ok {
		m.Strategies[name] = strategy
		return nil
	} else {
		return fmt.Errorf("strategy with name %s already exists", name)
	}
}

func (m *StrategyManager) RemoveStrategy(name string) error {
	if _, ok := m.Strategies[name]; ok {
		m.Strategies[name].Destroy()
		delete(m.Strategies, name)
		return nil
	} else {
		return fmt.Errorf("strategy with name %s does not exist", name)
	}
}

func (m *StrategyManager) Close(name string, status string) {

}
