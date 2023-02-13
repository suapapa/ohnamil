package main

import (
	"time"

	"github.com/pkg/errors"
	"github.kakaoenterprise.in/cassian-l/go_kakaowork-bot/calendar"
)

var (
	kepCalBot *calendar.Client
)

func getKepEvents(kepID string, from, to time.Time) ([]*calendar.Item, error) {
	if kepCalBot == nil {
		kepCalBot = calendar.NewClient(kepID)
	}

	resp, err := kepCalBot.PersonalEvents(from, to, true, "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get kep personal events")
	}

	return resp.Events.Items, nil
}
