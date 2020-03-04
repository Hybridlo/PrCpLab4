package main

import (
	"sync"
	"time"
)

func waysThread(g *GraphManager, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		g.addOrRemoveWay()

		time.Sleep(700 * time.Millisecond)
	}
}