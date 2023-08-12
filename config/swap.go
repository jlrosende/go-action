package config

type Swap struct {
	Mode string `json:"mode" yaml:"mode" mapstructure:"mode" validate:"required,oneof=slot traffic custom"`
	Canary string `json:"canary_time,omitempty" yaml:"canary_time,omitempty" mapstructure:"canary_time,omitempty" validate:"omitempty"`
	Custom string `json:"custom_mode,omitempty" yaml:"custom_mode,omitempty" mapstructure:"custom_mode,omitempty" validate:"required_if=Mode custom,omitempty"`
}

func NewSwap() *Swap {
	return &Swap{}
}
