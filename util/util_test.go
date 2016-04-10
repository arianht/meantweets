package util

import (
	"fmt"
	"testing"
)

func TestReadingTwitterAPICredentialsFromJSON(t *testing.T) {
	expectedConsumerKey := "test_key"
	expectedConsumerSecret := "test_secret"
	inputJSON := fmt.Sprintf("{\"ConsumerKey\":\"%v\",\"ConsumerSecret\":\"%v\"}",
		expectedConsumerKey, expectedConsumerSecret)

	credentials, err := ReadTwitterAPICredentials([]byte(inputJSON))

	if err != nil {
		t.Fatalf("Expected nil error, but was %v", err)
	}
	if credentials.ConsumerKey != expectedConsumerKey {
		t.Errorf("Expected consumer key %v, but found %v", expectedConsumerKey, credentials.ConsumerKey)
	}
	if credentials.ConsumerSecret != expectedConsumerSecret {
		t.Errorf("Expected consumer secret %v, but found %v", expectedConsumerSecret, credentials.ConsumerSecret)
	}
}
