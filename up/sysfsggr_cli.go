package main

import (
	"log"

	"fmt"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/sysfs"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	if err := sysfs.Pins[402].In(gpio.PullDown, gpio.BothEdges); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 50; i++ {
		fmt.Println(sysfs.Pins[402].Read())
	}

	//p := sysfs.Pin{number: 42, name: "foo", root: "/tmp/gpio/priv/"}
	//if l := p.Read(); l != gpio.Low {
	//	t.Fatal("broken pin is always low")
	//}
}
