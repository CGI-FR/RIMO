package model

import (
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

func GetJSONSchema() (string, error) {
	resBytes, err := json.MarshalIndent(jsonschema.Reflect(&Base{}), "", "  ") //nolint:exhaustruct
	if err != nil {
		return "", fmt.Errorf("couldn't unmarshall Base in JSON : %w", err)
	}

	return string(resBytes), nil
}
