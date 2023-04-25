package grpcake

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JsonPrettify adds indent to raw json string.
func JsonPrettify(jsonString string) (string, error) {
	var prettifiedJson bytes.Buffer
	err := json.Indent(&prettifiedJson, []byte(jsonString), "", "\t")
	if err != nil {
		return "", fmt.Errorf("error prettify-ing jsong string: %v", err)
	}

	return prettifiedJson.String(), nil
}
