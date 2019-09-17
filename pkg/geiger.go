package geiger

import "github.com/stianeikeland/go-rpio"

type GeigerCounter struct {
}

func (g GeigerCounter) Init(pinid int) error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	pin := rpio.Pin(pinid)
	pin.Input()
	return nil
}

func (g GeigerCounter) Read() (float32, error) {
	return 0, nil
}
