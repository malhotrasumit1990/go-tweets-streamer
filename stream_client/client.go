package stream_client

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/go-tweets-streamer/model"
)

// This function return a twitter client built on top of HTTP client to make requests with Twitter API
func Create_Streamer_Client(creds *model.Credentials) *twitter.Client {

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	return twitter.NewClient(httpClient)

}

// Get_Demux is to deMultiplex the message into a tweet or Direct message.
func Get_Demux() twitter.SwitchDemux {

	demux := twitter.NewSwitchDemux()

	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	return demux

}

// Start_Streamer will start Stream , look for  tweets where it finds a string "Dhoni"
func Start_Streamer(twitter_client *twitter.Client) (*twitter.Stream, error) {

	params := &twitter.StreamFilterParams{
		Track:         []string{"Dhoni"},
		StallWarnings: twitter.Bool(true),
	}
	return twitter_client.Streams.Filter(params)

}

//Post_Tweet to post a tweet
func Post_Tweet(twitter_client *twitter.Client, tweet_msg string) {

	tweet, resp, err := twitter_client.Statuses.Update(tweet_msg, nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)

}

// Search_Tweet will Search for a Tweet using a query string
func Search_Tweet(twitter_client *twitter.Client, search_string string) {

	search, resp, err := twitter_client.Search.Tweets(&twitter.SearchTweetParams{
		Query: search_string,
	})

	if err != nil {
		log.Print(err)
	}

	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", search)

}
