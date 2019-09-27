package geiger

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// GeigerCounter object
type GeigerCounter struct {
	pin  rpio.Pin
	c    int64
	c_pm int64
	c_ph int64

	iv_pm int64
	iv_ph int64
	debug bool

	factor float64
	stop   chan struct{}
}

func New(debug bool) *GeigerCounter {
	gc := GeigerCounter{debug: debug}
	gc.stop = make(chan struct{})
	gc.factor = 0.00812037037037 // J305β Geiger Tube
	return &gc
}

func (g *GeigerCounter) Init(pinid int) error {
	if err := rpio.Open(); err != nil {
		return err
	}
	g.pin = rpio.Pin(pinid)
	g.pin.Input()

	g.pin.PullOff()
	g.pin.Detect(rpio.FallEdge)

	go g.loop()
	go g.countPerHourLoop()
	go g.countPerMinLoop()

	log.Printf("Started loop on pin %d\n", pinid)
	return nil
}

func (g *GeigerCounter) loop() {
	for {
		select {
		case <-g.stop:
			break
		case <-time.After(500 * time.Millisecond):
			if g.pin.EdgeDetected() {
				g.c++
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

func (g *GeigerCounter) countPerMinLoop() {
	for {
		select {
		case <-g.stop:
			break
		case <-time.After(time.Minute):
			g.c_pm = g.c - g.iv_pm
			g.iv_pm = g.c
			if g.debug {
				log.Printf("CPM %d\n", g.c_pm)
			}
		}
	}
}
func (g *GeigerCounter) countPerHourLoop() {
	for {
		select {
		case <-g.stop:
			break
		case <-time.After(time.Hour):
			g.c_ph = g.c - g.iv_ph
			g.iv_ph = g.c
			if g.debug {
				log.Printf("CPH %d\n", g.c_ph)
			}
		}
	}
}
func (g GeigerCounter) PerMinute() int64 {
	return g.c_pm
}

func (g GeigerCounter) PerHour() int64 {
	return g.c_ph
}

// µSv/h
func (g GeigerCounter) Radiation() float64 {
	return g.factor * float64(g.c_pm)
}

func (g *GeigerCounter) Close() {
	g.stop <- struct{}{}
}
