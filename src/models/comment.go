package models

import (
	"time"
)

type Comment struct {
	ID               string
	CreationTime     time.Time
	ModificationTime time.Time
	Location         string
	Nick             string
	Mail             string
	Link             string
}
