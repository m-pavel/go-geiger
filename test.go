package main

import (
	"time"

	"flag"
	"fmt"

	geiger "github.com/m-pavel/go-geiger/pkg"
	"github.com/stianeikeland/go-rpio"
)

func main() {
	read := flag.Bool("read", false, "")
	pin := flag.Int("pin", 18, "")
	flag.Parse()

	if *read {
		pin := rpio.Pin(*pin)
		pin.Input()

		pin.PullOff()
		pin.Detect(rpio.FallEdge)
		for {
			fmt.Print(pin.Read())
			time.Sleep(500 * time.Millisecond)
		}
	} else {
		gc := geiger.New(true)
		gc.Init(*pin)
		defer gc.Close()
		time.Sleep(3 * time.Minute)
		fmt.Println(gc.PerMinute())
		fmt.Println(gc.PerHour())
		fmt.Println(gc.Radiation())
	}
}
