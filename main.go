package main

import (
	"sync"
	"time"

	"github.com/nproc/rpi-board-info/raspberry"
	"github.com/nproc/rpi-board-info/thermal"
)

func main() {
	raspberry := raspberry.NewPi(thermal.Thermometer{})

	ticker := time.NewTicker(5 * time.Second)

	stop := make(chan bool, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)

	defer func() {
		stop <- true
	}()

	go func() {
		for {
			select {
			case <-ticker.C:
				raspberry.RecordCurrentTemperature()
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()

	wg.Wait()
}
