package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
}

func main() {

	TBucket := TokenBucket{}

	TBucket.New(5, 2)

	done := make(chan bool)

	go handleShutdowns(done)

	for i := 1; i <= 20; i++ {
		if TBucket.removeToken() {
			fmt.Printf("Successfully for val: %v\n", i)
		} else {
			fmt.Printf("Droped val: %v\n", i)
		}
		time.Sleep(time.Duration((100 + 100*int((i/11)))) * time.Millisecond)
	}

	fmt.Println("Wating for shutdowns....")

	<-done

	fmt.Println("Shutting down")
}

func handleShutdowns(done chan<- bool) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGSEGV)

	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			fmt.Println("Encountered os interrupt")
			done <- true
		case syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGSEGV:
			fmt.Println("Received linux signel")
			done <- true
		}
	}()
}
