package signleton

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	var wg = new(sync.WaitGroup)
	wg.Add(30)
	for range 30 {
		go func(sync *sync.WaitGroup) {
			getInstance()
			sync.Done()
		}(wg)
	}
	wg.Wait()

}

func TestGetInstanceSyncOnce(t *testing.T) {
	var wg = new(sync.WaitGroup)
	wg.Add(30)
	for range 30 {
		go func(sync *sync.WaitGroup) {
			getInstanceSyncOnce()
			sync.Done()
		}(wg)
	}
	wg.Wait()

}
