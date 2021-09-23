package model

import "time"

type ListInfo struct {
	UserID      string
	ListTitle   string
	ListContext string
	StartTime   time.Time
	EndTime     time.Time
	TimeUp      bool
}

func NewListInfo(id string, title string, context string, start time.Time, end time.Time, timeup bool) *ListInfo {
	return &ListInfo{
		UserID:      id,
		ListTitle:   title,
		ListContext: context,
		StartTime:   start,
		EndTime:     end,
		TimeUp:      timeup,
	}
}
