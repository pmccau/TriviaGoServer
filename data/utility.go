package data

import (
	"crypto/rand"
	"fmt"
)

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