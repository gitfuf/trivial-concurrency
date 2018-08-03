##  Inspired by Exercise 3: Internet Café

### Task description:

A small internet café in a village just outside of Manilla has 8 computers, which are available on a first-come first-serve basis. When all the computers are taken, the next person in line has to wait until a computer frees up. This morning several groups of tourists, 25 people in all, are waiting when the doors open. Each person spends between 15 minutes and 2 hours online.

### Sample output:

```
Tourist 4 is online.
Tourist 25 is online.
....
Tourist 3 waiting for turn.
Tourist 18 waiting for turn.
....
Tourist 2 is done, having spent 15 minutes online.
Tourist 17 is online.
Tourist 5 is done, having spent 18 minutes online.
Tourist 6 is online.
Tourist 15 is done, having spent 40 minutes online.
....
Tourist 21 is done, having spent 111 minutes online.
The place is empty, let's close up and go to the beach!
```

### Run app

I've changed initial params in order not to wait minutes and hours so use **only seconds**. For setup can use next flags:

```
compNum := flag.Int("pc", 8, "number of availbale computers")
clientNum := flag.Int("clients", 25, "number of clients")
minT := flag.Int("min", 2, "minimum seconds online")
maxT := flag.Int("max", 5, "maximum seconds online")
```

`go run -race main.go -min=2 -max=4 -pc=2 -clients=4` has next output:

`-race` flas is not really need, just use it for check data race

```
Welcome to our cafe!
Tourist 1 waiting for turn.
Tourist 1 is online
Tourist 2 waiting for turn.
Tourist 2 is online
Tourist 3 waiting for turn.
Tourist 4 waiting for turn.
Tourist 2 is done, having spent 2 seconds online.
Tourist 3 is online
Tourist 1 is done, having spent 2 seconds online.
Tourist 4 is online
Tourist 3 is done, having spent 3 seconds online.
Tourist 4 is done, having spent 3 seconds online.
The place is empty, let's close up and go to the beach!
```