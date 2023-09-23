package config

type Cdn struct {
	Path		string 		`json:"path" yaml:"path" mapstructure:"path" validate:"required"`
	FrontDoor   *FrontDoor 	`json:"front-door,omitempty" yaml:"front-door,omitempty" mapstructure:"front-door,omitempty" validate:"omitempty,dive"`
	Akamai		*Akamai		`json:"akamai,omitempty" yaml:"akamai,omitempty" mapstructure:"akamai,omitempty" validate:"omitempty,dive"`
	CloudFront  *CloudFront	`json:"cloudfront,omitempty" yaml:"cloudfront,omitempty" mapstructure:"cloudfront,omitempty" validate:"omitempty,dive"`
}
type FrontDoor struct {
	ResourceGroup string `json:"resource_group" yaml:"resource_group" mapstructure:"resource_group" validate:"required"`
	Name      string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
	Endpoint  string `json:"endpoint" yaml:"endpoint" mapstructure:"endpoint" validate:"required"`
	Domain    string `json:"domain" yaml:"domain" mapstructure:"domain" validate:"required"`
}
type Akamai struct {
	Domain    string `json:"domain" yaml:"domain" mapstructure:"domain" validate:"required"`
}
type CloudFront struct {
	Domain    string `json:"domain" yaml:"domain" mapstructure:"domain" validate:"required"`
	Account       string `json:"account" yaml:"account" mapstructure:"account" validate:"required"`
}