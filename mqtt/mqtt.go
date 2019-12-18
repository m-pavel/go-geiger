package main

import (
	"flag"
	_ "net/http"
	_ "net/http/pprof"

	"fmt"

	"github.com/m-pavel/go-geiger/pkg"
	"github.com/m-pavel/go-geiger/rpi"
	"github.com/m-pavel/go-geiger/up"
	"github.com/m-pavel/go-hassio-mqtt/pkg"
)

type RadiationService struct {
	ghm.NonListerningService
	g   geiger.GeigerCounter
	pin *int
	typ *string
}
type Request struct {
	CountPerMinute int64   `json:"per_min"`
	CountPerHour   int64   `json:"per_hour"`
	Value          float64 `json:"value"`
}

func (ts *RadiationService) PrepareCommandLineParams() {
	ts.pin = flag.Int("pin", 18, "GPIO data pin")
	ts.typ = flag.String("type", "rpi", "Board type [rpi|up]")
}

func (ts RadiationService) Name() string { return "geiger" }

func (ts *RadiationService) Init(ctx *ghm.ServiceContext) error {
	if *ts.typ == "rpi" {
		ts.g = rpi.New(geiger.J305, ctx.Debug())
	} else if *ts.typ == "up" {
		ts.g = sysfs.New(geiger.J305, ctx.Debug())
	} else {
		return fmt.Errorf("Wrong board type %s", *ts.typ)
	}
	return ts.g.Init(*ts.pin)
}

func (ts RadiationService) Do() (interface{}, error) {
	return &Request{CountPerMinute: ts.g.PerMinute(), CountPerHour: ts.g.PerHour(), Value: ts.g.Radiation()}, nil
}

func (ts RadiationService) Close() error {
	ts.g.Close()
	return nil
}

func main() {
	ghm.NewStub(&RadiationService{}).Main()
}
