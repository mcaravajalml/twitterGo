package service

import (
	"fmt"

	"github.com/mcaravajalml/twitterGo/src/domain"
)

type TweetManager struct {
	Tweet    domain.Tweet
	Tweets   []domain.Tweet
	mapTweet map[string][]domain.Tweet
	writer   TweetWriter
}

var tweetManager *TweetManager

func NewMemoryTweetWriter() TweetWriter {

	return new(MemoryTweetWriter)
}

func NewTweetManager(writer TweetWriter) *TweetManager {

	tweetManager = new(TweetManager)
	tweetManager.mapTweet = make(map[string][]domain.Tweet)
	tweetManager.writer = writer

	return tweetManager
}

func (manager *TweetManager) GetTweetsByUser(user string) []domain.Tweet {

	return tweetManager.mapTweet[user]
}

func (manager *TweetManager) PublishTweet(tweet domain.Tweet) (error, int) {
	if tweet.GetUser() == "" {
		return fmt.Errorf("user is required"), 0
	}
	if tweet.GetUser() == "" {
		return fmt.Errorf("text is required"), 0
	}

	if lenght := len(tweet.GetText()); lenght > 140 {
		return fmt.Errorf("text is very long"), 0
	}

	tweetManager.Tweets = append(tweetManager.Tweets, tweet)

	if _, exist := tweetManager.mapTweet[tweet.GetUser()]; !exist {

		tweetManager.mapTweet[tweet.GetUser()] = make([]domain.Tweet, 0)
	}

	array := tweetManager.mapTweet[tweet.GetUser()]
	tweetManager.mapTweet[tweet.GetUser()] = append(array, tweet)

	tweetManager.writer.Save(tweet)

	return nil, 99
}

func GetTweet() domain.Tweet {
	return tweetManager.Tweet
}

func GetTweets() []domain.Tweet {
	return tweetManager.Tweets
}
