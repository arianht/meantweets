package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestServerGetTweetsEndpoint(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatalf("Could not get a context - %v", err)
	}
	defer done()

	request, _ := http.NewRequest("GET", "/get_tweets?test=one", nil)
	recorder := httptest.NewRecorder()
	tweetsHandler(ctx, recorder, request)

	if response := recorder.Body.String(); response != "OK" {
		t.Errorf("Expected body to be OK, but was %v", response)
	}
}
