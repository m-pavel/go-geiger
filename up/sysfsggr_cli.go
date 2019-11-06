package main

import (
	"log"

	"time"

	"fmt"

	"github.com/fsnotify/fsnotify"
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
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/sys/class/gpio/gpio402/value")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
