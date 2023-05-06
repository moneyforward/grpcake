package grpcake

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JsonPrettify adds indent to raw json string.
func JsonPrettify(jsonBytes []byte) (string, error) {
	var prettifiedJson bytes.Buffer
	err := json.Indent(&prettifiedJson, jsonBytes, "", "\t")
	if err != nil {
		return "", fmt.Errorf("error prettify-ing jsong string: %v", err)
	}

	return prettifiedJson.String(), nil
}
