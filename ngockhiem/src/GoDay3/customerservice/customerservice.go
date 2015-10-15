package main

import (
	"fmt"
	//	"math/rand"
	"sync"
	"time"
)

var (
	wg      sync.WaitGroup
	m       sync.Mutex
	clock   int              = 0
	workers map[*Worker]bool = make(map[*Worker]bool)
)

type Worker struct {
	name         string
	available_at int
}
type Request struct {
	name string
	time int
}

func Handle(ch chan Request, exit chan bool) {
	defer wg.Done()
	for {
		select {
		case rq := <-ch:
			wg.Add(1)
			go HandleCustomerReq(rq)
		case <-exit:
			return
		}
	}
}

func AddRequest(request_list []Request, request_chan chan<- Request, exit_chan chan<- bool) {
	for _, rq := range request_list {
		request_chan <- rq
	}
	time.Sleep(100 * time.Millisecond)
	exit_chan <- true
}

func HandleCustomerReq(rq Request) {
	defer wg.Done()
	m.Lock()
	for {
		for worker, available := range workers {
			if available == true {
				current_clock := clock
				fmt.Printf("Worker %v handle request %v(require %v),clock = %v \n", worker.name, rq.name, rq.time, current_clock)
				workers[worker] = false
				worker.available_at = worker.available_at + rq.time
				m.Unlock()
				time.Sleep(time.Second * time.Duration(rq.time))
				fmt.Printf("Worker %v finished request %v,clock = %v \n", worker.name, rq.name, current_clock+rq.time)
				if current_clock+rq.time >= clock {
					clock = current_clock + rq.time
				}
				workers[worker] = true
				return
			}
		}
	}
}

func ModifyClock(time int) {
	m.Lock()
	defer m.Unlock()
	if clock+time > clock {
		clock = clock + time
	}
}

func main() {
	request_list := []Request{Request{"X", 9}, Request{"W", 5}, Request{"D", 4}, Request{"K", 6}, Request{"V", 8}, Request{"H", 1}, Request{"Y", 6}, Request{"T", 8}}

	workers[&Worker{"A", 0}] = true
	workers[&Worker{"B", 0}] = true
	workers[&Worker{"C", 0}] = true

	request_chan := make(chan Request, 1)
	exit_chan := make(chan bool, 1)

	wg.Add(1)

	go AddRequest(request_list, request_chan, exit_chan)
	go Handle(request_chan, exit_chan)

	wg.Wait()

	fmt.Println("Current clock = ", clock)
	fmt.Println("Done")
}
