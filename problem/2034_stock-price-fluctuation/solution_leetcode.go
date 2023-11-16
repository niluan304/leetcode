package main

import "container/heap"

// 方法二：哈希表 + 两个优先队列
// 方法一使用一个有序集合存储每个股票价格的次数，在更新操作中将有序集合中的过期价格删除完毕，在其余操作中直接得到答案并返回。
// 可以换一个思路，删除过期价格不一定要在更新操作中完成，而是可以在返回股票最高价格操作和返回股票最低价格操作中完成，即延迟删除。
//
// 为了实现延迟删除，需要维护两个优先队列用于存储股票价格和时间戳，分别基于大根堆和小根堆实现，大根堆的堆顶元素对应股票最高价格，
// 小根堆的堆顶元素对应股票最低价格。以下将基于大根堆实现的优先队列称为最高价格队列，将基于小根堆实现的优先队列称为最低价格队列。
//
// 对于更新操作，使用 timestamp 更新最大时间戳，将 timestamp 和 price 存入哈希表，并将 (price,timestamp) 分别加入两个优先队列。
//
// 对于返回股票最新价格操作，从哈希表中得到最大时间戳对应的股票价格并返回。
//
// 对于返回股票最高价格操作，每次从最高价格队列的队首元素中得到价格和时间戳，并从哈希表中得到该时间戳对应的实际价格，
// 如果队首元素中的价格和实际价格不一致，则队首元素为过期价格，将队首元素删除，重复该操作直到队首元素不为过期价格，此时返回队首元素中的价格。
//
// 对于返回股票最低价格操作，每次从最低价格队列的队首元素中得到价格和时间戳，并从哈希表中得到该时间戳对应的实际价格，
// 如果队首元素中的价格和实际价格不一致，则队首元素为过期价格，将队首元素删除，重复该操作直到队首元素不为过期价格，此时返回队首元素中的价格
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/stock-price-fluctuation/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

type StockPrice2 struct {
	maxPrice, minPrice hp
	timePriceMap       map[int]int
	maxTimestamp       int
}

func Constructor2() StockPrice2 {
	return StockPrice2{timePriceMap: map[int]int{}}
}

func (sp *StockPrice2) Update(timestamp, price int) {
	heap.Push(&sp.maxPrice, pair{price: -price, timestamp: timestamp})
	heap.Push(&sp.minPrice, pair{price: price, timestamp: timestamp})
	sp.timePriceMap[timestamp] = price
	if timestamp > sp.maxTimestamp {
		sp.maxTimestamp = timestamp
	}
}

func (sp *StockPrice2) Current() int {
	return sp.timePriceMap[sp.maxTimestamp]
}

func (sp *StockPrice2) Maximum() int {
	for {
		if p := sp.maxPrice[0]; -p.price == sp.timePriceMap[p.timestamp] {
			return -p.price
		}
		heap.Pop(&sp.maxPrice)
	}
}

func (sp *StockPrice2) Minimum() int {
	for {
		if p := sp.minPrice[0]; p.price == sp.timePriceMap[p.timestamp] {
			return p.price
		}
		heap.Pop(&sp.minPrice)
	}
}

type pair struct{ price, timestamp int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
