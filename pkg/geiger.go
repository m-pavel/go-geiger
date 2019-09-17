package geiger

import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

type GeigerCounter struct {
	r   bool
	pin rpio.Pin
	c   int64
}

func New() *GeigerCounter {
	return &GeigerCounter{r: true}
}

func (g GeigerCounter) Init(pinid int) error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	g.pin = rpio.Pin(pinid)
	g.pin.Input()
	g.pin.PullUp()
	g.pin.Detect(rpio.FallEdge)
	go g.loop()
	return nil
}

func (g GeigerCounter) loop() {
	for {
		if !g.r {
			break
		}
		if g.pin.EdgeDetected() {
			g.c++
		}
		time.Sleep(time.Second / 2)
	}
	g.pin.Detect(rpio.NoEdge)
}
func (g GeigerCounter) Read() (int64, error) {
	return g.c, nil
}

func (g *GeigerCounter) Close() {
	g.r = false
}
