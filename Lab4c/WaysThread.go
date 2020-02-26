package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"strconv"
)

func waysThread(g *GraphManager, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		nodes := g.getNodesAndLock()

		if len(nodes) < 2 {
			g.justUnlock()
			continue
		}

		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		do := r.Intn(2)

		switch do {
		case 0:
			i_rand := r.Intn(len(nodes))
			node := nodes[i_rand]

			if len(node.ways) == 0 {
				g.justUnlock()
				continue
			}

			i_rand_2 := r.Intn(len(node.ways))
			way := node.ways[i_rand_2]
			g.removeWay(node.number, way.getOther(node).number)
			fmt.Printf("Remove way from node " + strconv.Itoa(node.number) + " to " + strconv.Itoa(way.getOther(node).number) + "\n")
		case 1:
			i_rand := r.Intn(len(nodes))
			if len(nodes[i_rand].ways) + 1 == len(nodes) {
				continue
			}
			i_rand_2 := r.Intn(len(nodes))
			wayExists := false

			for _, w := range nodes[i_rand].ways {
				otherNode := w.getOther(nodes[i_rand])
				if nodes[i_rand_2] == otherNode {
					wayExists = true
				}
			}

			for i_rand == i_rand_2 || wayExists {
				wayExists = false
				i_rand_2 = r.Intn(len(nodes))

				for _, w := range nodes[i_rand].ways {
					otherNode := w.getOther(nodes[i_rand])
					if nodes[i_rand_2] == otherNode {
						wayExists = true
					}
				}
			}
			i_rand_3 := r.Intn(10) + 1
			g.addWay(nodes[i_rand].number, nodes[i_rand_2].number, i_rand_3)
			fmt.Printf("Add way from node " + strconv.Itoa(nodes[i_rand].number) + " to " + strconv.Itoa(nodes[i_rand_2].number) + " cost " + strconv.Itoa(i_rand_3) + "\n")
		}

		time.Sleep(700 * time.Millisecond)
	}
}