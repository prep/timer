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
		case _ = <-ioChan:
			fmt.Println("Data received")
			timeout.Reset()

		case <-timeout.C:
			fmt.Println("Timer triggered!")
			return
		}
	}
}
```
