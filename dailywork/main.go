package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type person struct {
	name string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Let's go for a walk!")

	bob := person{name: "Bob"}
	alice := person{name: "Alice"}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go startPreparations(&wg, bob)
	go startPreparations(&wg, alice)
	wg.Wait()

	alarmC := make(chan struct{})
	go alarm(alarmC)

	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("Alarm is counting down")
	}()
	go startPuttingShoes(&wg, bob)
	go startPuttingShoes(&wg, alice)

	wg.Wait()
	fmt.Println("Exiting and locking the door")
	<-alarmC

}

func startPreparations(wg *sync.WaitGroup, p person) {
	defer wg.Done()
	fmt.Printf("%s started getting ready\n", p.name)

	rInt := randInt(60, 90)
	time.Sleep(time.Duration(rInt) * time.Second)
	fmt.Printf("%s spent %d seconds getting ready\n", p.name, time.Duration(rInt))
}

func startPuttingShoes(wg *sync.WaitGroup, p person) {
	defer wg.Done()

	fmt.Printf("%s started putting on shoes\n", p.name)

	rInt := randInt(35, 45)
	time.Sleep(time.Duration(rInt) * time.Second)
	fmt.Printf("%s spent %d seconds putting on shoes\n", p.name, time.Duration(rInt))
}

func alarm(doneC chan<- struct{}) {
	fmt.Println("Arming alarm")
	time.Sleep(60 * time.Second)
	fmt.Println("Alarm is armed")
	doneC <- struct{}{}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
