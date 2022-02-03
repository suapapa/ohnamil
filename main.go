package main

import (
	"flag"
	"log"
	"time"

	"github.com/fogleman/gg"
	"github.kakaoenterprise.in/IoTEngine/go_kakaowork-bot/calendar"
)

var (
	flagDryrun   bool
	flagInterval string
)

func main() {
	flag.BoolVar(&flagDryrun, "n", false, "dont disp. just save png")
	flag.StringVar(&flagInterval, "i", "15m", "display update interval")
	flag.Parse()

	initHW()

	userID := flag.Arg(0)
	showTodayEvents(userID)
}

func showTodayEvents(userID string) {
	calBol := calendar.NewClient(userID)

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
		resp, err := calBol.PersonalEvents(now, end, true, "")
		if err != nil {
			log.Fatal(err)
		}

		dc := gg.NewContext(dispW, dispH)
		drawDisp(dc, resp.Nickname, now, resp.Events)

		now = <-tkr.C
	}
}
