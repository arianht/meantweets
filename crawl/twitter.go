/*
Package for crawling for tweets, getting their sentiment score, and writing them
to the database.
*/
package crawl

type TwitterFacade interface {
	GetTweets(celebrity string) []string
}

type twitterFacadeMock struct {
	tweets map[string][]string
}

func (twitter twitterFacadeMock) GetTweets(celebrity string) []string {
	return twitter.tweets[celebrity]
}
