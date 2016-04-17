// Package crawl provides a means for crawling for tweets, getting their sentiment score, and writing them
// to the database.
package crawl

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/arianht/meantweets/util"
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
)

// TwitterFacade is an interface for retrieving celebrity tweets from Twitter.
type TwitterFacade interface {
	GetTweets(celebrity string, count uint) ([]twittergo.Tweet, error)
}

type twitterFacade struct {
	client TwitterClient
}

// TwitterClient is used to write unit tests for code based on the twittergo package.
type TwitterClient interface {
	SendRequest(req *http.Request) (resp *twittergo.APIResponse, err error)
}

func getTwitterClient(apiCredentialsFilename string) (client TwitterClient, err error) {
	credentials, err := util.GetTwitterAPICredentialsFromFile(apiCredentialsFilename)
	if err != nil {
		return
	}
	config := &oauth1a.ClientConfig{
		ConsumerKey:    credentials.ConsumerKey,
		ConsumerSecret: credentials.ConsumerSecret,
	}
	client = twittergo.NewClient(config, nil)
	return
}

func createSearchTweetsRequest(searchQuery string, count uint) (req *http.Request, err error) {
	url := fmt.Sprintf("/1.1/search/tweets.json?%v", getURLValues(searchQuery, count).Encode())
	req, err = http.NewRequest("GET", url, nil)
	return
}

func getURLValues(searchQuery string, count uint) (query *url.Values) {
	query = &url.Values{}
	query.Set("q", searchQuery+" -filter:retweets -filter:media -filter:replies -filter:links")
	query.Set("lang", "en")
	query.Set("count", fmt.Sprintf("%d", count))
	query.Set("result_type", "mixed") // Include both popular and real time results in the response
	return
}

// Searches Twitter for and returns tweets that contain the given string.
// The result will contain up to count tweets. As of 04/06/2016, Twitter will
// not return more than 100 results per request.
func (twitter *twitterFacade) GetTweets(celebrity string, count uint) (tweets []twittergo.Tweet, err error) {
	req, err := createSearchTweetsRequest(celebrity, count)
	if err != nil {
		return
	}

	resp, err := twitter.client.SendRequest(req)
	if err != nil {
		return
	}

	results := &twittergo.SearchResults{}
	err = resp.Parse(results)
	if err != nil {
		return
	}

	tweets = results.Statuses()
	return
}

// NewTwitterFacade returns a Twitter facade that can be used to get tweets based on a query string.
// The Twitter facade reads API credentials from the given JSON file which must have the following format:
// {
//   "ConsumerKey": "<consumer-key>",
//   "ConsumerSecret": "<consumer-secret>"
// }
func NewTwitterFacade(apiCredentialsFilename string) (twitter TwitterFacade, err error) {
	facade := &twitterFacade{}
	client, err := getTwitterClient(apiCredentialsFilename)
	if err != nil {
		return
	}
	facade.client = client
	twitter = facade
	return
}

// NewTwitterFacadeWithClient returns a Twitter facade with a custom client that can be used for unit testing.
func NewTwitterFacadeWithClient(client TwitterClient) TwitterFacade {
	return &twitterFacade{client}
}

// TwitterFacadeMock provides a mock TwitterFacade for unit tests of files that depend on a TwitterFacade.
type TwitterFacadeMock struct {
	Tweets map[string][]twittergo.Tweet
}

// GetTweets implementation for the TwitterFacade mock.
func (twitter TwitterFacadeMock) GetTweets(celebrity string, count uint) ([]twittergo.Tweet, error) {
	return twitter.Tweets[celebrity], nil
}
