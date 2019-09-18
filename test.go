package main

import (
	"time"

	"fmt"
	geiger "github.com/m-pavel/go-geiger/pkg"
)

func main() {
	gc := geiger.New(true)
	gc.Init(18)
	defer gc.Close()
	time.Sleep(3 * time.Minute)
	fmt.Println(gc.PerMinute())
	fmt.Println(gc.PerHour())
	fmt.Println(gc.Radiation())
}
