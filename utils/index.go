package utils

import (
	"encoding/json"
	"fmt"
)

func StructToString(data any) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(bytes)
}
