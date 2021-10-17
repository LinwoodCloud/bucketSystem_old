package main

type Bucket struct {
	Slug        string
	Name        map[string]string
	Description map[string]string
	Owner       []*User
	Maintainer  []*User
}
