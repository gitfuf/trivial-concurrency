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

type Dinner struct {
	dishes map[string]int
	mu     sync.Mutex
}

func main() {
	startTime := flag.Int("min", 30, "min duration for eatng morsel")
	stopTime := flag.Int("max", 180, "max duration for eatng morsel")
	flag.Parse()

	if *startTime >= *stopTime {
		fmt.Println("Wrong params ... exit")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
	var dishes = []string{"chorizo", "chopitos", "pimientos de padrón", "patatas bravas", "croquetas"}
	dinner := createDinner(dishes)
	fmt.Println("dinner:", dinner)
	fmt.Println("Bon appétit!")

	var wg sync.WaitGroup

	friends := []string{"Alice", "Bob", "Dave", "Charlie"}
	for _, friend := range friends {
		wg.Add(1)
		go eatDinner(&wg, friend, dinner, *startTime, *stopTime)
	}
	wg.Wait()
	fmt.Println("That was delicious!")
}

func eatDinner(wg *sync.WaitGroup, name string, dinner *Dinner, min, max int) {
	defer wg.Done()
	for {
		morsel := pickMorsel(dinner)
		if morsel == "" {
			fmt.Printf("%s finished eating\n", name)
			return
		}
		fmt.Printf("%s is enjoying some %s\n", name, morsel)
		time.Sleep(time.Duration(randInt(min, max)) * time.Second)
	}
}

func pickMorsel(dinner *Dinner) string {
	ret := ""
	dinner.mu.Lock()
	defer dinner.mu.Unlock()

	if len(dinner.dishes) == 0 {
		return ret
	}

	//get dish
	var list []string
	for k, _ := range dinner.dishes {
		list = append(list, k)
	}
	ret = getRandomDish(list)

	//get morsel from this dish
	markEatingMorsel(dinner, ret)

	return ret
}

func createDinner(list []string) *Dinner {
	d := make(map[string]int)
	for _, dish := range list {
		d[dish] = randInt(5, 10)
	}
	return &Dinner{dishes: d}
}

func getRandomDish(list []string) string {
	r := randInt(0, len(list))
	return list[r]
}

//func used to remove chosen morsel from the Dinner object
func markEatingMorsel(dinner *Dinner, dish string) {
	count := dinner.dishes[dish]
	if count == 1 {
		//it is last morsel so remove dish from dinner
		fmt.Printf("%s is last morsel \n", dish)
		delete(dinner.dishes, dish)
		return
	}
	dinner.dishes[dish] = count - 1
	fmt.Printf("%s is left %d \n", dish, count-1)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
