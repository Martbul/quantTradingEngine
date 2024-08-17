package orderbook

import (
	"math/rand"
	"time"
)

type Limit struct {
	price       float64
	orders      []*Order
	totalVolume float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		price:  price,
		orders: []*Order{},
	}
}

func (l *Limit) addOrder(o *Order){
	o.limitIndex = len(l.orders)
	l.orders = append(l.orders, o)
	l.totalVolume += o.size
}


//! dont understand this
// VISUAL EXAMPLE:
	// l.orders = [Order1, Order2, Order3, Order4] -> 
		//  l.orders[1] = l.orders[3]  // [Order1, Order4, Order3, Order4] ->
		 	//l.orders = l.orders[:3]  // [Order1, Order4, Order3]  (Order4 was moved to index 1)

func (l *Limit) deleteOrder(o *Order){
	l.orders[o.limitIndex] = l.orders[len(l.orders) -1]

	// `:` means range -> arr := []int{10, 20, 30, 40, 50}
								// subSlice := arr[:3]
								// fmt.Println(subSlice)  // Output: [10, 20, 30]
	l.orders = l.orders[:len(l.orders) - 1]

	if !o.isFilled() {
		l.totalVolume -= o.size
	}
}

type Order struct {
	id        int64
	size      float64
	timestamp int64
	isBid     bool //bool is 1 byte
	limitIndex int
}

func NewOrder(isBid bool, size float64) *Order {
	return &Order{
		id: rand.Int63n(1000000),
		size:size,
		timestamp: time.Now().UnixNano(),
		isBid:isBid,
	}
}


func NewBidOrder(size float64) *Order{
	return NewOrder(false, size)
}
func NewAskOrder(size float64) *Order{
	return NewOrder(false, size)
}

func (o *Order) isFilled() bool{
	return o.size == 0
}