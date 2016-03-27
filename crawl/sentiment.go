package crawl

type AnalyzeSentiment interface {
	GetScoreForTweet(tweet string) int
}
