# Channel Directions

 When using channels as function parameters, you can specify if a channel
 is meant to only send or receive values. This specificity increases the
 type-safety of the program.


This ping function only accepts a channel for sending values.
It would be a compile-time error to try to receive on this channel.
```golang
func ping(pings chan<- string, msg string) {
	pings <- msg
}
```

The pong function accepts one channel for receives (pings) and a second for sends (pongs).
```golang
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}
```
You can view and run this example here:
https://go.dev/play/p/3CU-Axp_AzF
