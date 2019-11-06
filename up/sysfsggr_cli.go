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

	if err := pin.Halt(); err != nil {
		log.Fatal(err)
	}
	if err := pin.In(gpio.PullNoChange, gpio.BothEdges); err != nil {
		log.Fatal(err)
	}
	fmt.Println(pin.Function())
	for i := 0; i < 50; i++ {
		fmt.Println(pin.Read())
		res := pin.WaitForEdge(time.Second * 10)
		fmt.Println(res)
		time.Sleep(100 * time.Millisecond)
	}

	//p := sysfs.Pin{number: 42, name: "foo", root: "/tmp/gpio/priv/"}
	//if l := p.Read(); l != gpio.Low {
	//	t.Fatal("broken pin is always low")
	//}
}
