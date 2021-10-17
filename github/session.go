package github

type Session struct {
	Token int
}

func (s Session) User() User {
	return User{}
}
