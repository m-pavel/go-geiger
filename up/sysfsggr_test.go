package sysfsggr

import (
	"log"
	"testing"

	"fmt"

	"github.com/google/periph/host"
	"github.com/google/periph/host/sysfs"
)

func Test1(t *testing.T) {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sysfs.Pins)
	//p := sysfs.Pin{number: 42, name: "foo", root: "/tmp/gpio/priv/"}
	//if l := p.Read(); l != gpio.Low {
	//	t.Fatal("broken pin is always low")
	//}
}
