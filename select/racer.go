package IDENT

import (
	"fmt"
	"net/http"
	"time"
)

/* select:
- Helps you wait on multiple channels
- Sometimes you'll want to includer time.After in one of the cases to prevent system blocking forever

httptest:
- A convenient way of creating test servers so you can have reliable and controllable tests.
- Uses the same interfaces as the real 'net/http' servers which is consistent and less  for you to learn */

var tenSeondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSeondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

/* we have to use make when creating a channel; rather than say var ch chan struct{}.
When you use var the variable will be initialised with the "zero" value of the type.
So for string it is "", int it is 0, etc. */

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
