timer
=====
This package implements an reusable timer in an easy-to-use interface that
wraps around time.Timer.

Usage
-----
```go
import "github.com/prep/timer"
```

```go
func readUntilTimeout(ioChan <-chan []byte) {
	timeout := timer.New(time.Second * 5)

	for {
		select {
		case _, ok := <-ioChan:
			if !ok {
				fmt.Println("Channel closed, so stop reading")
				timeout.Stop()
				return
			}

			fmt.Println("Data received, so reset the timer wait for the next event")
			timeout.Reset()

		case <-timeout.C:
			fmt.Println("Timer triggered, so stop reading")
			return
		}
	}
}
```
