package main

import (
	"embed"
	"fmt"
	"image/color"
	"sort"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.kakaocorp.com/DeviceDev/go_kakaowork-bot/calendar"
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

func drawDisp(dc *gg.Context, nick string, now time.Time, events calendar.Events) {
	dc.SetColor(color.White)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	h = 10.0

	drawString(dc, fmt.Sprintf("%s 의 오늘일정", nick), fsH2, 20, h+fsH2+10)
	h += fsH2 + 10
	h += 5

	items := calItems(events.Items)
	sort.Sort(items)

	for _, item := range events.Items {
		start := item.Start.DateTime
		end := item.End.DateTime

		str := stripStr(fmt.Sprintf("- %s ~ %s", start.Format("15:04"), end.Format("15:04")))
		drawString(dc, str, fsH1, 20, h+fsH1+10)
		h += fsH1 + 10
		str = stripStr(fmt.Sprintf("    %v", item.Summary))
		drawString(dc, str, fsH1, 20, h+fsH1+10)
		h += fsH1 + 10

		if h > dispH {
			break
		}

		// // str = stripStr(fmt.Sprintf("  %s ~ %s", start.Format("20060102-15:04"), end.Format("20060102-15:04")))
		// drawString(dc, str, fsH3, 10, h+fsH3+10)
		// h += fsH3 + 10
	}

	// draw footer
	_, ip, _, _ := resolveNet()
	drawStringAnchoredBR(dc, ip+"; "+now.Format("2006-01-02 15:04"), fsH3, dispW, dispH)

	if flagDryrun {
		dc.SavePNG(display)
	} else {
		updatePanel(dc.Image())
		updatePanel(dc.Image())
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

func loadFontFace(points float64) (font.Face, error) {
	ff := truetype.NewFace(ttf, &truetype.Options{
		Size:    points,
		Hinting: font.HintingFull,
		// Hinting: font.HintingNone,
	})
	return ff, nil
}
