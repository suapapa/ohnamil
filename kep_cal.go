package main

import (
	"time"

	"github.com/pkg/errors"
	"github.kakaoenterprise.in/gopher/go_kakaowork-bot/calendar"
)

var (
	kepCalBot *calendar.Client
)

func getKepCalItems(kepID string, from, to time.Time) (CalItems, error) {
	if kepCalBot == nil {
		kepCalBot = calendar.NewClient(kepID)
	}

	resp, err := kepCalBot.PersonalEvents(from, to, true, "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get kep personal events")
	}

	ret := make(CalItems, len(resp.Events.Items))
	for i, item := range resp.Events.Items {
		ret[i] = NewCalItemFromKepCalItem(item)
	}
	return ret, nil
}
