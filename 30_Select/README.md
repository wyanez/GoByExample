# Select

Go’s select lets you wait on multiple channel operations.  
Combining goroutines and channels with select is a powerful feature of Go.

We’ll use select to await both of these values simultaneously,
printing each one as it arrives.
```golang
select {
case msg1 := <-c1:
    fmt.Println("received ", msg1)
case msg2 := <-c2:
    fmt.Println("received ", msg2)
}
```

You can view and run this example here:
https://go.dev/play/p/BoTvzJ8eymD
