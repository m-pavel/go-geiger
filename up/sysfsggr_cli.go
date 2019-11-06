package main

import (
	"log"

	"time"

	"fmt"

	"os"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/sysfs"
)

func main() {
	fstest()

}

func gpiotest() {
	if s, err := host.Init(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(s)
	}

	pin := sysfs.Pins[402]

	if err := pin.Halt(); err != nil {
		log.Fatal(err)
	}
	if err := pin.In(gpio.PullNoChange, gpio.NoEdge); err != nil {
		log.Fatal(err)
	}
	fmt.Println(pin.Function())
	if err := pin.In(gpio.PullNoChange, gpio.RisingEdge); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 500; i++ {
		fmt.Println(pin.Read())
		//res := pin.WaitForEdge(time.Second * 10)
		//fmt.Println(res)
		time.Sleep(10 * time.Millisecond)
	}
}

func fstest() {
	f, err := os.Open("/sys/class/gpio/gpio402/value")
	if err != nil {
		log.Fatal(err)
	}
	bt := make([]byte, 1)
	for i := 0; i < 1000; i++ {
		f.Seek(0, 0)
		f.Read(bt)
		fmt.Printf("%d ", bt[0])
	}
	f.Close()
}
