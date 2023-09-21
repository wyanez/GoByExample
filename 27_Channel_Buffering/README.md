# Channel Buffering

By default channels are unbuffered, meaning that they will only accept
sends (chan <-) if there is a corresponding receive (<- chan) ready
to receive the sent value.
Buffered channels accept a limited number of values without a corresponding
receiver for those values.

Here we make a channel of strings buffering up to 2 values
```golang
messages := make(chan string, 2)
```

Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
```golang
	messages <- "buffered"
	messages <- "channel"
```
*Important: a third attempt  of send value to channel throw a runtime error (the channel cappacity is 2)*


Later we can receive these two values as usual.
```golang
	fmt.Println(<-messages)
	fmt.Println(<-messages)
```

You can view and run this example here:
https://go.dev/play/p/n9MUiMJx_mC