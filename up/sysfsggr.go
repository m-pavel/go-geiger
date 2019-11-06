package sysfs

import (
	"log"
	"time"

	"fmt"

	geiger "github.com/m-pavel/go-geiger/pkg"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/sysfs"
)

// GeigerCounter object
type sysfsCounter struct {
	pin    *sysfs.Pin
	c      *geiger.Counter
	stop   chan struct{}
	debug  bool
	factor float64
}

func New(factor float64, debug bool) geiger.GeigerCounter {
	gc := sysfsCounter{debug: debug, factor: factor}
	gc.stop = make(chan struct{})
	return &gc
}

func (g *sysfsCounter) Init(pinid int) error {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	var ok bool
	if g.pin, ok = sysfs.Pins[pinid]; !ok {
		return fmt.Errorf("Wrong pin %d", pinid)
	}
	if err := g.pin.Halt(); err != nil {
		return err
	}
	if err := g.pin.In(gpio.PullNoChange, gpio.RisingEdge); err != nil {
		log.Fatal(err)
	}

	go g.loop()
	g.c = geiger.NewCounter(g.factor, g.debug)

	log.Printf("Started loop on pin %d\n", pinid)
	return nil
}

func (g *sysfsCounter) loop() {
	for {
		select {
		case <-g.stop:
			break
		default:
			res := g.pin.WaitForEdge(time.Second * 10)
			if res {
				g.c.Bip()
				if g.debug {
					log.Println("bip")
				}
			}
		}
	}
	log.Println("GPIO routine exited.")
}

func (g *sysfsCounter) Close() error {
	g.stop <- struct{}{}
	g.c.Stop()
	return nil
}

func (g sysfsCounter) PerMinute() int64 {
	return g.c.PerMinute()
}

func (g sysfsCounter) PerHour() int64 {
	return g.c.PerHour()
}

// µSv/h
func (g sysfsCounter) Radiation() float64 {
	return g.c.Radiation()
}
