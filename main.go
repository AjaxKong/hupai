package main

import (
	"github.com/go-vgo/robotgo"
	"fmt"
	"github.com/kr/pretty"
	"time"
)

const (
	XOFFSET = 30
	YOFFSET = 20
)

func main() {
	fpid, err := robotgo.FindIds("iexplore.exe")
	if err == nil {
		fmt.Printf("%# v", pretty.Formatter(fpid))
	}
	robotgo.ActiveName("iexplore.exe")
	bitmaps := robotgo.OpenBitmap("targetRegion.png")
	bitmapBG := robotgo.OpenBitmap("bg.png")
	bitmap1 := robotgo.OpenBitmap("1.png")
	fxs, fys := robotgo.FindBitmap(bitmaps, bitmapBG)
	fmt.Println("FindBitmap------", fxs, fys)
	x1, y1 := robotgo.FindBitmap(bitmap1, bitmaps)
	fmt.Println("FindBitmap------", fxs+x1, fys+y1)
	bitmap2 := robotgo.OpenBitmap("2.png")
	x2, y2 := robotgo.FindBitmap(bitmap2, bitmaps)
	fmt.Println("FindBitmap------", fxs+x2, fys+y2)
	bitmap3 := robotgo.OpenBitmap("3.png")
	x3, y3 := robotgo.FindBitmap(bitmap3, bitmaps)
	fmt.Println("FindBitmap------", fxs+x3, fys+y3)
	bitmaps2 := robotgo.OpenBitmap("tg2.png")
	bitmapBG2 := robotgo.OpenBitmap("bg2.png")
	bitmap21 := robotgo.OpenBitmap("4.png")
	fxs2, fys2 := robotgo.FindBitmap(bitmaps2, bitmapBG2)
	fmt.Println("FindBitmap------", fxs2, fys2)
	x21, y21 := robotgo.FindBitmap(bitmap21, bitmaps2)
	fmt.Println("FindBitmap------", fxs2+x21, fys2+y21)
	keve := robotgo.AddEvent("space")
	if keve == 0 {

		robotgo.Move(fxs+x1+XOFFSET, fys+y1+YOFFSET)
		robotgo.MouseClick("left", true)
		robotgo.KeyTap("6")
		robotgo.KeyTap("0")
		robotgo.KeyTap("0")
		time.Sleep(time.Duration(100) * time.Millisecond)
		robotgo.Move(fxs+x2+XOFFSET, fys+y2+YOFFSET)
		robotgo.Click("left")

		robotgo.Move(fxs+x3+XOFFSET, fys+y3+YOFFSET)
		robotgo.Click("left")
		keves := robotgo.AddEvent("enter")
		if keves == 0 {
			robotgo.Move(fxs2+x21+XOFFSET, fys2+y21+YOFFSET)

				//if time.Now().Second() >= 57 {
					robotgo.Click("left")
				//}

		}
	}
	/*bitmapcj := robotgo.OpenBitmap("test3.png")
	fmt.Println("...", bitmapcj)

	fxcj, fycj := robotgo.FindBitmap(bitmapcj)

	bitmap := robotgo.OpenBitmap("test.png")
	fmt.Println("...", bitmap)
	robotgo.Move(fxcj + 10, fycj + 10)
	robotgo.Click("left")

	fx, fy := robotgo.FindBitmap(bitmap)
	fmt.Println("FindBitmap------", fx, fy)
	keve := robotgo.AddEvent("enter")
	if keve == 0 {
		robotgo.Move(fx + 10, fy + 10)
		robotgo.Click("left")
	}*/
}