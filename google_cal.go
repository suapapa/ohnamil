package main

import (
	"context"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var (
	googleCalSvc *calendar.Service
)

func getGoogleCalItems(from, to time.Time) (CalItems, error) {
	if googleCalSvc == nil {
		svc, err := NewGoogleCalSvc(context.Background())
		if err != nil {
			return nil, err
		}
		googleCalSvc = svc
	}

	minT := from.Format(time.RFC3339)
	maxT := to.Format(time.RFC3339)
	events, err := googleCalSvc.Events.List(flagGoogleCalID).ShowDeleted(false).
		SingleEvents(true).TimeMin(minT).TimeMax(maxT).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		return nil, err
	}

	ret := make(CalItems, len(events.Items))
	for i, item := range events.Items {
		ret[i] = NewCalItemFromGoogleCalItem(item)
	}
	return ret, nil
}

func NewGoogleCalSvc(ctx context.Context) (*calendar.Service, error) {
	c, err := google.DefaultClient(ctx, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, err
	}

	return calendar.NewService(ctx, option.WithHTTPClient(c))
}
