package crawl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/kurrik/twittergo"
)

type TwitterClientMock struct {
	responses map[string]*twittergo.APIResponse
}

func (twitterClient *TwitterClientMock) SendRequest(req *http.Request) (resp *twittergo.APIResponse, err error) {
	resp, exists := twitterClient.responses[req.URL.String()]
	if !exists {
		fmt.Printf("No response for request %v with URI %v\n", req, req.URL.String())
	}

	return
}

func TestGetTweets(t *testing.T) {
	var tweetsCount uint = 10
	expectedReqURL := fmt.Sprintf("/1.1/search/tweets.json?count=%d&lang=en&q=fake+name&result_type=mixed",
		tweetsCount)
	expectedTweets := []twittergo.Tweet{
		twittergo.Tweet{
			"id_str": "716102930609209345",
			"text":   "this is a test tweet",
		},
		twittergo.Tweet{
			"id_str": "711700393533640704",
			"text":   "this is another test tweet",
		},
	}
	expectedTweetsJSON, _ := json.Marshal(&twittergo.SearchResults{"statuses": expectedTweets})
	expectedResp := &twittergo.APIResponse{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader(expectedTweetsJSON)),
	}
	twitterClient := TwitterClientMock{map[string]*twittergo.APIResponse{expectedReqURL: expectedResp}}

	twitterFacade := NewTwitterFacadeWithClient(&twitterClient)
	tweets, err := twitterFacade.GetTweets("fake name", tweetsCount)

	if err != nil {
		t.Fatalf("Expected nil error, but was %v", err)
	}
	if !reflect.DeepEqual(tweets, expectedTweets) {
		t.Errorf("Expected tweets are %v, but were %v", expectedTweets, tweets)
	}
}

func TestGetTweetsWithDifferentCounts(t *testing.T) {
	tweets := []twittergo.Tweet{
		twittergo.Tweet{
			"id_str": "716102930609209345",
			"text":   "this is a test tweet",
		},
		twittergo.Tweet{
			"id_str": "711700393533640704",
			"text":   "this is another test tweet",
		},
		twittergo.Tweet{
			"id_str": "711700393533640705",
			"text":   "this is yet another test tweet",
		},
	}
	counts := []uint{1, 2, 3}
	expectedResponses := make(map[string]*twittergo.APIResponse)
	expectedTweets := make(map[uint][]twittergo.Tweet)
	for _, count := range counts {
		expectedTweets[count] = tweets[:count]
		expectedReqURL := fmt.Sprintf("/1.1/search/tweets.json?count=%d&lang=en&q=fake+name&result_type=mixed",
			count)
		expectedTweetsJSON, _ := json.Marshal(&twittergo.SearchResults{"statuses": expectedTweets[count]})
		expectedResp := &twittergo.APIResponse{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(expectedTweetsJSON)),
		}
		expectedResponses[expectedReqURL] = expectedResp
	}
	twitterClient := TwitterClientMock{expectedResponses}

	twitterFacade := NewTwitterFacadeWithClient(&twitterClient)
	for _, count := range counts {
		tweets, err := twitterFacade.GetTweets("fake name", count)

		if err != nil {
			t.Fatalf("Expected nil error, but was %v", err)
		}
		if !reflect.DeepEqual(tweets, expectedTweets[count]) {
			t.Errorf("Expected tweets are %v, but were %v", expectedTweets[count], tweets)
		}
	}

}
