package middleware

import (
	"encoding/json"
)

// interfaceToString is a quick helper to convert from an ambiguous string to a real one
func interfaceToString(toConvert interface{}) string {
	switch a := toConvert.(type) {
	case string:
		return a
	}
	return ""
}

func interfaceToInt(toConvert interface{}) int {
	switch a := toConvert.(type) {
	case int:
		return a
	case float64:
		return int(a)
	}
	return -1
}

func indentJSON(response interface{}) string {
	resp, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(resp)
}

