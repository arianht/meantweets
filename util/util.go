// Package util contains utility types and functions.
package util

import (
	"encoding/json"
	"io/ioutil"
)

// TwitterAPICredentials represents the API credentials needed to make API requests
// to Twitter.
type TwitterAPICredentials struct {
	ConsumerKey    string
	ConsumerSecret string
}

// ReadTwitterAPICredentials parses Twitter API credentials and returns them.
func ReadTwitterAPICredentials(data []byte) (credentials TwitterAPICredentials, err error) {
	err = json.Unmarshal(data, &credentials)
	return
}

// GetTwitterAPICredentialsFromFile reads a JSON file containing Twitter API credentials
// and returns the credentials.
func GetTwitterAPICredentialsFromFile(filename string) (credentials TwitterAPICredentials, err error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	credentials, err = ReadTwitterAPICredentials(fileData)
	return
}
