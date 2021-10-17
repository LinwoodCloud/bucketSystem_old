package main

var (
	EndpointGitHub = "https://api.github.com"
	UsersEndpoint  = EndpointGitHub + "/users"
	UserEndpoint   = func(uId string) string { return UsersEndpoint + "/" + uId }
)
