package main

import (
	"embed"
	"fmt"
	"image/color"
	"sort"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/lestrrat-go/dither"
	"github.kakaoenterprise.in/IoTEngine/go_kakaowork-bot/calendar"
	"golang.org/x/image/font"
)

const (
	dispW   = 800
	dispH   = 480
	display = "display.png"

	fsH1 = 40
	fsH2 = 30
	fsH3 = 20
)

var (
	//go:embed assets/RixPoetry.ttf
	resource embed.FS

	ttf *truetype.Font
	h   = 0.0
)

func init() {
	var err error
	fontdata, err := resource.ReadFile("assets/RixPoetry.ttf")
	if err != nil {
		panic(err)
	}
	ttf, err = truetype.Parse(fontdata)
	if err != nil {
		panic(err)
	}

}

var (
	lastDispItems []Event
)

func drawDisp(dc *gg.Context, nick string, now time.Time, events calendar.Events) {
	dc.SetColor(color.White)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	h = 10.0

	drawString(dc, fmt.Sprintf("%s 의 오늘 남은 일정", nick), fsH2, 20, h+fsH2+10)
	h += fsH2 + 10
	h += 5

	items := calItems(events.Items)
	sort.Sort(items)

	dispItems := make([]Event, len(events.Items))
	for i, item := range events.Items {
		dispItems[i].FromT = item.Start.DateTime.Format("15:04")
		dispItems[i].ToT = item.End.DateTime.Format("15:04")
		dispItems[i].Desc = item.Summary
	}

	if len(events.Items) == 0 {
		drawStringAnchoredCenter(dc, "없음", 100, dispW/2, dispH/2)
	} else {
		for _, item := range dispItems {
			str := stripStr(fmt.Sprintf("- %s ~ %s", item.FromT, item.ToT))
			drawString(dc, str, fsH1, 20, h+fsH1+10)
			h += fsH1 + 10
			str = stripStr(fmt.Sprintf("    %v", item.Desc))
			drawString(dc, str, fsH1, 20, h+fsH1+10)
			h += fsH1 + 10

			if h > dispH {
				break
			}
		}
	}

	// draw footer
	_, ip, _, _ := resolveNet()
	drawStringAnchoredBR(dc, ip+"; "+now.Format("2006-01-02 15:04")+"; "+flagInterval, fsH3, dispW, dispH)

	if flagDryrun {
		dc.SavePNG(display)
	} else {
		if isEqualEvents(dispItems, lastDispItems) {
			return
		}

		img := dc.Image()
		ditheredImg := dither.Monochrome(dither.Burkes, img, 1.18)
		updatePanel(ditheredImg)
		updatePanel(ditheredImg)

		lastDispItems = dispItems
	}
}

// ----

func drawString(dc *gg.Context, text string, fontSize, x, y float64) {
	ff, err := loadFontFace(fontSize)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(ff)
	dc.DrawString(text, x, y)
}

func drawStringAnchoredBR(dc *gg.Context, text string, fontSize, x, y float64) {
	ff, err := loadFontFace(fontSize)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(ff)
	dc.DrawStringAnchored(text, x-20, y, 1.0, -0.5)
}

func drawStringAnchoredCenter(dc *gg.Context, text string, fontSize, x, y float64) {
	ff, err := loadFontFace(fontSize)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(ff)
	dc.DrawStringAnchored(text, x-20, y, 0.5, 0.5)
}

func loadFontFace(points float64) (font.Face, error) {
	ff := truetype.NewFace(ttf, &truetype.Options{
		Size:    points,
		Hinting: font.HintingFull,
		// Hinting: font.HintingNone,
	})
	return ff, nil
}
