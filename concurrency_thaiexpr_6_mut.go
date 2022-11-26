package main

import "sync"

// mutex => Mutual execution = share resources between thread
// ไม่ใช้งานพร้อมกัน for avoid Race condition

// mutex.Lock()
// // จะมีแค่ thread เดียวที่เข้ามา run code ส่วนนี้ได้
// mutex.Unlock()

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increment the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can acces the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Value returns the current value of the counter for the given key
	defer c.mux.Unlock()
	return c.v[key]
}
func main() {

}
