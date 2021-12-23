package main

import (
	"flag"
	"image/color"
	"log"
	"time"

	"github.com/fogleman/gg"
	"github.kakaocorp.com/DeviceDev/go_kakaowork-bot/calendar"
)

var (
	flagDryrun bool
)

func main() {
	flag.BoolVar(&flagDryrun, "n", false, "dont disp. just save png")
	flag.Parse()

	initHW()

	userID := flag.Arg(0)
	showTodayEvents(userID)
}

func showTodayEvents(userID string) {
	calBol := calendar.NewClient(userID)

	tkr := time.NewTicker(15 * time.Minute)

	for {
		now := time.Now()
		resp, err := calBol.PersonalEvents(now, now.Add(24*time.Hour), true, "")
		if err != nil {
			log.Fatal(err)
		}

		dc := gg.NewContext(dispW, dispH)
		dc.SetColor(color.White)
		dc.Clear()
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
