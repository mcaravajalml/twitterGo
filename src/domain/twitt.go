package domain

import (
	"fmt"
	"time"
)

type TextTweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
}

type Stringer interface {
	String() string
}

type Tweet interface {
	PrintableTweet() string
	String() string
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetId() int
	SetId(int)
	SetText(string)
	SetUser(string)
}

var tweet *TextTweet

func NewTweet(user, text string) Tweet {

	date := time.Now()
	tweet = new(TextTweet)

	tweet.Date = &date
	tweet.User = user
	tweet.Text = text

	return tweet
}

func (tweet *TextTweet) PrintableTweet() string {

	print("%v", tweet.Text)
	return tweet.Text
}

func (tweet *TextTweet) String() string {
	aux := fmt.Sprintf("@" + tweet.User + ": " + tweet.Text)
	return aux
}

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) SetText(text string) {
	tweet.Text = text
}

func (tweet *TextTweet) SetUser(user string) {
	tweet.User = user
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}
