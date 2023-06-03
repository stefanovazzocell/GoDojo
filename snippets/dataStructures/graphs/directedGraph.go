package graphs

// A Node in a directed graph
type Node struct {
	Name   string
	Childs []*Node
	// Parents []*Node
}

// Returns true if a path is found between a start and end node
func (start *Node) HasPathTo(end *Node) bool {
	if start == nil || end == nil {
		return false
	}
	if start == end {
		return true
	}
	// Perform Breath-First Search,
	// keeping track of visited nodes
	nodesQueue := []*Node{start}
	seenNodes := map[*Node]bool{}
	for {
		node := nodesQueue[0]
		// Search Childs
		if node != nil {
			for _, child := range node.Childs {
				if child == nil {
					continue // Child is invalid
				}
				if child == end {
					return true // Child is end node
				}
				if !seenNodes[child] {
					seenNodes[child] = true
					nodesQueue = append(nodesQueue, child) // Queue
				}
			}
		}
		// Quit if queue empty, otherwise move to next node
		if len(nodesQueue) <= 1 {
			return false
		}
		nodesQueue = nodesQueue[1:]
	}
}
