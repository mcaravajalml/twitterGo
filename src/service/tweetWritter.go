package service

import (
	"github.com/mcaravajalml/twitterGo/src/domain"
)

type TweetWriter interface {
	Save(domain.Tweet)
	GetLastSavedTweet() domain.Tweet
}

type MemoryTweetWriter struct {
}

var TweetsBuffer []domain.Tweet = make([]domain.Tweet, 0)

func (d *MemoryTweetWriter) Save(obj domain.Tweet) {

	TweetsBuffer = append(TweetsBuffer, obj)
}

func (d *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	last := len(TweetsBuffer)
	return TweetsBuffer[last-1]
}
