package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"strconv"
)

func wayFinderThread(g *GraphManager, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		nodes := g.getNodesAndLock()

		if len(nodes) < 2 {
			g.justUnlock()
			continue
		}

		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		i_rand := r.Intn(len(nodes))
		i_rand_2 := r.Intn(len(nodes))
		for i_rand == i_rand_2 {
			i_rand_2 = r.Intn(len(nodes))
		}

		visited := make([]*Node, 1)
		cost := g.getCostAndUnlock(nodes[i_rand].number, nodes[i_rand_2].number, &visited)
		if cost == -1 {
			fmt.Printf("Way not found\n")
		} else {
			fmt.Printf("\tWay from node " + strconv.Itoa(nodes[i_rand].number) + " to " + strconv.Itoa(nodes[i_rand_2].number) + " costs " + strconv.Itoa(cost) + "\n")
		}

		time.Sleep(500 * time.Millisecond)
	}
}