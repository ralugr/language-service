package adapters

import (
	"encoding/json"
	"github.com/ralugr/language-service/logger"
)

// ConvertFromByteToStringArray converts a byte array to a string array
func ConvertFromByteToStringArray(b []byte) ([]string, error) {
	var stringList []string

	if err := json.Unmarshal(b, &stringList); err != nil {
		logger.Warning.Printf("Unmarshall failed for %v, error %v", string(b), err)
		return nil, err
	}
	logger.Info.Printf("Converted to %v", stringList)
	return stringList, nil
}
