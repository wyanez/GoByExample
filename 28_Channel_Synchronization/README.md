# Channel Synchronization

We can use channels to synchronize execution across goroutines.

Here’s an example of using a blocking receive to wait for a goroutine to finish.

When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.

You can view and run this example here:
https://go.dev/play/p/jskcpAXny66