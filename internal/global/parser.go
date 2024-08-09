package global

import (
	"encoding/json"
)

func StructToJson(val interface{}) string {
	b, err := json.Marshal(val)
	if err != nil {
		return err.Error()
	}

	return string(b)

}
