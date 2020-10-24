package main

import (
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Title:  "test",
			Height: 500,
			Width:  600,
		})
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()

		for {
			e := w.NextEvent() //イベントを受け取る

			switch e := e.(type) {
			case lifecycle.Event:
				//fmt.Printf("lifecycle.Event%v¥", e)
				if e.To == lifecycle.StageDead {
					return
				}

			case key.Event:
				//fmt.Printf("key.Event%v¥", e)
				if e.Code == key.CodeEscape {
					return
				}

			case error:
				log.Print(e)
			}
		}
	})
}
