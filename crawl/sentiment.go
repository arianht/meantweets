/*
Package for crawling for tweets, getting their sentiment score, and writing them
to the database.
*/
package crawl

// The interface for retrieving the sentiment score of a tweet.
type SentimentAnalyzer interface {
	GetScoreForTweet(tweet string) int
}

// Provide a mock Dao for unit tests of files that depend on a Dao.
type sentimentAnalyzerMock struct {
	scores map[string]int
}

func (analyzer sentimentAnalyzerMock) GetScoreForTweet(tweet string) int {
	return analyzer.scores[tweet]
}
