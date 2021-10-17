package main

import "bucketSystem/github"

type Bucket struct {
	Slug        string
	Name        map[string]string
	Description map[string]string
	Owner       []*github.User
	Maintainer  []*github.User
}
