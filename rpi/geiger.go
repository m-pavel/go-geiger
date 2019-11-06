package rpi

import (
	"log"
	"time"

	geiger "github.com/m-pavel/go-geiger/pkg"
	"github.com/stianeikeland/go-rpio"
)

// GeigerCounter object
type rpiGeigerCounter struct {
	pin    rpio.Pin
	c      *geiger.Counter
	stop   chan struct{}
	debug  bool
	factor float64
}

func New(factor float64, debug bool) geiger.GeigerCounter {
	gc := rpiGeigerCounter{debug: debug, factor: factor}
	gc.stop = make(chan struct{})
	return &gc
}

func (g *rpiGeigerCounter) Init(pinid int) error {
	if err := rpio.Open(); err != nil {
		return err
	}
	g.pin = rpio.Pin(pinid)
	g.pin.Input()

	g.pin.PullOff()
	g.pin.Detect(rpio.FallEdge)

	go g.loop()
	g.c = geiger.NewCounter(g.factor, g.debug)

	log.Printf("Started loop on pin %d\n", pinid)
	return nil
}

func (g *rpiGeigerCounter) loop() {
	for {
		select {
		case <-g.stop:
			break
		case <-time.After(500 * time.Millisecond):
			if g.pin.EdgeDetected() {
				g.c.Bip()
				if g.debug {
					log.Println("bip")
				}
			}
		}
	}
	g.pin.Detect(rpio.NoEdge)
	rpio.Close()
	log.Println("GPIO routine exited.")
}

func (g *rpiGeigerCounter) Close() error {
	g.stop <- struct{}{}
	g.c.Stop()
	return nil
}

func (g rpiGeigerCounter) PerMinute() int64 {
	return g.c.PerMinute()
}

func (g rpiGeigerCounter) PerHour() int64 {
	return g.c.PerHour()
}

// µSv/h
func (g rpiGeigerCounter) Radiation() float64 {
	return g.c.Radiation()
}
