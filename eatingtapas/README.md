##  Inspired by Exercise 2: Eating Tapas

### Task description:

Four friends are having dinner together. They ordered five dishes to share, each of which consists of between 5 and 10 morsels.

They eat leisurely, spending between 30 seconds and 3 minutes eating each morsel.

### Output:

```
Bon appétit!
Alice is enjoying some chorizo
Bob is enjoying some chopitos
Charlie is enjoying some pimientos de padrón
Dave is enjoying some croquetas
Alice is enjoying some patatas bravas
Charlie is enjoying some chorizo
Dave is enjoying some chopitos
Alice is enjoying some pimientos de padrón
Bob is enjoying some croquetas
Dave is enjoying some patatas bravas
Alice is enjoying some chorizo
Bob is enjoying some chopitos
Charlie is enjoying some pimientos de padrón
Alice is enjoying some croquetas
Bob is enjoying some patatas bravas
Dave is enjoying some chorizo
Charlie is enjoying some chopitos
Alice is enjoying some pimientos de padrón
Dave is enjoying some patatas bravas
Charlie is enjoying some croquetas
Bob is enjoying some chorizo
Alice is enjoying some chopitos
Charlie is enjoying some pimientos de padrón
Charlie is enjoying some patatas bravas
Dave is enjoying some croquetas
Alice is enjoying some chorizo
Bob is enjoying some chopitos
Bob is enjoying some pimientos de padrón
Charlie is enjoying some patatas bravas
Alice is enjoying some chopitos
Dave is enjoying some patatas bravas
Charlie is enjoying some chopitos
Alice is enjoying some patatas bravas
Bob is enjoying some patatas bravas
Dave is enjoying some patatas bravas
That was delicious!
```
### Run app

For setup can use next flags:

```
startTime := flag.Int("min", 30, "min duration for eatng morsel")
stopTime := flag.Int("max", 180, "max duration for eatng morsel")
```

I have two versions of this task:

* using mutex for access dinner food
```
cd mutex
go run -race main.go

```

* using channel over channel string: `ch chan chan string` 

```
cd channels
go run -race *

```

`-race` flas is not really need, just use it for check data race
