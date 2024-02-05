package mapping

import "encoding/json"

// MapToStruct map to struct
func MapToStruct(mapValue map[string]interface{}, structValue interface{}) error {

	bytes, err := json.Marshal(mapValue)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, structValue)
	if err != nil {
		return err
	}

	return nil
}
