/*
Package for crawling for tweets, getting their sentiment score, and writing them
to the database.
*/
package crawl

type TwitterFacade interface {
	GetTweets(celebrity string) []string
}

// Provide a mock TwitterFacade for unit tests of files that depend on a TwitterFacade.
type TwitterFacadeMock struct {
	Tweets map[string][]string
}

func (twitter TwitterFacadeMock) GetTweets(celebrity string) []string {
	return twitter.Tweets[celebrity]
}
