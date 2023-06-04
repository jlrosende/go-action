package config

// TODO Create a new struct for select and configure swap functions
type Swap struct {
	Mode   string
	Source string
	Target string
}

func NewSwap() *Swap {
	return &Swap{}
}
