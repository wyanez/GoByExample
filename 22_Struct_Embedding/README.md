# Struct Embedding

Go supports embedding of structs and interfaces to express a more seamless
composition of types. This is not to be confused with //go:embed which is
a go directive introduced in Go version 1.16+ to embed files and folders
into the application binary.

``` golang
type base struct {
	num int
}

//A container embeds a base. An embedding looks like a field without a name.
type container struct {
	base
	str string
}
```

You can view and run this example here:
https://go.dev/play/p/rABQsCbjGws