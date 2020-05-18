package middleware

import (
	"encoding/json"
	"fmt"
	"crypto/rand"
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

// GenerateGuid generates a simple GUID
// Credit: https://yourbasic.org/golang/generate-uuid-guid/
func GenerateGuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return string(uuid)
}