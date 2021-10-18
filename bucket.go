package main

import (
	"time"
)

type Bucket struct {
	IssueID                  int
	Slug                     string
	Name                     map[string]string
	Description              map[string]string
	Owner                    []*BucketUser
	Maintainer               []*BucketUser
	SourceNeeded             bool
	PermittedDownloadSources []string
	Items                    []*BucketItem
	Reviews                  []*BucketReview
}

type BucketItem struct {
	IssueID     int
	Slug        string
	Name        string
	Description string
	Source      string
	Tracks      []string
	Updates     []*BucketItemUpdate
	Created     time.Time
	Reviews     []*BucketReview
}

type BucketItemUpdate struct {
	Track        string
	Name         string
	Description  string
	IssueID      int
	Created      time.Time
	DownloadLink string
	Reviews      []*BucketReview
}

type BucketUser struct {
	ID int
}

type BucketReview struct {
	Author  *BucketUser
	Content string
}
