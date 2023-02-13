package main

import (
	"flag"
	"log"
	"time"

	"github.com/fogleman/gg"
)

var (
	flagDryrun   bool
	flagInterval string
	// flagKepID    string
)

func main() {
	flag.BoolVar(&flagDryrun, "n", false, "dont disp. just save png")
	flag.StringVar(&flagInterval, "i", "15m", "display update interval")
	// flag.StringVar(&flagKepID, "kep", "", "kep ldap id")
	flag.Parse()

	initHW()

	kepID := flag.Arg(0)
	showTodayEvents(kepID)
}

func showTodayEvents(kepID string) {
	dur, err := time.ParseDuration(flagInterval)
	if err != nil {
		log.Fatal(err)
	}
	tkr := time.NewTicker(dur)

	now := time.Now()
	for {
		end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
		log.Printf("update display for now: %v", now)
		// 오늘 남은 일정 조회
		evtItems, err := getKepEvents(kepID, now, end)
		if err != nil {
			log.Fatal(err)
		}

		dc := gg.NewContext(dispW, dispH)
		drawDisp(dc, kepID, now, evtItems)

		now = <-tkr.C
	}
}
