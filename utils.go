package grpcake

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONPrettify adds indent to raw json string.
func JSONPrettify(jsonBytes []byte) (string, error) {
	var prettifiedJSON bytes.Buffer
	err := json.Indent(&prettifiedJSON, jsonBytes, "", "\t")
	if err != nil {
		return "", fmt.Errorf("error prettify-ing jsong string: %v", err)
	}

	return prettifiedJSON.String(), nil
}
