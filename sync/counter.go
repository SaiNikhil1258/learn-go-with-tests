package sync

import "sync"

type Counter struct {
	sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

/*
NOTE: A wait group waits for a collection of goroutines to finish. The main go routine calls
Add to set the number of go routines to wait for. Then each of the go routine runs and calls
Done when finished. At the same time, wait can be used to block untill all  goroutines have finished.
*/
