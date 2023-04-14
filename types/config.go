package types

type Config struct {
	Env map[string][]Function `yaml:"environments" mapstructure:"environments" validate:"required,dive,dive"`
}

type Function struct {
	Name        string `yaml:"name" validate:"required"`
	PackagePath string `yaml:"package_path" mapstructure:"package_path" validate:"required,dirpath"`
	Region      string `yaml:"region" mapstructure:"region" validate:"required"`
	Cloud       string `yaml:"cloud" mapstructure:"cloud" validate:"required,oneof=azure aws gcp"`
	Runtime     string `yaml:"runtime" mapstructure:"runtime" validate:"required"`

	ResourceGroup string `yaml:"resource_group,omitempty" mapstructure:"resource_group,omitempty" validate:"required_if=Cloud azure,excluded_with=Account Project"`
	Account       string `yaml:"account,omitempty" mapstructure:"account,omitempty" validate:"required_if=Cloud aws,excluded_with=ResourceGroup Project"`
	Project       string `yaml:"project,omitempty" mapstructure:"project,omitempty" validate:"required_if=Cloud gcp,excluded_with=ResourceGroup Account"`

	StorageAccount string `yaml:"storage_account,omitempty" mapstructure:"storage_account,omitempty" validate:"required_if=Cloud azure,excluded_with=S3 GoogleStorage"`
	S3             string `yaml:"s3,omitempty" mapstructure:"s3,omitempty" validate:"required_if=Cloud aws,excluded_with=StorageAccount GoogleStorage"`
	GoogleStorage  string `yaml:"google_storage,omitempty" mapstructure:"google_storage,omitempty" validate:"required_if=Cloud gcp,excluded_with=StorageAccount S3"`

	Environment []string `yaml:"environment,omitempty" mapstructure:"environment,omitempty" validate:"omitempty"`

	Secrets []Secret `yaml:"secrets,omitempty" mapstructure:"secrets,omitempty" validate:"omitempty,excluded_if=Cloud azure,excluded_if=Cloud aws,dive"`
}

type Secret struct {
	Name    string `yaml:"name,omitempty" mapstructure:"name,omitempty" validate:"required"`
	Version string `yaml:"version,omitempty" mapstructure:"version,omitempty" validate:"required"`
	Path    string `yaml:"path,omitempty" mapstructure:"path,omitempty" validate:"required,filepath"`
}
