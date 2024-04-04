package publisher

import (
	"github.com/mastar3104/twitter-clone/user/domain/value"
	"net/http"
	"os"
)

type UserFollowPublisher struct{}

func (t UserFollowPublisher) Send(userID value.UserID, followUserID value.UserID) {
	req, err := http.NewRequest("POST", os.Getenv("FANOUT_SERVER")+"/v1/users/"+userID.ToString()+"/follow/"+followUserID.ToString(), nil)
	if err != nil {
		panic(err.Error())
	}
	httpClient := new(http.Client)
	_, _ = httpClient.Do(req)
}
