package lvlbp

import (
	"math/bits"
	"sync"
	"sync/atomic"
)

// pools contains pools for byte slices of various capacities.
//
//    pools[0] is for capacities from 0 to 8
//    pools[1] is for capacities from 9 to 16
//    pools[2] is for capacities from 17 to 32
//    ...
//    pools[n] is for capacities from 2^(n+2)+1 to 2^(n+3)
//
// Limit the maximum capacity to 2^18, since there are no performance benefits
// in caching byte slices with bigger capacities.
var pools [17]sync.Pool

var stats [17]int64
var overflow int64

// Get returns byte slice with the given capacity.
func Get(capacity int) *[]byte {
	id, capacityNeeded := getPoolIDAndCapacity(capacity)
	for i := 0; i < 2; i++ {
		if id < 0 || id >= len(pools) {
			break
		}
		if v := pools[id].Get(); v != nil {
			atomic.AddInt64(&stats[id], 1)
			return v.(*[]byte)
		}
		id++
	}
	atomic.AddInt64(&overflow, 1)
	b := make([]byte, 0, capacityNeeded)
	return &b
}

// Put returns byte slice to the pool.
func Put(b *[]byte) {
	if b == nil {
		return
	}
	capacity := cap(*b)
	id, poolCapacity := getPoolIDAndCapacity(capacity)
	if capacity <= poolCapacity {
		bb := (*b)[:0]
		pools[id].Put(&bb)
	}
}

func getPoolIDAndCapacity(size int) (idx, capacity int) {
	size--
	if size < 0 {
		size = 0
	}
	size >>= 3

	idx = bits.Len(uint(size))
	if idx >= len(pools) {
		idx = len(pools) - 1
	}
	return idx, (1 << (idx + 3))
}

func Stats() ([17]int64, int64) {
	s := [17]int64{}
	for i := 0; i < 17; i++ {
		s[i] = atomic.LoadInt64(&stats[i])
	}
	return s, atomic.LoadInt64(&overflow)
}
