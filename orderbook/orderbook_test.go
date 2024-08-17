package orderbook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitAddOrder(t *testing.T) {
	l := NewLimit(16000)
	n := 3
	size := 50.12

	for i := 0; i < n; i++ {

		o := NewAskOrder(size)
		l.addOrder(o)
		assert.Equal(t, i, o.limitIndex)

	}
					//t,  expectation, actual result
	assert.Equal(t,n, len(l.orders))
	assert.Equal(t, float64(n) * size, l.totalVolume)
}
