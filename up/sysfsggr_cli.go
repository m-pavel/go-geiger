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
	if s, err := host.Init(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(s)
	}

	pin := sysfs.Pins[402]

	if err := pin.Halt(); err != nil {
		log.Fatal(err)
	}
	if err := pin.In(gpio.PullNoChange, gpio.RisingEdge); err != nil {
		log.Fatal(err)
	}
	fmt.Println(pin.Function())
	for i := 0; i < 500; i++ {
		fmt.Println(pin.Read())
		//res := pin.WaitForEdge(time.Second * 10)
		//fmt.Println(res)
		time.Sleep(10 * time.Millisecond)
	}

}
