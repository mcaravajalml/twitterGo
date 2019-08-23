package service_test

import (
	"testing"

	"github.com/mcaravajalml/twitterGo/src/domain"
	"github.com/mcaravajalml/twitterGo/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var writer service.TweetWriter
	tweetManager := service.NewTweetManager(writer)

	var tweet domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	if publishedTweet.GetUser() != user &&
		publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.GetUser(), publishedTweet.GetText())
	}
	if publishedTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	var writer service.TweetWriter
	tweetManager := service.NewTweetManager(writer)
	var tweet domain.Tweet

	var user string
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err, _ = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization
	var writer service.TweetWriter
	tweetManager := service.NewTweetManager(writer)
	var tweet domain.Tweet

	user := "New user"
	text := ""
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err, _ = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWitchExceeding140CharacterIsNotPublished(t *testing.T) {
	// Initialization
	var writer service.TweetWriter
	tweetManager := service.NewTweetManager(writer)
	var tweet domain.Tweet

	user := "New user"
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas fringilla diam" +
		"nec mauris malesuada dictum. Integer malesuada purus a amet. "

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error

	err, _ = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text is very long" {
		t.Error("Expected error is user is required")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	var writer service.TweetWriter
	tweetManager := service.NewTweetManager(writer)

	var tweet, secondTweet domain.Tweet

	tweet = domain.NewTweet("New user", "Loren ipsum")
	secondTweet = domain.NewTweet("New user2", "Loren ipsum2")

	// Operation
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	if !isValidTweet(t, firstPublishedTweet, tweet.GetUser(), tweet.GetText()) {
		return
	}
	// Same for secondPublishedTweet

	if !isValidTweet(t, secondPublishedTweet, secondTweet.GetUser(), secondTweet.GetText()) {
		return
	}
}

func isValidTweet(t *testing.T, tweet domain.Tweet, user string, text string) bool {

	if tweet.GetUser() != user || tweet.GetText() != text {
		return false
	} else {
		return true
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	var writer service.TweetWriter
	tweetManager := service.NewTweetManager(writer)
	var tweet, secondTweet, thirdTweet domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	// publish the 3 tweets
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 { /* handle error */
		t.Error("Tienen que ser dos ")
	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet

	if !isValidTweet(t, firstPublishedTweet, tweet.GetUser(), tweet.GetText()) {
		return
	}
	// Same for secondPublishedTweet

	if !isValidTweet(t, secondPublishedTweet, secondTweet.GetUser(), secondTweet.GetText()) {
		return
	}
}

/*
	func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

    // Initialization
    tweet := domain.NewImageTweet("grupoesfera", "This is my image",
                "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
    // Operation
    text := tweet.PrintableTweet()
    // Validation
    expectedText := "@grupoesfera: This is my image
                    http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
    if text != expectedText {...}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
    // Initialization
    quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
    tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)
    // Validation
    expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
    if text != expectedText {...}
}
*/

func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {

	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet = new(domain.TextTweet)
	// Fill the tweet with data
	tweet.SetText("loren ipsum")
	tweet.SetUser("Test user")

	// Operation
	_, id := tweetManager.PublishTweet(tweet)

	// Validation
	memoryWriter := (tweetWriter).(*service.MemoryTweetWriter)
	savedTweet := memoryWriter.GetLastSavedTweet()

	if savedTweet == nil {
		t.Errorf("Tweet was not save")
	}

	if savedTweet.GetId() != id {
		t.Errorf("the id is not equals")
	}
}
