package main

import (
	"fmt"
	"image/color"
	"time"

	"tinygo.org/x/drivers/examples/ili9341/initdisplay"
	"tinygo.org/x/drivers/ili9341"
)

var (
	black   = color.RGBA{0, 0, 0, 255}
	white   = color.RGBA{255, 255, 255, 255}
	blue    = color.RGBA{0, 0, 255, 255}
	display *ili9341.Device
)

func main() {
	display = initdisplay.InitDisplay()
	width, height := display.Size()
	display.FillRectangle(0, 0, width, height, black)

	if true {
		go func() {
			for {
				display.FillRectangle(0, height/2, width/2, height/2, black)
				time.Sleep(600 * time.Millisecond)
				display.FillRectangle(0, height/2, width/2, height/2, white)
				time.Sleep(600 * time.Millisecond)
			}
		}()
	}

	if true {
		go func() {
			for {
				display.FillRectangle(width/2, height/2, width/2, height/2, blue)
				time.Sleep(1000 * time.Millisecond)
				display.FillRectangle(width/2, height/2, width/2, height/2, white)
				time.Sleep(1000 * time.Millisecond)
			}
		}()
	}

	if true {
		main2()
	}

	select {}
}

func main2() {
	label1 := NewLabel(320, 30)
	label2 := NewLabel(320, 30)

	chCnt1 := make(chan uint32, 1)
	chCnt2 := make(chan uint32, 1)

	go timer77ms(chCnt1)
	go timer500ms(chCnt2)

	for {
		select {
		case cnt := <-chCnt1:
			label1.SetText(fmt.Sprintf("77ms : %04X", cnt), white)
			display.DrawRGBBitmap(0, 40, label1.buf, label1.w, label1.h)
		case cnt := <-chCnt2:
			label2.SetText(fmt.Sprintf("500ms: %04X", cnt), white)
			display.DrawRGBBitmap(0, 80, label2.buf, label2.w, label2.h)
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func timer77ms(ch chan<- uint32) {
	cnt := uint32(0)
	for {
		ch <- cnt
		cnt++
		time.Sleep(77 * time.Millisecond)
	}
}

func timer500ms(ch chan<- uint32) {
	cnt := uint32(0)
	for {
		ch <- cnt
		cnt++
		time.Sleep(500 * time.Millisecond)
	}
}
