package concurrent

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestStripedMutex_GetLock(t *testing.T) {

	c := NewStripedMutex(64)
	for i := 0; i < 100; i++ {
		c.GetLock(strconv.Itoa(rand.Int()))
	}
}
