package config

type Cdn struct {
	ResourceGroup string `json:"resource_group,omitempty" yaml:"resource_group,omitempty" mapstructure:"resource_group,omitempty" validate:"required_if=Cloud azure,excluded_with=Account Project"`
	Account       string `json:"account,omitempty" yaml:"account,omitempty" mapstructure:"account,omitempty" validate:"required_if=Cloud aws,excluded_with=ResourceGroup Project"`
	Project       string `json:"project,omitempty" yaml:"project,omitempty" mapstructure:"project,omitempty" validate:"required_if=Cloud gcp,excluded_with=ResourceGroup Account"`

	Name      string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
	Endpoint  string `json:"endpoint" yaml:"endpoint" mapstructure:"endpoint" validate:"required"`
	Domain    string `json:"domain" yaml:"domain" mapstructure:"domain" validate:"required"`
	CachePath string `json:"cache_path" yaml:"cache_path" mapstructure:"cache_path" validate:"required"`
}
