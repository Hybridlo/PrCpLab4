package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"strconv"
)

func costChangerThread(g *GraphManager, wg *sync.WaitGroup) {
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
		node := nodes[i_rand]

		if len(node.ways) == 0 {
			g.justUnlock()
			continue
		}

		i_rand_2 := r.Intn(len(node.ways))
		way := node.ways[i_rand_2]
		i_rand_3 := r.Intn(10) + 1
		g.modifyWay(nodes[i_rand].number, way.getOther(node).number, i_rand_3)

		fmt.Printf("Change way cost from node " + strconv.Itoa(node.number) + " to " + strconv.Itoa(way.getOther(node).number) + " new cost " + strconv.Itoa(i_rand_3) + "\n")

		time.Sleep(2000 * time.Millisecond)
	}
}