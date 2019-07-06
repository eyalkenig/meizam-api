package api_test

import (
	"errors"
	"fmt"
	"os"
)

var uri string

func getHTTPURI() string {
	if uri != "" {
		return uri
	}
	envVarName := "GATEWAY_URI"
	uri, ok := os.LookupEnv(envVarName)
	if !ok || uri == "" {
		panic(errors.New(fmt.Sprintf("missing environment variable %s", envVarName)))
	}
	return uri
}
