package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"strconv"
)

func nodesThread(g *GraphManager, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		nodes := g.getNodesAndLock()

		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		do := r.Intn(2)

		switch do {
		case 0:
			if len(nodes) < 5 {
				g.justUnlock()
				continue
			}

			i_rand := r.Intn(len(nodes))
			node := nodes[i_rand]
			g.removeNode(node.number)
			fmt.Printf("Remove node " + strconv.Itoa(node.number) + "\n")
		case 1:
			max := 0
			for _, n := range nodes {
				if n.number > max {
					max = n.number
				}
			}
			max = max + 1
			g.addNode(max)
			fmt.Printf("Add node " + strconv.Itoa(max) + "\n")
		}

		

		time.Sleep(1 * time.Second)
	}
}