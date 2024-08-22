package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		want := 1000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(want)
		for i := 0; i < want; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, want)

	})

}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.value != want {
		t.Errorf("got %d,want %d", got.value, want)
	}

}
