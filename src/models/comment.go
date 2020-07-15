package models

import (
	"time"
)

type Comment struct {
	ID               string
	CreationTime     time.Time
	ModificationTime time.Time
	Content          string
	Location         string
	Nick             string
	Mail             string
	Link             string
}

type CommentQuery struct {
	ID               string
	CreationTime     time.Time
	ModificationTime time.Time
	Content          string
	Location         string
	Nick             string
	Mail             string
	Link             string
	Offset           int
	Limit            int
}
