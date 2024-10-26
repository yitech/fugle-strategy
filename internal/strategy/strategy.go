package strategy

type Strategy interface {
	Name() string
	Family() string
	Status() string
	Destroy()
}
