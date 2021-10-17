package github

import "time"

type Label struct {
	Id          int64  `json:"id"`
	NodeId      string `json:"node_id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
	Description string `json:"description"`
}

type MessageReaction struct {
	Url        string `json:"url"`
	TotalCount int    `json:"total_count"`
	Field3     int    `json:"+1"`
	Field4     int    `json:"-1"`
	Laugh      int    `json:"laugh"`
	Hooray     int    `json:"hooray"`
	Confused   int    `json:"confused"`
	Heart      int    `json:"heart"`
	Rocket     int    `json:"rocket"`
	Eyes       int    `json:"eyes"`
}

type Issue struct {
	Url                   string          `json:"url"`
	RepositoryUrl         string          `json:"repository_url"`
	LabelsUrl             string          `json:"labels_url"`
	CommentsUrl           string          `json:"comments_url"`
	EventsUrl             string          `json:"events_url"`
	HtmlUrl               string          `json:"html_url"`
	Id                    int             `json:"id"`
	NodeId                string          `json:"node_id"`
	Number                int             `json:"number"`
	Title                 string          `json:"title"`
	User                  *SmallUser      `json:"user"`
	Labels                []*Label        `json:"labels"`
	State                 string          `json:"state"`
	Locked                bool            `json:"locked"`
	Assignee              *SmallUser      `json:"assignee"`
	Assignees             []*SmallUser    `json:"assignees"`
	Milestone             interface{}     `json:"milestone"`
	Comments              int             `json:"comments"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`
	ClosedAt              interface{}     `json:"closed_at"`
	AuthorAssociation     string          `json:"author_association"`
	ActiveLockReason      interface{}     `json:"active_lock_reason"`
	Body                  string          `json:"body"`
	ClosedBy              interface{}     `json:"closed_by"`
	Reactions             MessageReaction `json:"reactions"`
	TimelineUrl           string          `json:"timeline_url"`
	PerformedViaGithubApp interface{}     `json:"performed_via_github_app"`
}
