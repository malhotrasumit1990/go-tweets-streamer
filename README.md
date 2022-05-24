<h1>Go-Tweets-Streamer Application</h1>

<h2>This MVP is responsible to stream the real time tweets. To test the Streamer Application : </h2>

1. Add Proper credentials (API_KEY, API_KEY_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET) in the environment variables.
2. Inside `go-tweets-streamer` folder run command `go run main.go`
3. This will start the streamer application, and will start looking for tweets matching keywords mentioned in `Track` attribute.

<h2>Other functionalities:</h2>

<h3>Post a tweet</h3>
        1. To Post a tweet: Uncomment line number 52 in main.go and rerun the application.

        2. This will post `Test Tweet`

<h3>Search tweets</h3>
        1. To Search for tweets: Uncomment line number 53 in main.go and rerun the application.

        2. This will search tweets with  `Test` Keyword in it.   
        
<h3> Unit Tests </h3>

        I did not add unit tests here because It would just have been mocking the twitter client and testing mocked request and responses.

        If It was like, processing the response and then doing some thing with it. I would have added unit tests as well.
