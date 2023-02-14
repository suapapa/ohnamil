package main

import (
	"flag"
	"log"
	"time"

	"github.com/fogleman/gg"
)

var (
	flagDryrun      bool
	flagInterval    string
	flagGoogleCalID string
	flagKepID       string
	flagDisplayName string
)

func main() {
	flag.BoolVar(&flagDryrun, "n", false, "dont disp. just save png")
	flag.StringVar(&flagInterval, "i", "15m", "display update interval")
	flag.StringVar(&flagGoogleCalID, "gcal", "primary", "google calendar id")
	flag.StringVar(&flagKepID, "kep", "", "kep ldap id")
	flag.StringVar(&flagDisplayName, "d", "", "display user name")
	flag.Parse()

	initHW()

	showTodayEvents()
}

func showTodayEvents() {
	dur, err := time.ParseDuration(flagInterval)
	if err != nil {
		log.Fatal(err)
	}
	tkr := time.NewTicker(dur)

	now := time.Now()
	var items CalItems
	for {
		end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
		log.Printf("update display for now: %v", now)

		// 오늘 남은 일정 조회
		kepItems, err := getKepCalItems(flagKepID, now, end)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, kepItems...)
		googleItems, err := getGoogleCalItems(now, end)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, googleItems...)

		dc := gg.NewContext(dispW, dispH)
		drawDisp(dc, flagDisplayName, now, items)

		now = <-tkr.C
	}
}
