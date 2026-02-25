package main

import (
	"fmt"
	"time"

	"github.com/sav7ng/go-screws/array"
	"github.com/sav7ng/go-screws/calendar"
	"github.com/sav7ng/go-screws/geo"
	"github.com/sav7ng/go-screws/qrcode"
	"github.com/sav7ng/go-screws/random"
	gotime "github.com/sav7ng/go-screws/time"
	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("=== go-screws hello world demo ===")

	fmt.Println("\n--- RandomTool ---")
	rt := &random.RandomTool{}
	fmt.Println("Random chars (16):", rt.GetRandomChars(16))
	fmt.Println("Random int [1,100]:", rt.GetRandomInt(1, 100))
	fmt.Println("Random digits (8):", rt.GetRandomString(8))

	fmt.Println("\n--- ArrayTool ---")
	at := &array.ArrayTool{}
	merged := at.ArrayStringMerge([]string{"a", "b", "c"}, []string{"c", "d", "e"})
	fmt.Println("Merged [a,b,c]+[c,d,e]:", merged)
	deduped := at.ArrayStringUnique(merged)
	fmt.Println("Deduplicated:", deduped)
	fmt.Println("'c' in array?", at.StrInArray("c", merged))
	generic := array.RemoveRepeatElement([]int{1, 2, 2, 3, 3, 4})
	fmt.Println("Generic dedup [1,2,2,3,3,4]:", generic)

	fmt.Println("\n--- CalendarTool ---")
	ct := &calendar.CalendarTool{}
	now := time.Now()
	fmt.Println("First day of year:", ct.FirstDayOfYear(now))
	fmt.Println("Last day of year:", ct.LastDayOfYear(now))
	fmt.Println("First day of month:", ct.FirstDayOfMonth(now))
	fmt.Println("Last day of month:", ct.LastDayOfMonth(now))

	fmt.Println("\n--- TimeTool ---")
	tt := &gotime.TimeTool{}
	t1 := time.Now()
	t2 := t1.AddDate(0, 0, -10)
	fmt.Println("SubDays (10 days apart):", tt.SubDays(t1, t2))
	fmt.Println("Format time:", tt.FromTime(t1))

	fmt.Println("\n--- GeoHash ---")
	bits, latRange, lngRange := geo.NewGeoHash("39.9042", "116.4074", 32)
	centerLat, centerLng := geo.BoxCenter(latRange, lngRange)
	fmt.Printf("GeoHash bits: %d\n", bits)
	fmt.Printf("Center: (%s, %s)\n", centerLat.StringFixed(4), centerLng.StringFixed(4))
	_ = decimal.Zero

	fmt.Println("\n--- QRCodeTool ---")
	qt := &qrcode.QRCodeTool{}
	b64, err := qt.GenerateQrcodeToBase64("https://github.com/sav7ng/go-screws", 256)
	if err != nil {
		fmt.Println("QR error:", err)
	} else {
		fmt.Printf("QR code base64 length: %d chars\n", len(*b64))
	}

	fmt.Println("\n=== All tools working! ===")
}
