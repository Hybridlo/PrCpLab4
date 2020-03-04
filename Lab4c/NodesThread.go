package main

import (
	"sync"
	"time"
)

func nodesThread(g *GraphManager, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		g.addOrRemoveNode()

		

		time.Sleep(1 * time.Second)
	}
}