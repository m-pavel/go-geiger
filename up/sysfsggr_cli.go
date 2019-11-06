package main

import (
	"fmt"
	"log"

	"periph.io/x/periph/host"
	"periph.io/x/periph/host/sysfs"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sysfs.Pins[402])
	//p := sysfs.Pin{number: 42, name: "foo", root: "/tmp/gpio/priv/"}
	//if l := p.Read(); l != gpio.Low {
	//	t.Fatal("broken pin is always low")
	//}
}
