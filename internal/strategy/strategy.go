package strategy

type Strategy interface {
	Run() error
	Name() string
	Family() string
	Status() string
	Destroy()
}
