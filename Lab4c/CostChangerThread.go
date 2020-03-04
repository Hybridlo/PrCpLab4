package main

import (
	"sync"
	"time"
)

func costChangerThread(g *GraphManager, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		g.modifyWay()

		time.Sleep(2000 * time.Millisecond)
	}
}