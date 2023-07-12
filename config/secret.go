package config

type Secret struct {
	Name    string `json:"name" yaml:"name,omitempty" mapstructure:"name,omitempty" validate:"required"`
	Version string `json:"version" yaml:"version,omitempty" mapstructure:"version,omitempty" validate:"required"`
	Path    string `json:"path" yaml:"path,omitempty" mapstructure:"path,omitempty" validate:"required,filepath"`
}
