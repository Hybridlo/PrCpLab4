package main

import (
	"sync"
	"time"
)

func wayFinderThread(g *GraphManager, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		g.findWay()

		time.Sleep(500 * time.Millisecond)
	}
}