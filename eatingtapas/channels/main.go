package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	startTime := flag.Int("min", 30, "min duration for eatng morsel")
	stopTime := flag.Int("max", 180, "max duration for eatng morsel")
	flag.Parse()

	if *startTime > *stopTime {
		fmt.Println("Wrong params ... exit")
		os.Exit(1)
	}
	dishes := []string{"chorizo", "chopitos", "pimientos de padrón", "patatas bravas", "croquetas"}
	dinner := createDinner(dishes)

	go dinner.start()
	fmt.Println("Bon appétit!")

	friends := []string{"Alice", "Bob", "Dave", "Charlie"}
	var wg sync.WaitGroup
	for _, friend := range friends {
		wg.Add(1)
		go eat(&wg, friend, dinner, *startTime, *stopTime)
	}

	wg.Wait()
	close(dinner.ch)
	fmt.Println("That was delicious!")
}

func eat(wg *sync.WaitGroup, name string, dinner *Dinner, min, max int) {
	defer wg.Done()
	for {
		dishC := make(chan string)
		dinner.ch <- dishC
		dish := <-dishC
		if dish == "" {
			fmt.Printf("%s finished eating\n", name)
			return
		}
		fmt.Printf("%s is enjoying some %s\n", name, dish)
		time.Sleep(time.Duration(randInt(min, max)) * time.Second)
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
