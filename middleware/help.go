package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func getHelpResponse() *map[string]interface{} {
	data, err := ioutil.ReadFile("assets/help.json")
	if err != nil {
		panic(err)
	}
	var output map[string]interface{}
	err = json.Unmarshal(data, &output)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("Output:", output)
	return &output
}