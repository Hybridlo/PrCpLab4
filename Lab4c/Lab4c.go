package main

import (
	"sync"
	"time"
)

func main() {
	g := NewGraphManager()
	var wg sync.WaitGroup

	wg.Add(5)

	go nodesThread(&g, &wg)
	time.Sleep(5 * time.Second)
	go waysThread(&g, &wg)
	time.Sleep(5 * time.Second)
	go costChangerThread(&g, &wg)
	time.Sleep(5 * time.Second)
	go wayFinderThread(&g, &wg)
	go wayFinderThread(&g, &wg)

	wg.Wait()
}