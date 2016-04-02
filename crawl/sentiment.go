/*
Package for crawling for tweets, getting their sentiment score, and writing them
to the database.
*/
package crawl

// The interface for retrieving the sentiment score of a tweet.
type SentimentAnalyzer interface {
	GetScoreForTweet(tweet string) int32
}

// Provide a mock SentimentAnalyzer for unit tests of files that depend on a SentimentAnalyzer.
type SentimentAnalyzerMock struct {
	Scores map[string]int32
}

func (analyzer SentimentAnalyzerMock) GetScoreForTweet(tweet string) int32 {
	return analyzer.Scores[tweet]
}
