package config

type Function struct {
	Name        string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
	Type        string `json:"type" yaml:"type" mapstructure:"type" validate:"required,oneof=back front"`
	PackagePath string `json:"package_path" yaml:"package_path" mapstructure:"package_path" validate:"required"`
	Region      string `json:"region" yaml:"region" mapstructure:"region" validate:"required"`
	Cloud       string `json:"cloud" yaml:"cloud" mapstructure:"cloud" validate:"required,oneof=azure aws gcp"`
	Runtime     string `json:"runtime" yaml:"runtime" mapstructure:"runtime" validate:"required"`

	ArtifactId string `json:"artifact_id" yaml:"artifact_id" mapstructure:"artifact_id" validate:"required"`

	Profile string `json:"profile" yaml:"profile" mapstructure:"profile" validate:"required"`

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

	Vault *Vault `json:"vault,omitempty" yaml:"vault,omitempty" mapstructure:"vault,omitempty" validate:"omitempty,dive"`

	Testing *Testing `json:"testing,omitempty" yaml:"testing,omitempty" mapstructure:"testing,omitempty" validate:"omitempty,dive"`

	Swap *Swap `json:"swap" yaml:"swap" mapstructure:"swap" validate:"required_if=Type back,dive"`
}
