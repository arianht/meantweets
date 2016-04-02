/*
Package for crawling for tweets, getting their sentiment score, and writing them
to the database.
*/
package crawl

// A tweet from Twitter.
type Tweet struct {
	Id      int64
	Content string
}

// The interface for retrieving celebrity tweets from Twitter.
type TwitterFacade interface {
	GetTweets(celebrity string) []Tweet
}

// Provide a mock TwitterFacade for unit tests of files that depend on a TwitterFacade.
type TwitterFacadeMock struct {
	Tweets map[string][]Tweet
}

func (twitter TwitterFacadeMock) GetTweets(celebrity string) []Tweet {
	return twitter.Tweets[celebrity]
}
