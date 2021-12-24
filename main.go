package main

import (
	"flag"
	"log"
	"time"

	"github.com/fogleman/gg"
	"github.kakaocorp.com/DeviceDev/go_kakaowork-bot/calendar"
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

	for {
		now := time.Now()
		resp, err := calBol.PersonalEvents(now, now.Add(24*time.Hour), true, "")
		if err != nil {
			log.Fatal(err)
		}

		dc := gg.NewContext(dispW, dispH)
		drawDisp(dc, resp.Nickname, now, resp.Events)

		<-tkr.C
	}
	// // jEnc := json.NewEncoder(os.Stdout)
	// // jEnc.SetIndent("", "  ")
	// for _, e := range resp.Events.Items {
	// 	// jEnc.Encode(e)
	// 	fmt.Printf("%s - %v ~ %v", e.Summary, e.Start.DateTime, e.End.DateTime)
	// }
}
