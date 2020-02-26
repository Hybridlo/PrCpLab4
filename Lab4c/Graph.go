package main

import (
	"fmt"
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

func (g *GraphManager) addNode(number int) {
	g.RUnlock()
	
	var newNode Node
	newNode.number = number

	g.Lock()
	defer g.Unlock()

	g.nodes = append(g.nodes, &newNode)
}

func (g *GraphManager) removeNode(number int) {
	g.RUnlock()
	g.Lock()
	defer g.Unlock()

	i, rmNode := g.getNode(number)
	for _, w := range rmNode.ways {
		w.getOther(rmNode).removeWay(w)
	}
	
	g.nodes = append(g.nodes[:i], g.nodes[i+1:]...)
}

func (g *GraphManager) addWay(a int, b int, weight int) {
	g.RUnlock()
	g.Lock()
	defer g.Unlock()

	_, aNode := g.getNode(a)
	_, bNode := g.getNode(b)

	way := NewWay(aNode, bNode, weight)

	aNode.ways = append(aNode.ways, &way)
	bNode.ways = append(bNode.ways, &way)
}

func (g *GraphManager) removeWay(a int, b int) {
	g.RUnlock()
	g.Lock()
	defer g.Unlock()

	_, aNode := g.getNode(a)
	_, bNode := g.getNode(b)

	rmWay := aNode.getWay(bNode)

	aNode.removeWay(rmWay)
	bNode.removeWay(rmWay)
}

func (g *GraphManager) modifyWay(a int, b int, newWeight int) {
	g.RUnlock()
	g.Lock()
	defer g.Unlock()

	_, aNode := g.getNode(a)
	_, bNode := g.getNode(b)

	changeWay := aNode.getWay(bNode)

	changeWay.weight = newWeight
}

func (g *GraphManager) getNodesAndLock() []*Node {
	g.RLock()
	return g.nodes
}

func (g *GraphManager) justUnlock() {
	g.RUnlock()
}

func (g *GraphManager) getCostAndUnlock(a int, b int, visited *[]*Node) int {
	defer g.RUnlock()

	return g.dfs(a, b, visited)
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