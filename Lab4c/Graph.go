package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

type Way struct {
	a *Node
	b *Node
	weight int
}

func NewWay(a *Node, b *Node, weight int) Way {
	w := Way{a, b, weight}
	return w
}

func (w Way) getOther(from *Node) *Node {
	if from == w.a {
		return w.b
	} else {
		if from == w.b {
			return w.a
		} else {
			return nil
		}
	}
}







type Node struct {
	ways []*Way
	number int
}

func (n *Node) getWay(toNode *Node) *Way {
	for _, w := range n.ways {
		if w.getOther(n) == toNode {
			return w
		}
	}

	return nil
}

func (n *Node) removeWay(rmWay *Way) {
	for i, w := range n.ways {
		if w == rmWay {
			n.ways = append(n.ways[:i], n.ways[i+1:]...)
			break
		}
	}
}







type GraphManager struct {
	lockRead chan bool
	lockWrite chan bool
	lockWriteRequest chan bool
	nodes []*Node
}

func NewGraphManager() GraphManager {
	var g GraphManager
	g.lockRead = make(chan bool, 100)
	g.lockWrite = make(chan bool, 1)
	g.lockWriteRequest = make(chan bool, 1)
	return g
}

func (g *GraphManager) RLock() {
	for len(g.lockWrite) > 0 || len(g.lockWriteRequest) > 0 {
	}

	g.lockRead<-true
}

func (g *GraphManager) RUnlock() {
	if len(g.lockRead) == 0 {
		fmt.Printf("RUnlock before RLock\n")
	}

	<- g.lockRead
}

func (g *GraphManager) Lock() {
	g.lockWriteRequest <- true
	for len(g.lockRead) > 0 || len(g.lockWrite) > 0 {
	}

	g.lockWrite <- true
	<- g.lockWriteRequest
}

func (g *GraphManager) Unlock() {
	if len(g.lockWrite) == 0 {
		fmt.Printf("Unlock before Lock")
	}

	<- g.lockWrite
}

func (g *GraphManager) getNode(number int) (int, *Node) {
	for i, n := range g.nodes {
		if n.number == number {
			return i, n
		}
	}

	return 0, nil
}

func (g *GraphManager) addOrRemoveNode() {

	g.Lock()
	defer g.Unlock()

	nodes := g.nodes

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	do := r.Intn(2)

	switch do {
	case 0:
		if len(nodes) < 5 {
			return
		}

		i_rand := r.Intn(len(nodes))
		node := nodes[i_rand]

		for _, w := range node.ways {
			w.getOther(node).removeWay(w)
		}
		
		g.nodes = append(g.nodes[:i_rand], g.nodes[i_rand+1:]...)

		fmt.Printf("Remove node " + strconv.Itoa(node.number) + "\n")
	case 1:
		max := 0
		for _, n := range nodes {
			if n.number > max {
				max = n.number
			}
		}
		max = max + 1
		
		var newNode Node
		newNode.number = max

		g.nodes = append(g.nodes, &newNode)

		fmt.Printf("Add node " + strconv.Itoa(max) + "\n")
	}
}

func (g *GraphManager) addOrRemoveWay() {
	g.Lock()
	defer g.Unlock()

	nodes := g.nodes

	if len(nodes) < 2 {
		return
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	do := r.Intn(2)

	switch do {
	case 0:
		i_rand := r.Intn(len(nodes))
		node := nodes[i_rand]

		if len(node.ways) == 0 {
			return
		}

		i_rand_2 := r.Intn(len(node.ways))
		way := node.ways[i_rand_2]

		node.removeWay(way)
		way.getOther(node).removeWay(way)
		
		fmt.Printf("Remove way from node " + strconv.Itoa(node.number) + " to " + strconv.Itoa(way.getOther(node).number) + "\n")
	case 1:
		i_rand := r.Intn(len(nodes))
		if len(nodes[i_rand].ways) + 1 == len(nodes) {
			return
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

		way := NewWay(nodes[i_rand], nodes[i_rand_2], i_rand_3)

		nodes[i_rand].ways = append(nodes[i_rand].ways, &way)
		nodes[i_rand_2].ways = append(nodes[i_rand_2].ways, &way)

		fmt.Printf("Add way from node " + strconv.Itoa(nodes[i_rand].number) + " to " + strconv.Itoa(nodes[i_rand_2].number) + " cost " + strconv.Itoa(i_rand_3) + "\n")
	}
}

func (g *GraphManager) modifyWay() {
	g.Lock()
	defer g.Unlock()

	nodes := g.nodes

	if len(nodes) < 2 {
		return
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	i_rand := r.Intn(len(nodes))
	node := nodes[i_rand]

	if len(node.ways) == 0 {
		return
	}

	i_rand_2 := r.Intn(len(node.ways))
	way := node.ways[i_rand_2]
	i_rand_3 := r.Intn(10) + 1

	way.weight = i_rand_3

	fmt.Printf("Change way cost from node " + strconv.Itoa(node.number) + " to " + strconv.Itoa(way.getOther(node).number) + " new cost " + strconv.Itoa(i_rand_3) + "\n")
}

func (g *GraphManager) findWay() {
	g.RLock()
	defer g.RUnlock()

	nodes := g.nodes

	if len(nodes) < 2 {
		return
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	i_rand := r.Intn(len(nodes))
	i_rand_2 := r.Intn(len(nodes))

	for i_rand == i_rand_2 {
		i_rand_2 = r.Intn(len(nodes))
	}

	visited := make([]*Node, 1)
	cost := g.dfs(nodes[i_rand].number, nodes[i_rand_2].number, &visited)
	
	if cost == -1 {
		fmt.Printf("Way not found\n")
	} else {
		fmt.Printf("\tWay from node " + strconv.Itoa(nodes[i_rand].number) + " to " + strconv.Itoa(nodes[i_rand_2].number) + " costs " + strconv.Itoa(cost) + "\n")
	}
}

func (g *GraphManager) dfs(a int, b int, visited *[]*Node) int {

	if a == b {
		return 0
	}

	var result int
	_, aNode := g.getNode(a)

	*visited = append(*visited, aNode)

	for _, w := range aNode.ways {
		found := true
		otherNode := w.getOther(aNode)

		for _, n := range *visited {
			if n == otherNode {
				found = false
				break
			}
		}

		if found {
			result = g.dfs(otherNode.number, b, visited)
		}

		if result != -1 {
			return result + w.weight
		}
	}

	return -1
}