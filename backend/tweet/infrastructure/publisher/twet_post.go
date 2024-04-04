package publisher

import (
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
	"net/http"
	"os"
)

type TweetPostPublisher struct{}

func (t TweetPostPublisher) Send(userID value.UserID, tweetID value.TweetID) {
	req, err := http.NewRequest("POST", os.Getenv("FANOUT_SERVER")+"/v1/users/"+userID.ToString()+"/tweets/"+tweetID.ToString(), nil)
	if err != nil {
		panic(err.Error())
	}
	httpClient := new(http.Client)
	_, _ = httpClient.Do(req)
}
