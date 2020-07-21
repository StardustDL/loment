package models

import (
	"time"
)

type Comment struct {
	Id               string
	CreationTime     time.Time
	ModificationTime time.Time
	Content          string
	Uri              string
	Author           string
	Email            string
	Link             string
	Extra            string
}

type CommentQuery struct {
	Id               string
	CreationTime     time.Time
	ModificationTime time.Time
	Content          string
	Uri              string
	Author           string
	Email            string
	Link             string
	Offset           int
	Limit            int
}
