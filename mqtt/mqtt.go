package main

import (
	_ "net/http"
	_ "net/http/pprof"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/m-pavel/go-geiger/pkg"
	"github.com/m-pavel/go-hassio-mqtt/pkg"
)

type RadiationService struct {
	g *geiger.GeigerCounter
}
type Request struct {
	Radiation int64 `json:"radiation"`
}

func (ts RadiationService) PrepareCommandLineParams() {}
func (ts RadiationService) Name() string              { return "geiger" }

func (ts *RadiationService) Init(client MQTT.Client, topic, topicc, topica string, debug bool) error {
	ts.g = geiger.New()
	return ts.g.Init(12)
}

func (ts RadiationService) Do(client MQTT.Client) (interface{}, error) {
	v, err := ts.g.Read()
	if err != nil {
		return nil, err
	}
	return &Request{Radiation: v}, nil
}

func (ts RadiationService) Close() error {
	ts.g.Close()
	return nil
}

func main() {
	ghm.NewStub(&RadiationService{}).Main()
}
