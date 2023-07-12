package config

type Testing struct {
	Repository string `json:"repository" yaml:"repository" mapstructure:"repository" validate:"required"`
}
