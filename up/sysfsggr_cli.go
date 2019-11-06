package main

import (
	"log"

	"time"

	"fmt"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/sysfs"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	pin := sysfs.Pins[402]
	if err := pin.In(gpio.PullNoChange, gpio.BothEdges); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 50; i++ {
		fmt.Println(pin.Read())
		res := pin.WaitForEdge(time.Minute)
		fmt.Println(res)
		fmt.Println(sysfs.Pins[402].Read())
		time.Sleep(500 * time.Millisecond)
	}

	//p := sysfs.Pin{number: 42, name: "foo", root: "/tmp/gpio/priv/"}
	//if l := p.Read(); l != gpio.Low {
	//	t.Fatal("broken pin is always low")
	//}
}
