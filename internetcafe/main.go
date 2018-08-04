package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Client struct {
	name    int
	finishC chan struct{}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

}

func main() {
	compNum := flag.Int("pc", 8, "number of availbale computers")
	clientNum := flag.Int("clients", 25, "number of clients")
	minT := flag.Int("min", 2, "minimum seconds online")
	maxT := flag.Int("max", 5, "maximum seconds online")

	flag.Parse()

	wg := sync.WaitGroup{}
	//client queue
	queue := make(chan Client, *clientNum)

	//stop channel for computers
	stopCh := make(chan struct{})

	fmt.Println("Welcome to our cafe!")
	//power on each computer
	for i := 1; i <= *compNum; i++ {
		go func(computerNum int) {
			for {
				select {
				case <-stopCh:
					fmt.Println("stopCh came ... power off computer ", computerNum)
					return
				case client := <-queue:
					fmt.Printf("Tourist %d is online\n", client.name)
					dur := randInt(*minT, *maxT)
					time.Sleep(time.Duration(dur) * time.Second)
					fmt.Printf("Tourist %d is done, having spent %d seconds online.\n", client.name, dur)
					client.finishC <- struct{}{}
				}
			}
		}(i)
	}

	//clients go into the cafe. run for each goroutine in order to do queue
	//wait when client finished with internet and then exit goroutine
	for i := 1; i <= *clientNum; i++ {
		client := Client{name: i, finishC: make(chan struct{})}
		wg.Add(1)

		go func(wg *sync.WaitGroup, client Client) {
			defer wg.Done()
			queue <- client
			fmt.Printf("Tourist %d waiting for turn.\n", client.name)
			for {
				select {
				case <-client.finishC:
					//client done with internet
					return

				}
			}
		}(&wg, client)

	}
	wg.Wait()

	fmt.Println("The place is empty, let's close up and go to the beach!")
	close(stopCh)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
