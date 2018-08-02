package main

import (
	"fmt"
)

type Dish struct {
	name      string
	morselNum int
}

type Dinner struct {
	dishes []Dish
	ch     chan chan string
}

func createDinner(list []string) *Dinner {
	ch := make(chan chan string)
	d := make([]Dish, len(list))
	for i, dish := range list {
		d[i] = Dish{
			name:      dish,
			morselNum: randInt(5, 10),
		}
	}
	return &Dinner{dishes: d, ch: ch}
}

func (d *Dinner) start() {
	for retCh := range d.ch {
		if len(d.dishes) == 0 {
			//say that no food left
			retCh <- ""
			continue
		}

		ind := randInt(0, len(d.dishes))
		dish := d.dishes[ind].name

		count := d.dishes[ind].morselNum
		if count == 1 {
			//it is last morsel so remove dish from dinner
			fmt.Printf("%s is last morsel \n", dish)
			d.dishes = append(d.dishes[0:ind], d.dishes[ind+1:]...)
		} else {
			d.dishes[ind].morselNum--
		}
		retCh <- dish
		//fmt.Printf("%s is left %d \n", dish, count-1)
	}
	fmt.Println("stop dinner")
}
