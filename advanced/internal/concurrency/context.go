package concurrency

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ContextMaze() {
	base := context.Background()
	//go waitForContextCancel(base, "Base")

	childWithCancel, cancelFn := context.WithCancel(base)
	go waitForContextCancel(childWithCancel, "Child-with-Cancel")
	grandChild, _ := context.WithCancel(childWithCancel)
	go waitForContextCancel(grandChild, "Grand-Child-with-Cancel")
	grandGrandChild := context.WithoutCancel(grandChild)
	go waitForContextCancel(grandGrandChild, "Grand-Grand-Child-without-Cancel")

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Canceling Child context")
		cancelFn()
	}()

	childWithTimeout, _ := context.WithTimeout(base, time.Second) // why does Goland complain?
	go waitForContextCancel(childWithTimeout, "Child-with-Timeout")

	childWithValue := context.WithValue(base, "some-key", "some-value")
	value, ok := childWithValue.Value("some-key").(string)
	if ok {
		fmt.Printf("The value in the context was %v\n", value)
	}

	time.Sleep(3 * time.Second)
}

func waitForContextCancel(ctx context.Context, name string) {
	select {
	case <-ctx.Done():
		fmt.Printf("%s context canceled\n", name)
	}
}

// FROM https://github.com/kubernetes-sigs/controller-runtime, for demonstration purposes
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
var onlyOneSignalHandler = make(chan struct{})

// SetupSignalHandler registers for SIGTERM and SIGINT. A context is returned
// which is canceled on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() context.Context {
	close(onlyOneSignalHandler) // panics when called twice

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		cancel()
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return ctx
}
