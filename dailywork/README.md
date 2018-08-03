##  Inspired by Exercise 1: The Daily Walk

### Task description:

Every morning, Alice and Bob go for a walk, and being creatures of habit, they follow the same routine every day.

First, they both prepare, grabbing sunglasses, perhaps a belt, closing open windows, turning off ceiling fans, and pocketing their phones and keys.

Once they’re both ready, which typically takes each of them between 60 and 90 seconds, they arm the alarm, which has a 60 second delay.

While the alarm is counting down, they both put on their shoes, a process which tends to take each of them between 35 and 45 seconds.

Then they leave the house together and lock the door, before the alarm has finished its countdown.

### Ouput:

```
Let's go for a walk!
Bob started getting ready
Alice started getting ready
Alice spent 72 seconds getting ready
Bob spent 76 seconds getting ready
Arming alarm.
Bob started putting on shoes
Alarm is counting down.
Alice started putting on shoes
Alice spent 37 seconds putting on shoes
Bob spent 39 seconds putting on shoes
Exiting and locking the door.
Alarm is armed.
```

### Run app

`go run -race main.go`

`-race` flas is not really need, just use it for check data race