package lvlbp

import (
	"fmt"
	"testing"
	"time"
)

func TestGetPutConcurrent(t *testing.T) {
	const concurrency = 10
	doneCh := make(chan struct{}, concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			for capacity := -1; capacity < 100; capacity++ {
				bb := *Get(capacity)
				if len(bb) > 0 {
					panic(fmt.Errorf("len(bb) must be zero; got %d", len(bb)))
				}
				if capacity < 0 {
					capacity = 0
				}
				bb = append(bb, make([]byte, capacity)...)
				Put(&bb)
			}
			doneCh <- struct{}{}
		}()
	}

	tc := time.After(10 * time.Second)
	for i := 0; i < concurrency; i++ {
		select {
		case <-tc:
			t.Fatalf("timeout")
		case <-doneCh:
		}
	}

	stats, overflow := Stats()
	totalCount := int(overflow)
	for _, count := range stats {
		totalCount += int(count)
	}
	if want := 1000; totalCount != want {
		t.Fatalf("got %d want %d", totalCount, want)
	}
}
