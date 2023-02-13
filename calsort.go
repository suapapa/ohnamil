package main

import "github.kakaoenterprise.in/cassian-l/go_kakaowork-bot/calendar"

type calItems []*calendar.Item

func (s calItems) Len() int {
	return len(s)
}

func (s calItems) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func (s calItems) Less(i, j int) bool {
	iStart := s[i].Start.DateTime.Unix()
	jStart := s[j].Start.DateTime.Unix()
	return iStart < jStart
}
