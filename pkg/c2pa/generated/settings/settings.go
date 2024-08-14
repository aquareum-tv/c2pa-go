// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package settings

type Builder struct {
	// AutoThumbnail corresponds to the JSON schema field "auto_thumbnail".
	AutoThumbnail bool `json:"auto_thumbnail" yaml:"auto_thumbnail" mapstructure:"auto_thumbnail"`
}

type Core struct {
	// CompressManifests corresponds to the JSON schema field "compress_manifests".
	CompressManifests bool `json:"compress_manifests" yaml:"compress_manifests" mapstructure:"compress_manifests"`

	// Debug corresponds to the JSON schema field "debug".
	Debug bool `json:"debug" yaml:"debug" mapstructure:"debug"`

	// HashAlg corresponds to the JSON schema field "hash_alg".
	HashAlg string `json:"hash_alg" yaml:"hash_alg" mapstructure:"hash_alg"`

	// MaxMemoryUsage corresponds to the JSON schema field "max_memory_usage".
	MaxMemoryUsage *int `json:"max_memory_usage,omitempty" yaml:"max_memory_usage,omitempty" mapstructure:"max_memory_usage,omitempty"`

	// PreferBmffMerkleTree corresponds to the JSON schema field
	// "prefer_bmff_merkle_tree".
	PreferBmffMerkleTree bool `json:"prefer_bmff_merkle_tree" yaml:"prefer_bmff_merkle_tree" mapstructure:"prefer_bmff_merkle_tree"`

	// PreferBoxHash corresponds to the JSON schema field "prefer_box_hash".
	PreferBoxHash bool `json:"prefer_box_hash" yaml:"prefer_box_hash" mapstructure:"prefer_box_hash"`

	// SaltJumbfBoxes corresponds to the JSON schema field "salt_jumbf_boxes".
	SaltJumbfBoxes bool `json:"salt_jumbf_boxes" yaml:"salt_jumbf_boxes" mapstructure:"salt_jumbf_boxes"`
}

type SettingsSchemaJson struct {
	// Builder corresponds to the JSON schema field "builder".
	Builder Builder `json:"builder" yaml:"builder" mapstructure:"builder"`

	// Core corresponds to the JSON schema field "core".
	Core Core `json:"core" yaml:"core" mapstructure:"core"`

	// Trust corresponds to the JSON schema field "trust".
	Trust Trust `json:"trust" yaml:"trust" mapstructure:"trust"`

	// Verify corresponds to the JSON schema field "verify".
	Verify Verify `json:"verify" yaml:"verify" mapstructure:"verify"`
}

type Trust struct {
	// AllowedList corresponds to the JSON schema field "allowed_list".
	AllowedList *string `json:"allowed_list,omitempty" yaml:"allowed_list,omitempty" mapstructure:"allowed_list,omitempty"`

	// PrivateAnchors corresponds to the JSON schema field "private_anchors".
	PrivateAnchors *string `json:"private_anchors,omitempty" yaml:"private_anchors,omitempty" mapstructure:"private_anchors,omitempty"`

	// TrustAnchors corresponds to the JSON schema field "trust_anchors".
	TrustAnchors *string `json:"trust_anchors,omitempty" yaml:"trust_anchors,omitempty" mapstructure:"trust_anchors,omitempty"`

	// TrustConfig corresponds to the JSON schema field "trust_config".
	TrustConfig *string `json:"trust_config,omitempty" yaml:"trust_config,omitempty" mapstructure:"trust_config,omitempty"`
}

type Verify struct {
	// CheckIngredientTrust corresponds to the JSON schema field
	// "check_ingredient_trust".
	CheckIngredientTrust bool `json:"check_ingredient_trust" yaml:"check_ingredient_trust" mapstructure:"check_ingredient_trust"`

	// OcspFetch corresponds to the JSON schema field "ocsp_fetch".
	OcspFetch bool `json:"ocsp_fetch" yaml:"ocsp_fetch" mapstructure:"ocsp_fetch"`

	// RemoteManifestFetch corresponds to the JSON schema field
	// "remote_manifest_fetch".
	RemoteManifestFetch bool `json:"remote_manifest_fetch" yaml:"remote_manifest_fetch" mapstructure:"remote_manifest_fetch"`

	// VerifyAfterReading corresponds to the JSON schema field "verify_after_reading".
	VerifyAfterReading bool `json:"verify_after_reading" yaml:"verify_after_reading" mapstructure:"verify_after_reading"`

	// VerifyAfterSign corresponds to the JSON schema field "verify_after_sign".
	VerifyAfterSign bool `json:"verify_after_sign" yaml:"verify_after_sign" mapstructure:"verify_after_sign"`

	// VerifyTrust corresponds to the JSON schema field "verify_trust".
	VerifyTrust bool `json:"verify_trust" yaml:"verify_trust" mapstructure:"verify_trust"`
}