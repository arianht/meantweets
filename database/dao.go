/*
Package database provides a means for reading and writing from the database. Perform all database
operations through this package.
*/
package database

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const (
	datastoreKind = "Tweet"
)

// Dao defines the interface for the data access object that abstracts database interactions.
type Dao interface {
	WriteCelebrityTweets(tweets []Tweet) (err error)
	GetCelebrityTweets(celebrityName string) (tweets []Tweet, err error)
	DeleteAllTweetsForCelebrity(celebrityName string) (err error)
}

// DatastoreDao is a DAO for interacting with App Engine's Datastore.
type DatastoreDao struct {
	Ctx context.Context // App Engine Context which can be obtained from an HTTP request.
}

// Tweet is an entity for storing data in the datastore.
type Tweet struct {
	CelebrityName string `json:",omitempty"`
	Id            int64
	Score         int32
}

// WriteCelebrityTweets saves the slice of tweets to the database. Note that duplicates aren't
// caught here because of writing asynchronicity.
func (datastoreDao DatastoreDao) WriteCelebrityTweets(tweets []Tweet) (err error) {
	keys := make([]*datastore.Key, len(tweets))
	for i, _ := range tweets {
		keys[i] = datastore.NewIncompleteKey(datastoreDao.Ctx, datastoreKind, nil)
	}
	_, err = datastore.PutMulti(datastoreDao.Ctx, keys, tweets)
	return
}

// GetCelebrityTweets retrieves all the celebrity tweets related to a celebrity sorted with highest
// scores first. Duplicate tweets will be filtered out.
func (datastoreDao DatastoreDao) GetCelebrityTweets(celebrityName string) (tweets []Tweet,
	err error) {
	q := datastore.NewQuery(datastoreKind).
		Filter("CelebrityName = ", celebrityName).
		Project("Id", "Score").
		Order("Score")
	var results []Tweet
	if _, err = q.GetAll(datastoreDao.Ctx, &results); err != nil {
		return
	}
	tweetSet := make(map[int64]bool)
	for _, tweetResult := range results {
		if !tweetSet[tweetResult.Id] {
			tweets = append(tweets, tweetResult)
			tweetSet[tweetResult.Id] = true
		}
	}
	return
}

// DeleteAllTweetsForCelebrity deletes all tweets for a provided celebirty name.
func (datastoreDao DatastoreDao) DeleteAllTweetsForCelebrity(celebrityName string) (err error) {
	q := datastore.NewQuery(datastoreKind).
		Filter("CelebrityName = ", celebrityName).
		KeysOnly()
	keys, err := q.GetAll(datastoreDao.Ctx, nil)
	if err != nil {
		return
	}
	err = datastore.DeleteMulti(datastoreDao.Ctx, keys)
	return
}

// DaoMock provides a mock Dao for unit tests of files that depend on a Dao.
type DaoMock struct {
	Tweets *[]Tweet // Use a pointer so all copies of DaoMock modify the same "database".
}

// WriteCelebrityTweet implementation for DaoMock.
func (dao DaoMock) WriteCelebrityTweets(tweets []Tweet) (err error) {
	*dao.Tweets = append(*dao.Tweets, tweets...)
	return
}

// GetCelebrityTweets implementation for DaoMock.
func (dao DaoMock) GetCelebrityTweets(celebrityName string) (tweets []Tweet, err error) {
	for _, tweet := range *dao.Tweets {
		if tweet.CelebrityName == celebrityName {
			tweets = append(tweets, tweet)
		}
	}
	return
}

// DeleteAllTweetsForCelebrity implementation for DaoMock.
func (dao DaoMock) DeleteAllTweetsForCelebrity(celebrityName string) (err error) {
	newTweets := []Tweet{}
	for _, tweet := range *dao.Tweets {
		if tweet.CelebrityName != celebrityName {
			newTweets = append(newTweets, tweet)
		}
	}
	*dao.Tweets = newTweets
	return
}
