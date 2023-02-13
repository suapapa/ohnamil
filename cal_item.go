package main

import (
	"time"

	kep_cal "github.kakaoenterprise.in/cassian-l/go_kakaowork-bot/calendar"
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
