package modelv2

type Config struct {
	Concept      string   `json:"concept"      yaml:"concept"      jsonschema:"required"`
	Constraint   []string `json:"constraint"   yaml:"constraint"   jsonschema:"required"`
	Confidential *bool    `json:"confidential" yaml:"confidential" jsonschema:"required"`
}
