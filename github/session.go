package github

import "net/http"

type Session struct {
	Token int
}

func (s Session) User(uId string) (User, error) {
	result, err := http.Get(UserEndpoint(uId))
	if err != nil {
		return User{}, err
	}
	if result.StatusCode != 200 {

	}
	return User{}, nil
}
