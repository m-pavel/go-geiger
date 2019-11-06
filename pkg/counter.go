package geiger

import (
	"log"
	"time"
)

const J305 = 0.00812037037037 // J305β Geiger Tube

type Counter struct {
	c    int64
	c_pm int64
	c_ph int64

	iv_pm int64
	iv_ph int64

	factor         float64
	stop_h, stop_m chan struct{}
	debug          bool
}

func NewCounter(factor float64, debug bool) *Counter {
	c := Counter{factor: factor, debug: debug}
	c.stop_h = make(chan struct{})
	c.stop_m = make(chan struct{})
	go c.countPerHourLoop()
	go c.countPerMinLoop()
	return &c
}

func (c *Counter) countPerMinLoop() {
	for {
		select {
		case <-c.stop_m:
			break
		case <-time.After(time.Minute):
			c.c_pm = c.c - c.iv_pm
			c.iv_pm = c.c
			if c.debug {
				log.Printf("CPM %d\n", c.c_pm)
			}
		}
	}
}
func (c *Counter) countPerHourLoop() {
	for {
		select {
		case <-c.stop_h:
			break
		case <-time.After(time.Hour):
			c.c_ph = c.c - c.iv_ph
			c.iv_ph = c.c
			if c.debug {
				log.Printf("CPH %d\n", c.c_ph)
			}
		}
	}
}

func (c *Counter) Stop() {
	c.stop_h <- struct{}{}
	c.stop_m <- struct{}{}
}

func (c *Counter) Bip() {
	c.c = c.c + 1
}

func (c Counter) PerMinute() int64 {
	return c.c_pm
}

func (c Counter) PerHour() int64 {
	return c.c_ph
}

// µSv/h
func (c Counter) Radiation() float64 {
	return c.factor * float64(c.c_pm)
}
