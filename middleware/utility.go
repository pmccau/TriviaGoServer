package middleware

import (
	"encoding/json"
	"net/http"
)

// interfaceToString is a quick helper to convert from an ambiguous string to a verified one
func interfaceToString(toConvert interface{}) string {
	switch a := toConvert.(type) {
	case string:
		return a
	}
	return ""
}

// interfaceToInt is a quick helper to convert from an ambiguous int to a verified one
func interfaceToInt(toConvert interface{}) int {
	switch a := toConvert.(type) {
	case int:
		return a
	case float64:
		return int(a)
	}
	return -1
}

// indentJSON properly indents an interface as a stringified JSON
func indentJSON(response interface{}) string {
	resp, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(resp)
}

// SetResponseHeaders will set the headers appropriately before sending back a response
// for all middleware components
func SetResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Context-Type", "application/results-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
