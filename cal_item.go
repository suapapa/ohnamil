package main

import (
	"time"

	kep_cal "github.kakaoenterprise.in/cassian-l/go_kakaowork-bot/calendar"
	google_cal "google.golang.org/api/calendar/v3"
)

type CalItem struct {
	Start   time.Time
	End     time.Time
	Summary string
}

func NewCalItemFromKepCalItem(item *kep_cal.Item) *CalItem {
	return &CalItem{
		Start:   item.Start.DateTime,
		End:     item.End.DateTime,
		Summary: item.Summary,
	}
}

func NewCalItemFromGoogleCalItem(item *google_cal.Event) *CalItem {
	start, _ := time.Parse(time.RFC3339, item.Start.DateTime)
	end, _ := time.Parse(time.RFC3339, item.End.DateTime)
	return &CalItem{
		Start:   start,
		End:     end,
		Summary: item.Summary,
	}
}

type CalItems []*CalItem

func (s CalItems) Len() int {
	return len(s)
}

func (s CalItems) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func (s CalItems) Less(i, j int) bool {
	iStart := s[i].Start.Unix()
	jStart := s[j].End.Unix()
	return iStart < jStart
}
