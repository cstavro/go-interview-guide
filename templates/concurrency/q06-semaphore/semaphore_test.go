package semaphore

import (
	"sync"
	"testing"
)

func TestSemaphoreLimitsConcurrency(t *testing.T) {
	sem := New(3)
	var active int
	var maxActive int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem.Acquire()
			mu.Lock()
			active++
			if active > maxActive {
				maxActive = active
			}
			mu.Unlock()

			mu.Lock()
			active--
			mu.Unlock()
			sem.Release()
		}()
	}
	wg.Wait()

	if maxActive > 3 {
		t.Errorf("maxActive = %d, want <= 3", maxActive)
	}
}
