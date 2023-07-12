package config

// TODO Create a new struct for select and configure swap functions
type Swap struct {
	Mode string `json:"mode" yaml:"mode" mapstructure:"mode" validate:"required,oneof=slot front"`

	FrontDoor *FrontDoor `json:"front_door,omitempty" yaml:"front_door,omitempty" mapstructure:"front_door,omitempty" validate:"required_if=Mode front"`

	AppInsights *AppInsights `json:"app_insights,omitempty" yaml:"app_insights,omitempty" mapstructure:"app_insights,omitempty" validate:"required_if=Mode front"`
}

type FrontDoor struct {
	ResourceGroup string `json:"resource_group" yaml:"resource_group" mapstructure:"resource_group" validate:"required"`
	Name          string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
	OriginGroup   string `json:"origin_group" yaml:"origin_group" mapstructure:"origin_group" validate:"required"`
}

type AppInsights struct {
	ResourceGroup string `json:"resource_group" yaml:"resource_group" mapstructure:"resource_group" validate:"required"`
	Name          string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
}

func NewSwap() *Swap {
	return &Swap{}
}
