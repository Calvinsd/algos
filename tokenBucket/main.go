package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity    int       // capacity of the bucket
	rate        int       // no of tokens to put into bucket every second.
	tokens      int       // no of tokens
	lastUpdated time.Time // keeps track of updated time
	lock        sync.Mutex
}

func (tb *TokenBucket) New(size int, rate int) {
	tb.capacity = size
	tb.rate = rate
	tb.tokens = size
	tb.lastUpdated = time.Now()
}

func (tb *TokenBucket) removeToken() bool {
	tb.lock.Lock()

	defer tb.lock.Unlock()

	tb.refill()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	} else {
		return false
	}
}

func (tb *TokenBucket) refill() {
	elapsedTime := time.Now().Sub(tb.lastUpdated)

	refillTokens := (int(elapsedTime.Seconds()) * tb.rate)

	if refillTokens > 0 {
		tb.lastUpdated = time.Now()
		tb.tokens += refillTokens
	}

	if tb.tokens+refillTokens > tb.capacity {
		tb.tokens = tb.capacity
	}
}

func main() {

	TBucket := TokenBucket{}

	TBucket.New(5, 5)

	for i := 1; i <= 20; i++ {
		if TBucket.removeToken() {
			fmt.Printf("Successfully for val: %v\n", i)
		} else {
			fmt.Printf("Droped val: %v\n", i)
		}
		time.Sleep(time.Duration((100 + 100*int((i/11)))) * time.Millisecond)
	}
}
