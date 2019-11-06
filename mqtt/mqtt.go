package main

import (
	"flag"
	_ "net/http"
	_ "net/http/pprof"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/m-pavel/go-geiger/pkg"
	"github.com/m-pavel/go-hassio-mqtt/pkg"
)

type RadiationService struct {
	g   *geiger.GeigerCounter
	pin *int
	dev *string
}
type Request struct {
	CountPerMinute int64   `json:"per_min"`
	CountPerHour   int64   `json:"per_hour"`
	Value          float64 `json:"value"`
}

func (ts *RadiationService) PrepareCommandLineParams() {
	ts.pin = flag.Int("pin", 18, "GPIO data pin")
	ts.dev = flag.String("dev", "/dev/", "GPIO device")
}

func (ts RadiationService) Name() string { return "geiger" }

func (ts *RadiationService) Init(client MQTT.Client, topic, topicc, topica string, debug bool, ss ghm.SendState) error {
	ts.g = geiger.New(debug)
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
