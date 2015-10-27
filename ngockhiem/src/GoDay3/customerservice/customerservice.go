package main

import (
	"fmt"
	//"math/rand"
	"sync"
	"time"
)

var (
	wg        sync.WaitGroup
	m         sync.Mutex
	c         sync.Mutex
	clock     int              = 0
	workers   map[*Worker]bool = make(map[*Worker]bool)
	exit_chan chan bool        = make(chan bool)
)

type Worker struct {
	name         string
	available_at int
}
type Request struct {
	name    string
	require int
	appear  int
}

func Handle(ch chan Request, exit chan bool) {
	defer wg.Done()
	for {
		select {
		case rq := <-ch:
			wg.Add(1)
			go HandleCustomerReq(rq)
		case <-exit:
			close(exit)
			return
		}
	}
}

func AddRequest(request_list []Request, request_chan chan<- Request, exit_chan chan<- bool) {
	for _, rq := range request_list {

		fmt.Printf("%v, Add request %v(require %v) \n", clock, rq.name, rq.require)
		request_chan <- rq
		time.Sleep(time.Second)

	}
	time.Sleep(100 * time.Millisecond)
	exit_chan <- true
}

func HandleCustomerReq(rq Request) {
	defer wg.Done()
	m.Lock()
	for {
		for worker, available := range workers {
			if available == true && worker.available_at <= clock {
				current_clock := clock
				fmt.Printf("%v, Start converstion: %v - %v \n", clock, worker.name, rq.name)
				workers[worker] = false
				worker.available_at = worker.available_at + rq.require
				m.Unlock()
				for {
					if current_clock+rq.require <= clock {
						//time.Sleep(time.Second * time.Duration(rq.require))
						fmt.Printf("%v, End conversation: %v - %v \n", current_clock+rq.require, worker.name, rq.name)
						// if current_clock+rq.require >= clock {
						// 	clock = current_clock + rq.require
						// }
						ModifyClock(current_clock, rq.require)
						workers[worker] = true
						return
					}
				}
			}
		}
	}
}

func ModifyClock(curr_clock, time int) {
	c.Lock()
	defer c.Unlock()
	if curr_clock+time > clock {
		clock = curr_clock + time
	}
}

func IncreaseClock() {
	for {
		can_support := false
		for _, available := range workers {
			if available == true {
				can_support = true
			}
		}
		select {
		case <-exit_chan:
			clock++
			time.Sleep(time.Second)
		default:
			if !can_support {
				clock++
				time.Sleep(time.Second)
			}
		}

	}
}

func main() {
	request_list := []Request{Request{"X", 9, 1}, Request{"W", 5, 2}, Request{"D", 4, 4}, Request{"K", 6, 4}, Request{"V", 8, 5}, Request{"H", 1, 6}, Request{"Y", 6, 7}, Request{"T", 8, 9}}

	workers[&Worker{"A", 0}] = true
	workers[&Worker{"B", 0}] = true
	workers[&Worker{"C", 0}] = true

	request_chan := make(chan Request)

	wg.Add(1)

	go AddRequest(request_list, request_chan, exit_chan)
	go Handle(request_chan, exit_chan)
	go IncreaseClock()
	wg.Wait()

	fmt.Println("Done")
}
