Go-Tweets-Streamer Application

This MVP is responsible to stream the real time tweets 

To test the Streamer Application : 
    1. Add Proper credentials (API_KEY, API_KEY_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET) in the environment variables.
    2. Inside `go-tweets-streamer` folder run command `go run main.go`
    3. This will start the streamer application, and will start looking for tweets matching keywords mentioned in `Track` attribute.

Other functionalities:
    Post a tweet.
        1. To Post a tweet: Uncomment line number 52 in main.go and rerun the application.
        2. This will post "Test Tweet"

    Search Tweet
        1. To Search for tweets: Uncomment line number 53 in main.go and rerun the application.
        2. This will search tweets with  "Test" Keyword in it.    