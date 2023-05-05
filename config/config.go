package config

type Config struct {
	Version string                `yaml:"version" mapstructure:"version" validate:"required"`
	Env     map[string][]Function `yaml:"environments" mapstructure:"environments" validate:"required,dive,dive"`
}

type Function struct {
	Name        string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
	Type        string `json:"type" yaml:"type" mapstructure:"type" validate:"required,oneof=back front"`
	PackagePath string `json:"package_path" yaml:"package_path" mapstructure:"package_path" validate:"required,dirpath"`
	Region      string `json:"region" yaml:"region" mapstructure:"region" validate:"required"`
	Cloud       string `json:"cloud" yaml:"cloud" mapstructure:"cloud" validate:"required,oneof=azure aws gcp"`
	Runtime     string `json:"runtime" yaml:"runtime" mapstructure:"runtime" validate:"required"`

	ResourceGroup string `json:"resource_group,omitempty" yaml:"resource_group,omitempty" mapstructure:"resource_group,omitempty" validate:"required_if=Cloud azure,excluded_with=Account Project"`
	Account       string `json:"account,omitempty" yaml:"account,omitempty" mapstructure:"account,omitempty" validate:"required_if=Cloud aws,excluded_with=ResourceGroup Project"`
	Project       string `json:"project,omitempty" yaml:"project,omitempty" mapstructure:"project,omitempty" validate:"required_if=Cloud gcp,excluded_with=ResourceGroup Account"`

	StorageAccount string `json:"storage_account,omitempty" yaml:"storage_account,omitempty" mapstructure:"storage_account,omitempty" validate:"required_if=Cloud azure Type front,excluded_with=S3 GoogleStorage"`
	S3             string `json:"s3,omitempty" yaml:"s3,omitempty" mapstructure:"s3,omitempty" validate:"required_if=Cloud aws Type front,excluded_with=StorageAccount GoogleStorage"`
	GoogleStorage  string `json:"google_storage,omitempty" yaml:"google_storage,omitempty" mapstructure:"google_storage,omitempty" validate:"required_if=Cloud gcp Type front,excluded_with=StorageAccount S3"`

	Environments []string `json:"environments,omitempty" yaml:"environments,omitempty" mapstructure:"environments,omitempty" validate:"omitempty"`

	Secrets []*Secret `json:"secrets,omitempty" yaml:"secrets,omitempty" mapstructure:"secrets,omitempty" validate:"omitempty,excluded_if=Cloud azure,excluded_if=Cloud aws,dive"`

	Functions *[]string `json:"functions,omitempty" yaml:"functions,omitempty" mapstructure:"functions" validate:"omitempty"`

	Cdn *Cdn `json:"cdn,omitempty" yaml:"cdn,omitempty" mapstructure:"cdn,omitempty" validate:"omitempty,dive"`

	Database *Database `json:"db,omitempty" yaml:"db,omitempty" mapstructure:"db,omitempty" validate:"omitempty,dive"`

	Vault *Vault `json:"vault,omitempty" yaml:"db,omitempty" mapstructure:"db,omitempty" validate:"omitempty,dive"`
}

type Secret struct {
	Name    string `json:"name" yaml:"name,omitempty" mapstructure:"name,omitempty" validate:"required"`
	Version string `json:"version" yaml:"version,omitempty" mapstructure:"version,omitempty" validate:"required"`
	Path    string `json:"path" yaml:"path,omitempty" mapstructure:"path,omitempty" validate:"required,filepath"`
}

type Cdn struct {
	ResourceGroup string `json:"resource_group,omitempty" yaml:"resource_group,omitempty" mapstructure:"resource_group,omitempty" validate:"required_if=Cloud azure,excluded_with=Account Project"`
	Account       string `json:"account,omitempty" yaml:"account,omitempty" mapstructure:"account,omitempty" validate:"required_if=Cloud aws,excluded_with=ResourceGroup Project"`
	Project       string `json:"project,omitempty" yaml:"project,omitempty" mapstructure:"project,omitempty" validate:"required_if=Cloud gcp,excluded_with=ResourceGroup Account"`

	Name      string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
	Endpoint  string `json:"endpoint" yaml:"endpoint" mapstructure:"endpoint" validate:"required"`
	Domain    string `json:"domain" yaml:"domain" mapstructure:"domain" validate:"required"`
	CachePath string `json:"cache_path" yaml:"cache_path" mapstructure:"cache_path" validate:"required"`
}

type Database struct {
	Enabled bool `json:"enabled" yaml:"enabled" mapstructure:"enabled" validate:"required,boolean"`

	ResourceGroup string `json:"resource_group,omitempty" yaml:"resource_group,omitempty" mapstructure:"resource_group,omitempty" validate:"required_if=Cloud azure,excluded_with=Account Project"`
	Account       string `json:"account,omitempty" yaml:"account,omitempty" mapstructure:"account,omitempty" validate:"required_if=Cloud aws,excluded_with=ResourceGroup Project"`
	Project       string `json:"project,omitempty" yaml:"project,omitempty" mapstructure:"project,omitempty" validate:"required_if=Cloud gcp,excluded_with=ResourceGroup Account"`

	Name string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
	Type string `json:"type" yaml:"type" mapstructure:"type" validate:"required,oneof=postgresql mysql mongodb"`
}

type Vault struct {
	Enabled bool `json:"enabled" yaml:"enabled" mapstructure:"enabled" validate:"required,boolean"`

	ResourceGroup string `json:"resource_group,omitempty" yaml:"resource_group,omitempty" mapstructure:"resource_group,omitempty" validate:"required_if=Cloud azure,excluded_with=Account Project"`
	Account       string `json:"account,omitempty" yaml:"account,omitempty" mapstructure:"account,omitempty" validate:"required_if=Cloud aws,excluded_with=ResourceGroup Project"`
	Project       string `json:"project,omitempty" yaml:"project,omitempty" mapstructure:"project,omitempty" validate:"required_if=Cloud gcp,excluded_with=ResourceGroup Account"`

	Name string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
}
